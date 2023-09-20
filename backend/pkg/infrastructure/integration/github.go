package integration

import (
	"context"
	"errors"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/shurcooL/graphql"
	"os"
	"regexp"
	"time"
)

type githubService struct {
	client *graphql.Client
	ctx    context.Context
}

func NewGithubService() service.GithubService {
	graphqlClient := NewGraphQLClient(
		"https://api.github.com/graphql",
		os.Getenv("GITHUB_TOKEN"),
	)
	return &githubService{
		client: graphqlClient.Client,
		ctx:    context.Background(),
	}
}

func (s *githubService) GetStarCountByURL(url string) (int, error) {
	owner, name, err := s.extractOwnerAndName(url)
	if err != nil {
		return 0, err
	}
	var query struct {
		Repository struct {
			Stargazers struct {
				TotalCount graphql.Int
			}
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": graphql.String(owner),
		"name":  graphql.String(name),
	}
	if err := s.client.Query(s.ctx, &query, variables); err != nil {
		return 0, fmt.Errorf("Failed to parse query: %v", err)
	}
	return int(query.Repository.Stargazers.TotalCount), nil
}

func (s *githubService) FindTopStarRepo(minStar, limit int, language string) ([]*dto.GithubRepo, error) {
	const topicLimit = 3
	var query struct {
		Search struct {
			Edges []struct {
				Node struct {
					Repository struct {
						Name             graphql.String
						Owner            struct{ Login graphql.String }
						Stargazers       struct{ TotalCount graphql.Int }
						PrimaryLanguage  struct{ Name graphql.String }
						Description      graphql.String
						HomepageURL      graphql.String
						CreatedAt        time.Time
						RepositoryTopics struct {
							Nodes []struct {
								Topic struct{ Name graphql.String }
							}
						} `graphql:"repositoryTopics(first: $topicLimit)"`
					} `graphql:"... on Repository"`
				}
			}
		} `graphql:"search(query: $query, type: REPOSITORY, first: $first)"`
	}
	created := "2023-05-01"
	variables := map[string]interface{}{
		"topicLimit": graphql.Int(topicLimit),
		"query":      graphql.String(fmt.Sprintf("stars:>%d created:>%s language: %s", minStar, created, language)),
		"first":      graphql.Int(limit),
	}

	err := s.client.Query(s.ctx, &query, variables)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse query: %v", err)
	}

	var result []*dto.GithubRepo
	for _, edge := range query.Search.Edges {
		repo := edge.Node.Repository
		var topics []model.Topic
		for _, node := range repo.RepositoryTopics.Nodes {
			topics = append(topics, model.Topic{Name: string(node.Topic.Name)})
		}
		result = append(result, &dto.GithubRepo{
			Name:            string(repo.Name),
			Owner:           string(repo.Owner.Login),
			StarCount:       int(repo.Stargazers.TotalCount),
			HomepageURL:     string(repo.HomepageURL),
			Description:     string(repo.Description),
			PrimaryLanguage: string(repo.PrimaryLanguage.Name),
			CreatedAt:       repo.CreatedAt,
			Topics:          topics,
		})
	}
	return result, nil
}

func (s *githubService) extractOwnerAndName(repoURL string) (owner, name string, err error) {
	r := regexp.MustCompile(`github\.com/([^/]+)/([^/]+)`)
	matches := r.FindStringSubmatch(repoURL)
	if len(matches) < 3 {
		return "", "", errors.New("invalid GitHub repository URL")
	}
	return matches[1], matches[2], nil
}

//
//type
