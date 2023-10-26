package utils

import "sort"

type KeyValuePair struct {
	Key   int
	Value float64
}

func SortMapKeysByFloatValue(inputMap map[int]float64, threshold float64) []int {
	keyValuePairs := make([]KeyValuePair, 0)
	for key, value := range inputMap {
		// 閾値の設定
		if value > threshold {
			keyValuePairs = append(keyValuePairs, KeyValuePair{key, value})
		}
	}
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})
	sortedKeys := make([]int, len(keyValuePairs))
	for i, pair := range keyValuePairs {
		sortedKeys[i] = pair.Key
	}
	return sortedKeys
}
