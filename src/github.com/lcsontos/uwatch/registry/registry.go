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

package registry

import (
	"appengine"
	"appengine/urlfetch"

	"net/http"

	"github.com/lcsontos/uwatch/catalog"
	"github.com/lcsontos/uwatch/store"
	"github.com/lcsontos/uwatch/youtube"
)

func GetVideoCatalog(videoType catalog.VideoType, req *http.Request) catalog.VideoCatalog {
	// TODO : this package should be appengine agnostic

	context := appengine.NewContext(req)

	transport := &urlfetch.Transport{Context: context}

	videoCatalog, err := youtube.NewWithRoundTripper(transport)

	if err != nil {
		panic(err)
	}

	return videoCatalog
}

func GetVideoStore(req *http.Request) store.VideoStore {
	// TODO

	return nil
}
