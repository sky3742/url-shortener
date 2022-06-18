package main

import (
	"testing"

	"url.shortener/model"
	"url.shortener/tool"
)

func TestUnique(t *testing.T) {
	var urls []model.Url_Shortener
	urls = append(urls, model.Url_Shortener{Id: tool.GenerateId(nil, 6)})
	if tool.IsUniqueId(urls, urls[0].Id) {
		t.Error("Expected: True, Received: False")
	}
}

func TestNotUnique(t *testing.T) {
	var urls []model.Url_Shortener
	urls = append(urls, model.Url_Shortener{Id: tool.GenerateId(nil, 6)})
	if !tool.IsUniqueId(urls, "123456") {
		t.Error("Expected: False, Received: True")
	}
}

func TestIdLengthIsSix(t *testing.T) {
	id := tool.GenerateId(nil, 6)
	if len(id) != 6 {
		t.Errorf("Length: Expected 6, Received %d", len(id))
	}
}

func TestIdLengthIsNotSix(t *testing.T) {
	length := 5
	id := tool.GenerateId(nil, length)
	if len(id) == 6 {
		t.Errorf("Length: Expected %d, Received %d", length, len(id))
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	id := tool.GenerateId(nil, 6)
	if !tool.IsAlphaNumericOnly(id) {
		t.Errorf("Expected alphanumeric, received %s", id)
	}
}

func TestIsNotAlphaNumeric(t *testing.T) {
	id := "@@@@@@"
	if tool.IsAlphaNumericOnly(id) {
		t.Errorf("Expected alphanumeric, received %s", id)
	}
}
