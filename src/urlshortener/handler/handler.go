package handler

import (
	"net/http"
	"urlshortener/model"
	"github.com/gorilla/mux"
	"encoding/json"
)

func ShortUrl(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)

		url := vars["url"]

		urlInfo := &model.UrlInfo{
			Url: url,
		}

		shortUrl := model.PutUrl(urlInfo)

		var res struct {
			Result string `json:"result"`
		}

		res.Result = shortUrl

		json.NewEncoder(writer).Encode(res)
}

func ResolveShorturl(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	shortUrl := vars["short_url"]

	res, urlInfo := model.GetUrl(shortUrl)

	if !res {
		http.NotFound(writer,request)
	}

	http.Redirect(writer, request, urlInfo.GetLink(), http.StatusFound)
}