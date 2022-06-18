package service

import (
	"time"

	"url.shortener/model"
	"url.shortener/tool"
)

var shortenedUrl []model.Url_Shortener = []model.Url_Shortener{}

func GetNewUrl(destination string) model.Url_Shortener {
	id := tool.GenerateId(shortenedUrl, 6)
	url := model.Url_Shortener{
		Id:            id,
		Destination:   destination,
		RedirectCount: 0,
		CreatedAt:     time.Now().UnixMilli(),
	}
	saveUrl(url)
	return url
}

func GetUrls() []model.Url_Shortener {
	return shortenedUrl
}

func GetUrlById(id string) (model.Url_Shortener, bool) {
	for _, url := range shortenedUrl {
		if url.Id == id {
			return url, true
		}
	}
	return model.Url_Shortener{}, false
}

func UpdateUrlRedirectCount(url model.Url_Shortener) {
	url.IncrementCount()
	saveUrl(url)
}

func saveUrl(url model.Url_Shortener) {
	for i := range shortenedUrl {
		if shortenedUrl[i].Id == url.Id {
			shortenedUrl[i] = url
			return
		}
	}

	shortenedUrl = append(shortenedUrl, url)
}
