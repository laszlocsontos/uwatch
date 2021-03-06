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

package html

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/lcsontos/uwatch/registry"
	"github.com/lcsontos/uwatch/util"
)

type invalidLongUrlError struct {
	longUrl string
}

var page *template.Template
var pattern *regexp.Regexp

func (err *invalidLongUrlError) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid long video URL", err.longUrl)
}

func ProcessTemplate(rw http.ResponseWriter, req *http.Request) {
	tc := make(map[string]interface{})

	urlId, _, err := getPathTokens(req.URL)

	if util.HandleError(rw, req, err, err, true) {
		return
	}

	tc["URL"], err = getVideoUrl(urlId, req)

	if util.HandleError(rw, req, err, err, true) {
		return
	}

	if err := page.Execute(rw, tc); err != nil {
		util.HandleError(rw, req, err, nil, false)
	}
}

func getPathTokens(url *url.URL) (int64, string, error) {
	matches := pattern.FindStringSubmatch(url.Path)

	if matches == nil {
		return -1, "", &invalidLongUrlError{url.RequestURI()}
	}

	urlIdString, normalizedTitle := matches[1], matches[2]

	urlId, err := strconv.ParseInt(urlIdString, 10, 64)

	if err != nil {
		return -1, "", err
	}

	return urlId, normalizedTitle, nil
}

func getVideoUrl(urlId int64, req *http.Request) (string, error) {
	videoStore := registry.GetVideoStore(req)

	longVideoUrl, err := videoStore.FindLongVideoUrlByID(urlId)

	if err != nil {
		return "", err
	}

	// TODO

	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", longVideoUrl.VideoKey.VideoId)

	return url, nil
}

func init() {
	page = template.Must(template.ParseFiles("templates/html/index.html"))
	pattern = regexp.MustCompile("/(\\S+)/(\\S+)/?")
}
