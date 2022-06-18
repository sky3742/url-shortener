package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const port = "8888"

func main() {
	r := mux.NewRouter()
	r.Use(jsonMiddleware)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
