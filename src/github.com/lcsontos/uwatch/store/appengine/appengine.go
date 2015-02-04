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

package appengine

import (
	"appengine"
	"appengine/datastore"

	"github.com/lcsontos/uwatch/catalog"
)

type Store struct {
	context *appengine.Context
}

const _KIND = "LongVideoUrl"

func (store Store) FindLongVideoUrlByID(id int64) (*catalog.LongVideoUrl, error) {
	key := datastore.NewKey(*store.context, _KIND, "", id, nil)

	return store.getLongVideoUrl(key)
}

func (store Store) FindLongVideoUrlByVideoKey(videoKey *catalog.VideoKey) (*catalog.LongVideoUrl, error) {
	key := datastore.NewKey(*store.context, _KIND, "", 0, nil)

	return store.getLongVideoUrl(key)
}

func (store Store) getLongVideoUrl(key *datastore.Key) (*catalog.LongVideoUrl, error) {
	var longVideoUrl catalog.LongVideoUrl

	err := datastore.Get(*store.context, key, &longVideoUrl)

	if err != nil {
		return nil, err
	}

	return &longVideoUrl, nil
}
