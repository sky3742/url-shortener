package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"url.shortener/service"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	urls := service.GetUrls()
	fmt.Println(urls)
	if err := json.NewEncoder(w).Encode(urls); err != nil {
		w.Write([]byte(err.Error()))
	}
}

func GetStatById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url, exist := service.GetUrlById(id)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(url)
}
