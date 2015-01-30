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

package service

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/lcsontos/uwatch/catalog"
	"github.com/lcsontos/uwatch/normalizer"
	"github.com/lcsontos/uwatch/store"
)

type InvalidVideoUrlError struct {
	VideoUrl string
}

type UnsupportedVideoTypeError struct {
	VideoType catalog.VideoType
}

type urlPattern struct {
	videoType catalog.VideoType
	pattern   *regexp.Regexp
}

var urlPatterns = []urlPattern{
	urlPattern{catalog.YouTube, regexp.MustCompile("http.+youtube\\.com\\/watch\\?v=(\\S+)")},
	urlPattern{catalog.YouTube, regexp.MustCompile("http.+youtu\\.be\\/(\\S+)")},
}

func (err *InvalidVideoUrlError) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video URL", err.VideoUrl)
}

func (err *UnsupportedVideoTypeError) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video type", err.VideoType)
}

func LongVideoUrl(videoCatalog catalog.VideoCatalog, videoType catalog.VideoType, videoId string, req *http.Request) (*catalog.LongVideoUrl, error) {
	if videoType != catalog.YouTube {
		return nil, &UnsupportedVideoTypeError{videoType}
	}

	videoRecord, err := videoCatalog.SearchByID(videoId)

	if err != nil {
		return nil, err
	}

	normalizedTitle := normalizer.Normalize(videoRecord.Title)

	urlId, err := store.PutVideoRecord(videoRecord, req)

	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("%d/%s", urlId, normalizedTitle)

	longVideoUrl := &catalog.LongVideoUrl{
		catalog.VideoKey{videoId, videoType},
		videoRecord.Title, urlId, urlPath,
	}

	return longVideoUrl, nil
}

func ParseVideoUrl(videoUrl string) (*catalog.VideoKey, error) {
	if videoUrl == "" {
		return nil, &InvalidVideoUrlError{""}
	}

	videoKey := &catalog.VideoKey{VideoType: catalog.Unknown}

	for _, urlPattern := range urlPatterns {
		matches := urlPattern.pattern.FindStringSubmatch(videoUrl)

		if matches != nil {
			videoKey.VideoId = matches[1]
			videoKey.VideoType = urlPattern.videoType

			break
		}
	}

	if videoKey.VideoType == catalog.Unknown {
		return nil, &InvalidVideoUrlError{videoUrl}
	}

	return videoKey, nil
}
