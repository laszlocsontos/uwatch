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
	"reflect"
	"testing"

	"github.com/lcsontos/uwatch/youtube"
)

type longVideoUrlTestCase struct {
	want      string
	videoType VideoType
	videoId   string
}

type parseVideoUrlTestCase struct {
	want     *ParsedVideoUrl
	videoUrl string
}

const (
	_EMPTY_STRING      = ""
	_INVALID_VIDEO_ID  = "abcdefgh"
	_INVALID_VIDEO_URL = "http://www.google.com"
)

var _EMPTY_PVU = ParsedVideoUrl{"", 0}

var longVideoUrlData = []longVideoUrlTestCase{
	{"meditations-with-sri-chinmoy-vol--1", YouTube, "5VAYzfvNI1w"},
	{"zen---musica-de-relajacion-y-balance-espiritual", YouTube, "Gd0TiO0iMfc"},
	{_EMPTY_STRING, YouTube, _EMPTY_STRING},
	{_EMPTY_STRING, Vimeo, _EMPTY_STRING},
	{_EMPTY_STRING, Youku, _EMPTY_STRING},
	{_EMPTY_STRING, Rutube, _EMPTY_STRING},
	{_EMPTY_STRING, YouTube, _INVALID_VIDEO_ID},
	{_EMPTY_STRING, Vimeo, _INVALID_VIDEO_ID},
	{_EMPTY_STRING, Youku, _INVALID_VIDEO_ID},
	{_EMPTY_STRING, Rutube, _INVALID_VIDEO_ID},
}

var parseVideoUrlData = []parseVideoUrlTestCase{
	{&ParsedVideoUrl{"Gd0TiO0iMfc", YouTube}, "https://www.youtube.com/watch?v=Gd0TiO0iMfc"},
	{&ParsedVideoUrl{"Gd0TiO0iMfc", YouTube}, "http://youtu.be/Gd0TiO0iMfc"},
	{&ParsedVideoUrl{"5VAYzfvNI1w", YouTube}, "https://www.youtube.com/watch?v=5VAYzfvNI1w"},
	{&ParsedVideoUrl{"5VAYzfvNI1w", YouTube}, "http://youtu.be/5VAYzfvNI1w"},
	{nil, _INVALID_VIDEO_URL},
}

func TestLongVideoUrl(t *testing.T) {
	for _, data := range longVideoUrlData {

		videoCatalog, err := youtube.New()

		if err != nil {
			t.Fatal(err)
		}

		result, err := LongVideoUrl(videoCatalog, data.videoType, data.videoId)

		if err != nil && data.want != _EMPTY_STRING {
			t.Fatal(err)
		}

		got := _EMPTY_STRING

		if result != nil {
			got = result.UrlPath
		}

		want := data.want

		if got != want {
			t.Fatalf("TestLongVideoUrl(%v, %v).urlPath = %v, but wanted %v",
				data.videoType, data.videoId, got, want)
		}
	}
}

func TestParseVideoUrl(t *testing.T) {
	for _, data := range parseVideoUrlData {
		result, err := ParseVideoUrl(data.videoUrl)

		if err != nil && data.want != nil {
			t.Fatal(err)
		}

		got := _EMPTY_PVU

		if result != nil {
			got = *result
		}

		want := _EMPTY_PVU

		if w := data.want; w != nil {
			want = *w
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("TestParseVideoUrl(%v) = %v, but wanted %v",
				data.videoUrl, got, want)
		}
	}
}
