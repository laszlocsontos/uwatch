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

package webservice

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lcsontos/uwatch/catalog"
	"github.com/lcsontos/uwatch/registry"
	"github.com/lcsontos/uwatch/service"
	"github.com/lcsontos/uwatch/util"
)

func GetLongVideoUrl(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	videoTypeName := vars["videoTypeName"]
	videoId := vars["videoId"]

	videoType, err := catalog.GetVideoTypeByName(videoTypeName)

	if apperr, isAppErr := err.(*catalog.InvalidVideoTypeNameError); util.HandleError(rw, req, err, apperr, isAppErr) {
		return
	}

	videoCatalog := registry.GetVideoCatalog(videoType, req)

	videoKey := &catalog.VideoKey{VideoId: videoId, VideoType: videoType}

	longVideoUrl, err := service.LongVideoUrl(videoCatalog, videoKey, req)

	handledError := false

	switch apperr := err.(type) {
	case *service.UnsupportedVideoTypeError:
		handledError = util.HandleError(rw, req, err, apperr, true)
	case *catalog.NoSuchVideoError:
		handledError = util.HandleError(rw, req, err, apperr, true)
	default:
		handledError = util.HandleError(rw, req, err, apperr, false)
	}

	if handledError {
		return
	}

	json.NewEncoder(rw).Encode(*longVideoUrl)
}

func GetParseVideoUrl(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	videoUrl := vars["videoUrl"]

	log.Printf("videoUrl=%s", videoUrl)

	videoKey, err := service.ParseVideoUrl(videoUrl)

	if apperr, isAppErr := err.(*service.InvalidVideoUrlError); util.HandleError(rw, req, err, apperr, isAppErr) {
		return
	}

	json.NewEncoder(rw).Encode(*videoKey)
}
