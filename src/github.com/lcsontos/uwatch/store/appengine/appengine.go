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

	// "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lcsontos/uwatch/catalog"
	"github.com/lcsontos/uwatch/registry"
	"github.com/lcsontos/uwatch/store"
)

type Store struct {
	context *appengine.Context
}

type StoreFactory struct {
}

const _KIND_LONG_VIDEO_URL = "LongVideoUrl"
const _KIND_VIDEO_KEY = "VideoKey"

func (store Store) FindLongVideoUrlByID(id int64) (*catalog.LongVideoUrl, error) {
	key := datastore.NewKey(*store.context, _KIND_VIDEO_KEY, "", id, nil)

	videoKey, err := store.getVideoKey(key)

	if err != nil {
		return nil, err
	}

	return store.FindLongVideoUrlByVideoKey(videoKey)
}

func (store Store) FindLongVideoUrlByVideoKey(videoKey *catalog.VideoKey) (*catalog.LongVideoUrl, error) {
	key := datastore.NewKey(*store.context, _KIND_LONG_VIDEO_URL, videoKey.String(), 0, nil)

	longVideoUrl, err := store.getLongVideoUrl(key)

	if err != nil {
		return nil, err
	}

	longVideoUrl.VideoKey = videoKey

	return longVideoUrl, nil
}

func (storeFactory StoreFactory) NewStore(args interface{}) store.VideoStore {
	req := args.(*http.Request)

	context := appengine.NewContext(req)

	return Store{context: &context}
}

func (store Store) SaveLongVideoUrl(longVideoUrl *catalog.LongVideoUrl) error {
	transactionOptions := &datastore.TransactionOptions{XG: true}

	return datastore.RunInTransaction(*store.context, func(context appengine.Context) error {
		return doSaveLongVideoUrl(context, longVideoUrl)
	}, transactionOptions)
}

func (s Store) getLongVideoUrl(key *datastore.Key) (*catalog.LongVideoUrl, error) {
	var longVideoUrl catalog.LongVideoUrl

	err := datastore.Get(*s.context, key, &longVideoUrl)

	switch {
	case err == nil:
		return &longVideoUrl, nil
	case err == datastore.ErrNoSuchEntity:
		parts := strings.Split(key.StringID(), "@")

		var videoKey *catalog.VideoKey = nil

		if len(parts) == 2 {
			videoType, _ := strconv.Atoi(parts[0])

			videoKey = &catalog.VideoKey{
				VideoType: catalog.VideoType(videoType),
				VideoId:   parts[1],
			}
		}

		return nil, &store.NoSuchLongVideoUrl{
			Id: key.IntID(), VideoKey: videoKey,
		}
	default:
		return nil, err
	}
}

func (s Store) getVideoKey(key *datastore.Key) (*catalog.VideoKey, error) {
	var videoKey catalog.VideoKey

	err := datastore.Get(*s.context, key, &videoKey)

	switch {
	case err == nil:
		return &videoKey, nil
	case err == datastore.ErrNoSuchEntity:
		return nil, &store.NoSuchLongVideoUrl{Id: key.IntID()}
	default:
		return nil, err
	}
}

func doSaveLongVideoUrl(context appengine.Context, longVideoUrl *catalog.LongVideoUrl) error {
	videoKey := longVideoUrl.VideoKey

	// log.Printf("videoKey: %s", *videoKey)

	videoKeyKey := datastore.NewIncompleteKey(context, _KIND_VIDEO_KEY, nil)
	videoKeyKey, err := datastore.Put(context, videoKeyKey, videoKey)

	// log.Printf("videoKeyKey, err: %s, %s", *videoKeyKey, err)

	if err != nil {
		return err
	}

	longVideoUrl.Id = videoKeyKey.IntID()

	// log.Printf("longVideoUrl: %s", longVideoUrl)

	longVideoUrlKey := datastore.NewKey(context, _KIND_LONG_VIDEO_URL, videoKey.String(), 0, nil)
	longVideoUrlKey, err = datastore.Put(context, longVideoUrlKey, longVideoUrl)

	// log.Printf("longVideoUrlKey, err: %s", longVideoUrlKey, err)

	return err
}

func init() {
	registry.RegisterVideoStore(&StoreFactory{})
}
