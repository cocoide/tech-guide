package integration

import (
	"context"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

type GraphQLClient struct {
	Client *graphql.Client
}

func NewGraphQLClient(url, tkn string) *GraphQLClient {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tkn},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient(url, httpClient)
	return &GraphQLClient{
		Client: client,
	}
}
