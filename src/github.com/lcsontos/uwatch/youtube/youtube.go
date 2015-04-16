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

package youtube

import (
	"flag"
	"net/http"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"

	"github.com/lcsontos/uwatch/catalog"
)

const _PART = "id,snippet"

type Service struct {
	httpClient     *http.Client
	youTubeService *youtube.Service
}

var developerKey string

func New() (*Service, error) {
	service, err := NewWithRoundTripper(nil)

	return service, err
}

func NewWithRoundTripper(roundTripper http.RoundTripper) (*Service, error) {
	transport := &transport.APIKey{
		Key:       developerKey,
		Transport: roundTripper,
	}

	client := &http.Client{
		Transport: transport,
	}

	service, err := youtube.New(client)

	if err != nil {
		return nil, err
	}

	return &Service{httpClient: client, youTubeService: service}, nil
}

func (service Service) SearchByID(videoId string) (*catalog.VideoRecord, error) {
	call := service.getVideosListCall(videoId)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	if cap(response.Items) == 0 {
		return nil, &catalog.NoSuchVideoError{videoId}
	}

	videoId = response.Items[0].Id
	videoKey := &catalog.VideoKey{videoId, catalog.YouTube}

	publishedTime := parsePublishedAt(response.Items[0].Snippet.PublishedAt)

	title := response.Items[0].Snippet.Title

	videoRecord := catalog.NewVideoRecord(videoKey, publishedTime, title)

	return videoRecord, nil
}

func (service Service) SearchByTitle(title string, maxResults int64) ([]catalog.VideoRecord, error) {
	call := service.getSearchListCall(title, maxResults)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	items := response.Items
	itemCount := cap(items)

	var videoRecords []catalog.VideoRecord

	if itemCount == 0 {
		videoRecords = []catalog.VideoRecord{}
	} else {
		videoRecords = make([]catalog.VideoRecord, itemCount)

		for index, item := range items {
			videoKey := &catalog.VideoKey{item.Id.VideoId, catalog.YouTube}

			publishedTime := parsePublishedAt(item.Snippet.PublishedAt)
			title := item.Snippet.Title

			videoRecord := catalog.NewVideoRecord(videoKey, publishedTime, title)

			videoRecords[index] = *videoRecord
		}
	}

	return videoRecords, nil
}

func (service Service) getSearchListCall(searchTerm string, maxResults int64) *youtube.SearchListCall {
	call := service.youTubeService.Search.List(_PART)

	call.Q(searchTerm)
	call.MaxResults(maxResults)

	return call
}

func (service Service) getVideosListCall(videoId string) *youtube.VideosListCall {
	call := service.youTubeService.Videos.List(_PART)

	return call.Id(videoId)
}

func init() {
	developerKey = *flag.String("YouTubeDevKey", "", "")
}

func parsePublishedAt(publishedAt string) time.Time {
	if publishedTime, err := time.Parse(time.RFC3339Nano, publishedAt); err == nil {
		return publishedTime
	} else {
		return time.Unix(0, 0)
	}
}

func setDeveloperKey(devKey string) {
	developerKey = devKey
}
