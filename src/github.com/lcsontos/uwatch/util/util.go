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

		logError(err, req)

		msg := fmt.Sprintf("INTERNAL ERROR: %s", err.Error())

		http.Error(rw, msg, status)
	}

	return true
}

func HandlePanic(err interface{}, rw http.ResponseWriter, req *http.Request) {
	logError(err, req)

	http.Error(rw, "SYSTEM ERROR", http.StatusInternalServerError)
}

func logError(err interface{}, req *http.Request) {
	var stack [4096]byte

	runtime.Stack(stack[:], false)

	log.Printf(
		"Handler for %s[url:%s, data:%s] has failed with %s\nStack trace:\n %s\n",
		req.Method, req.URL, req.Form, err, stack[:])
}
