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

func (err *NoSuchVideoError) Error() string {
	return fmt.Sprintf("Video with Id %s does not exist", err.VideoId)
}

func NewVideoRecord(videoId, title, description string, publishedAt time.Time) *VideoRecord {
	return &VideoRecord{Description: description, PublishedAt: publishedAt, VideoId: videoId, Title: title}
}

func (videoRecord *VideoRecord) String() string {
	return fmt.Sprintf("[%v] %v: %v", videoRecord.VideoId, videoRecord.Title, videoRecord.Description)
}
