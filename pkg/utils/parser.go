package utils

import "github.com/graph-gophers/dataloader"

func NewStringsFromKeys(keys dataloader.Keys) []string {
	var result []string

	for _, key := range keys {
		result = append(result, key.String())
	}

	return result
}
