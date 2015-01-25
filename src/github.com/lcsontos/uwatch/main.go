package main

import (
	"fmt"
	"github.com/lcsontos/uwatch/webservice"
	"net/http"
)

func init() {
	http.HandleFunc("/api/", webservice.ServeHTTP)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
