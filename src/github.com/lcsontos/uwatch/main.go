package main

import (
	"fmt"
	"net/http"

	"github.com/lcsontos/uwatch/webservice"
)

func init() {
	http.HandleFunc("/api/", webservice.ServeHTTP)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
