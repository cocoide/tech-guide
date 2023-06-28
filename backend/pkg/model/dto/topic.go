package dto

import "sort"

type TopicWeight struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

func GetTopWeightedTopics(topicWeights []TopicWeight, limit int) []TopicWeight {
	sort.Slice(topicWeights, func(i, j int) bool {
		return topicWeights[i].Weight > topicWeights[j].Weight
	})

	if len(topicWeights) > limit {
		topicWeights = topicWeights[:limit]
	}

	return topicWeights
}
