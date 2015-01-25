package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lcsontos/uwatch/webservice"
)

func init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/long_video_url/{videoType}/{videoId}", webservice.GetLongVideoUrl).Methods("GET")
	router.HandleFunc("/api/parse_video_url/{videoUrl}", webservice.GetParseVideoUrl).Methods("GET")

	http.Handle("/", router)
}
