package webservice

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"

	"github.com/lcsontos/uwatch/catalog"
	"github.com/lcsontos/uwatch/youtube"
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
	VideoId   string
	VideoType VideoType
}

type LengthenVideoUrl struct {
	ParsedVideoUrl
	UrlId   int64
	UrlPath string
}

type UnsupportedVideoType struct {
	VideoType VideoType
}

type urlPattern struct {
	videoType VideoType
	pattern   *regexp.Regexp
}

var urlPatterns = []urlPattern{
	urlPattern{YouTube, regexp.MustCompile("http.+youtube\\.com\\/watch\\?v=(\\S+)")},
	urlPattern{YouTube, regexp.MustCompile("http.+youtu\\.be\\/(\\S+)")},
}

var videoCatalogRegistry = make(map[VideoType]catalog.VideoCatalog)

func (err *InvalidVideoUrl) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video URL", err.VideoUrl)
}

func (err *UnsupportedVideoType) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video type", err.VideoType)
}

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	for key, value := range vars {
		fmt.Fprintf(rw, "Key: %v, Value: %v\n", key, value)
	}
}

func (url *LengthenVideoUrl) String() string {
	return ""
}

func init() {
	var err error

	// TODO create factory for creating wrapper objects to
	// video sharing services
	videoCatalogRegistry[YouTube], err = youtube.New()

	if err != nil {
		panic(err)
	}
}

func longVideoUrl(videoType VideoType, videoId string) (*LengthenVideoUrl, error) {
	if videoType != YouTube {
		return nil, &UnsupportedVideoType{videoType}
	}

	videoCatalog := videoCatalogRegistry[videoType]

	videoRecord, err := videoCatalog.SearchByID(videoId)

	if err != nil {
		return nil, err
	}

	// TODO implement title converter here
	title := videoRecord.Title

	lengthenVideoUrl := &LengthenVideoUrl{
		ParsedVideoUrl{videoId, videoType},
		0, title,
	}

	fmt.Println(title)

	return lengthenVideoUrl, nil
}

func parseVideoUrl(videoUrl string) (*ParsedVideoUrl, error) {
	if videoUrl == "" {
		return nil, &InvalidVideoUrl{""}
	}

	parsedVideoUrl := &ParsedVideoUrl{VideoType: unknown}

	for _, urlPattern := range urlPatterns {
		matches := urlPattern.pattern.FindStringSubmatch(videoUrl)

		if matches != nil {
			parsedVideoUrl.VideoId = matches[1]
			parsedVideoUrl.VideoType = urlPattern.videoType

			break
		}
	}

	if parsedVideoUrl.VideoType == unknown {
		return nil, &InvalidVideoUrl{videoUrl}
	}

	return parsedVideoUrl, nil
}
