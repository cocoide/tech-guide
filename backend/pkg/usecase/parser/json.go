package parser

import (
	"encoding/json"
	"fmt"
)

func Serialize(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func Deserialize[T any](data string) (T, error) {
	var result T
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, fmt.Errorf("Failed to unmarshal: %v", err)
	}
	return result, nil
}
