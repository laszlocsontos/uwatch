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

type VideoKey struct {
	VideoId   string
	VideoType VideoType
}

type LongVideoUrl struct {
	Id int64
	*VideoKey

	CreatedAt   time.Time
	PublishedAt time.Time

	NormalizedTitle string
	UrlPath         string
}

type VideoRecord struct {
	*VideoKey

	PublishedAt time.Time

	Title string
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

func NewLongVideoUrl(videoRecord *VideoRecord) *LongVideoUrl {
	return &LongVideoUrl{
		VideoKey:    videoRecord.VideoKey,
		CreatedAt:   time.Now(),
		PublishedAt: videoRecord.PublishedAt,
	}
}

func NewVideoRecord(videoKey *VideoKey, publishedAt time.Time, title string) *VideoRecord {
	return &VideoRecord{
		VideoKey: videoKey, PublishedAt: publishedAt, Title: title}
}

func (url *LongVideoUrl) String() string {
	return ""
}

func (videoKey *VideoKey) String() string {
	return fmt.Sprintf("%s@%s", videoKey.VideoType, videoKey.VideoId)
}

func (videoType VideoType) String() string {
	return videoTypesStringMap[videoType]
}

func (videoRecord *VideoRecord) String() string {
	return fmt.Sprintf("[%v] %v: %v", videoRecord.VideoId, videoRecord.Title)
}

func init() {
	for videoType, videoTypeName := range videoTypesStringMap {
		videoTypesLookupMap[videoTypeName] = videoType
	}
}
