package util

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func HandleError(rw http.ResponseWriter, req *http.Request, err, apperr error, isAppErr bool) bool {
	if err == nil {
		return false
	}

	if status := http.StatusBadRequest; isAppErr {
		http.Error(rw, apperr.Error(), status)
	} else {
		status = http.StatusInternalServerError

		// TODO Generalize error handling
		http.Error(rw, fmt.Sprintf("INTERNAL ERROR: %s", err.Error()), status)
		log.Printf(err.Error())
	}

	return true
}

func HandlePanic(err interface{}, rw http.ResponseWriter, req *http.Request) {
	var stack [4096]byte

	runtime.Stack(stack[:], false)

	log.Printf(
		"Handler for %s[url:%s, data:%s] has failed with %s\nStack trace:\n %s\n",
		req.Method, req.URL, req.Form, err, stack[:])

	http.Error(rw, "SYSTEM ERROR", http.StatusInternalServerError)
}
