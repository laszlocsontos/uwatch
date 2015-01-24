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

func (err *InvalidVideoUrl) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video URL", err.VideoUrl)
}

func LongVideoUrl(videoType VideoType, videoId string) (string, error) {
	return "", nil
}

func ParseVideoUrl(videoUrl string) (VideoType, string, error) {
	return YouTube, "", nil
}

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
}
