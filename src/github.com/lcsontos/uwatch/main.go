package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lcsontos/uwatch/webservice"
)

func init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/long_video_url/{videoType}/{videoId}", webservice.ServeHTTP)
	router.HandleFunc("/api/parse_video_url/{videoUrl}", webservice.ServeHTTP)

	http.Handle("/", router)
}
