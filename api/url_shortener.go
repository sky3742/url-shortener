package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"url.shortener/service"
)

type url_Request struct {
	Url string `json:"url"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var body url_Request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	url := service.GetNewUrl(body.Url)
	body.Url = fmt.Sprintf("%s/%s", r.Host, url.Id)
	json.NewEncoder(w).Encode(body)
}

func RedirectWithId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	url, exist := service.GetUrlById(id)
	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	service.UpdateUrlRedirectCount(url)
	http.Redirect(w, r, url.Destination, http.StatusTemporaryRedirect)
}
