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
	want     string
	videoUrl string
}

var longVideoUrlData = []longVideoUrlTestCase{}

var parseVideoUrlData = []parseVideoUrlTestCase{}

func TestLongVideoUrl(t *testing.T) {

}

func TestParseVideoUrl(t *testing.T) {

}
