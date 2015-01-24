package webservice

import (
	"fmt"
	"net/http"
	"regexp"
)

type VideoType int

const (
	YouTube = (iota)

	// Reserved for future implementation
	Vimeo
	Youku
	Rutube

	// Internal use only!
	unknown = -1
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

type urlPattern struct {
	videoType VideoType
	pattern   *regexp.Regexp
}

var urlPatterns = []urlPattern{
	urlPattern{YouTube, regexp.MustCompile("http.+youtube\\.com\\/watch\\?v=(\\S+)")},
	urlPattern{YouTube, regexp.MustCompile("http.+youtu\\.be\\/(\\S+)")},
}

func (err *InvalidVideoUrl) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video URL", err.VideoUrl)
}

func LongVideoUrl(videoType VideoType, videoId string) (*LengthenVideoUrl, error) {
	return nil, nil
}

func ParseVideoUrl(videoUrl string) (*ParsedVideoUrl, error) {
	if videoUrl == "" {
		return nil, &InvalidVideoUrl{""}
	}

	parsedVideoUrl := &ParsedVideoUrl{videoType: unknown}

	for _, urlPattern := range urlPatterns {
		matches := urlPattern.pattern.FindStringSubmatch(videoUrl)

		if matches != nil {
			parsedVideoUrl.videoId = matches[1]
			parsedVideoUrl.videoType = urlPattern.videoType

			break
		}
	}

	if parsedVideoUrl.videoType == unknown {
		return nil, &InvalidVideoUrl{videoUrl}
	}

	return parsedVideoUrl, nil
}

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
}

func (url *LengthenVideoUrl) String() string {
	return ""
}
