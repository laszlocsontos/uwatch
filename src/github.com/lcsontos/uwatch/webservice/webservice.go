package webservice

import (
	"fmt"
	"net/http"
)

type VideoType int

const (
	YouTube = (iota)

	// Reserved for future implementation
	Vimeo
	Youku
	Rutube
)

type InvalidVideoUrl struct {
	VideoUrl string
}

type ParsedVideoUrl struct {
	videoId   string
	videoType VideoType
}

type LengthenVideoUrl struct {
	ParsedVideoUrl
	urlId   int64
	urlPath string
}

func (err *InvalidVideoUrl) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video URL", err.VideoUrl)
}

func LongVideoUrl(videoType VideoType, videoId string) (*LengthenVideoUrl, error) {
	return nil, nil
}

func ParseVideoUrl(videoUrl string) (*ParsedVideoUrl, error) {
	return nil, nil
}

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
}

func (url *LengthenVideoUrl) String() string {
	return ""
}
