package tool

import (
	"regexp"

	"github.com/google/uuid"
	"url.shortener/model"
)

func GenerateId(arr []model.Url_Shortener, length int) string {
	for {
		id := uuid.New().String()[1 : length+1]
		if IsUniqueId(arr, id) && IsAlphaNumericOnly(id) {
			return id
		}
	}
}

func IsUniqueId(arr []model.Url_Shortener, id string) bool {
	for _, url := range arr {
		if url.Id == id {
			return false
		}
	}
	return true
}

func IsAlphaNumericOnly(id string) bool {
	matched, _ := regexp.Match("[^a-zA-Z0-9]", []byte(id))
	return !matched
}
