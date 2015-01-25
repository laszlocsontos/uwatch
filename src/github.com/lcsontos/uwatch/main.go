package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"

	"github.com/lcsontos/uwatch/webservice"
)

type handlerFuncType func(http.ResponseWriter, *http.Request)

func handleSafely(handlerFunc handlerFuncType) handlerFuncType {
	return func(rw http.ResponseWriter, req *http.Request) {
		// Recover Panic
		defer func() {
			if err := recover(); err != nil {
				panicHandler(err, rw, req)
			}
		}()

		// Call original handler
		handlerFunc(rw, req)
	}
}

func init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(
		"/api/long_video_url/{videoType}/{videoId}",
		handleSafely(webservice.GetLongVideoUrl)).Methods("GET")

	router.HandleFunc(
		"/api/parse_video_url/{videoUrl}",
		handleSafely(webservice.GetParseVideoUrl)).Methods("GET")

	http.Handle("/", router)
}

func panicHandler(err interface{}, rw http.ResponseWriter, req *http.Request) {
	var stack [4096]byte

	runtime.Stack(stack[:], false)

	log.Printf(
		"Handler for %s[url:%s, data:%s] has failed with %s\nStack trace:\n %s\n",
		req.Method, req.URL, req.Form, err, stack[:])

	http.Error(rw, "SYSTEM ERROR", http.StatusInternalServerError)
}
