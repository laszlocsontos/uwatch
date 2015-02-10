//
// Copyright (C) 2015-present  Laszlo Csontos
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>. */
//

package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lcsontos/uwatch/html"
	"github.com/lcsontos/uwatch/util"
	"github.com/lcsontos/uwatch/webservice"

	_ "github.com/lcsontos/uwatch/store/appengine"
	_ "github.com/lcsontos/uwatch/youtube/appengine"
)

type handlerFuncType func(http.ResponseWriter, *http.Request)

func handleSafely(handlerFunc handlerFuncType) handlerFuncType {
	return func(rw http.ResponseWriter, req *http.Request) {
		// Recover Panic
		defer func() {
			if err := recover(); err != nil {
				util.HandlePanic(err, rw, req)
			}
		}()

		// Call original handler
		handlerFunc(rw, req)
	}
}

func init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(
		"/api/long_video_url/{videoTypeName}/{videoId}",
		handleSafely(webservice.GetLongVideoUrl)).Methods("GET")

	router.HandleFunc(
		"/api/parse_video_url",
		handleSafely(webservice.GetParseVideoUrl)).Methods("POST")

	router.HandleFunc(
		"/{urlId}/{urlPath}",
		handleSafely(html.ProcessTemplate)).Methods("GET")

	http.Handle("/", router)
}
