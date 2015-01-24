package webservice

import (
	"testing"
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
	_EMPTY             = ""
	_INVALID_VIDEO_ID  = "abcdefgh"
	_INVALID_VIDEO_URL = "http://www.google.com"
)

var longVideoUrlData = []longVideoUrlTestCase{
	{"meditations-with-sri-chinmoy-vol--1", YouTube, "5VAYzfvNI1w"},
	{"zen---musica-de-relajacion-y-balance-espiritual", YouTube, "Gd0TiO0iMfc"},
	{_EMPTY, YouTube, _EMPTY},
	{_EMPTY, Vimeo, _EMPTY},
	{_EMPTY, Youku, _EMPTY},
	{_EMPTY, Rutube, _EMPTY},
	{_EMPTY, YouTube, _INVALID_VIDEO_ID},
	{_EMPTY, Vimeo, _INVALID_VIDEO_ID},
	{_EMPTY, Youku, _INVALID_VIDEO_ID},
	{_EMPTY, Rutube, _INVALID_VIDEO_ID},
}

var parseVideoUrlData = []parseVideoUrlTestCase{
	{&ParsedVideoUrl{"Gd0TiO0iMfc", YouTube}, "https://www.youtube.com/watch?v=Gd0TiO0iMfc"},
	{&ParsedVideoUrl{"Gd0TiO0iMfc", YouTube}, "http://youtu.be/Gd0TiO0iMfc"},
	{&ParsedVideoUrl{"5VAYzfvNI1w", YouTube}, "https://www.youtube.com/watch?v=5VAYzfvNI1w"},
	{&ParsedVideoUrl{"5VAYzfvNI1w", YouTube}, "http://youtu.be/5VAYzfvNI1w"},
	{nil, _INVALID_VIDEO_URL},
}

func TestLongVideoUrl(t *testing.T) {

}

func TestParseVideoUrl(t *testing.T) {

}
