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

package catalog

import (
	"fmt"
	"time"
)

type VideoType int

type InvalidVideoTypeNameError struct {
	VideoTypeName string
}

type ParsedVideoUrl struct {
	VideoId   string
	VideoType VideoType
}

type LengthenedVideoUrl struct {
	ParsedVideoUrl
	Title   string
	UrlId   int64
	UrlPath string
}

type VideoRecord struct {
	Id          int64
	Description string
	PublishedAt time.Time
	VideoId     string
	Title       string
}

type NoSuchVideoError struct {
	VideoId string
}

type VideoCatalog interface {
	SearchByID(videoId string) (*VideoRecord, error)
	SearchByTitle(title string, maxResults int64) ([]VideoRecord, error)
}

const (
	YouTube VideoType = (iota)

	// Reserved for future implementation
	Vimeo
	Youku
	Rutube

	// Internal use only!
	Unknown = -1
)

var videoTypesLookupMap = make(map[string]VideoType)

var videoTypesStringMap = map[VideoType]string{
	YouTube: "YouTube",
	Vimeo:   "Vimeo",
	Youku:   "Youku",
	Rutube:  "Rutube",
}

func (err *InvalidVideoTypeNameError) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video name", err.VideoTypeName)
}

func (err *NoSuchVideoError) Error() string {
	return fmt.Sprintf("Video with Id %s does not exist", err.VideoId)
}

func GetVideoTypeByName(videoTypeName string) (VideoType, error) {
	if videoType, ok := videoTypesLookupMap[videoTypeName]; !ok {
		return Unknown, &InvalidVideoTypeNameError{videoTypeName}
	} else {
		return videoType, nil
	}
}

func NewVideoRecord(videoId, title, description string, publishedAt time.Time) *VideoRecord {
	return &VideoRecord{Description: description, PublishedAt: publishedAt, VideoId: videoId, Title: title}
}

func (url *LengthenedVideoUrl) String() string {
	return ""
}

func (videoType VideoType) String() string {
	return videoTypesStringMap[videoType]
}

func (videoRecord *VideoRecord) String() string {
	return fmt.Sprintf("[%v] %v: %v", videoRecord.VideoId, videoRecord.Title, videoRecord.Description)
}

func init() {
	for videoType, videoTypeName := range videoTypesStringMap {
		videoTypesLookupMap[videoTypeName] = videoType
	}
}
