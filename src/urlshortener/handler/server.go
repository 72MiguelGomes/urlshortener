package handler

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func StartServer() {

	r := mux.NewRouter()

	r.HandleFunc("/short/{url}", ShortUrl)
	r.HandleFunc("/resolve/{short_url}", ResolveShorturl)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":7272", nil))
}