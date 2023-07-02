package util

import "sort"

type KeyValuePair[T int | string] struct {
	Key   T
	Value float64
}

func SortMapKeysByFloatValue[T int | string](inputMap map[T]float64, threshold float64) []T {
	keyValuePairs := make([]KeyValuePair[T], 0)
	for key, value := range inputMap {
		// 閾値の設定
		if value > threshold {
			keyValuePairs = append(keyValuePairs, KeyValuePair[T]{key, value})
		}
	}
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})
	sortedKeys := make([]T, len(keyValuePairs))
	for i, pair := range keyValuePairs {
		sortedKeys[i] = pair.Key
	}
	return sortedKeys
}
