package webservice

import (
	"appengine"
	"appengine/urlfetch"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/gorilla/mux"

	"github.com/lcsontos/uwatch/catalog"
	"github.com/lcsontos/uwatch/youtube"
)

type VideoType int

const (
	YouTube VideoType = (iota)

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

type LengthenedVideoUrl struct {
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

var videoCatalogRegistryRWM sync.RWMutex

var videoTypesLookupMap = make(map[string]VideoType)

var videoTypesStringMap = map[VideoType]string{
	YouTube: "YouTube", Vimeo: "Vimeo", Youku: "Youku", Rutube: "Rutube",
}

func (err *InvalidVideoUrl) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video URL", err.VideoUrl)
}

func (err *UnsupportedVideoType) Error() string {
	return fmt.Sprintf("\"%s\" is an invalid video type", err.VideoType)
}

func GetLongVideoUrl(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	videoType := vars["videoType"]
	videoId := vars["videoId"]

	lengthenedVideoUrl, err := longVideoUrl(videoTypesLookupMap[videoType], videoId, req)

	apperr, isAppErr := err.(*UnsupportedVideoType)

	if handledError(rw, req, err, apperr, isAppErr) {
		return
	}

	json.NewEncoder(rw).Encode(*lengthenedVideoUrl)
}

func GetParseVideoUrl(rw http.ResponseWriter, req *http.Request) {
}

func (videoType VideoType) String() string {
	return videoTypesStringMap[videoType]
}

func (url *LengthenedVideoUrl) String() string {
	return ""
}

func init() {
	// Initialize videoTypesLookupMap
	initVideoTypesLookupMap()

	// Initialize catalog registry
	// initVideoCatalogRegistry()
}

// func initVideoCatalogRegistry() {
// 	var err error

// 	// TODO create factory for creating wrapper objects to
// 	// video sharing services
// 	videoCatalogRegistry[YouTube], err = youtube.New()

// 	if err != nil {
// 		panic(err)
// 	}
// }

func initVideoTypesLookupMap() {
	for videoType, videoTypeName := range videoTypesStringMap {
		videoTypesLookupMap[videoTypeName] = videoType
	}
}

// I needed this "hack", because app engine requires a http.Request object to
// instanciate Transport objects. Why on earth do I have to do this???
// Reference: https://cloud.google.com/appengine/docs/go/urlfetch/
func getVideoCatalog(videoType VideoType, req *http.Request) catalog.VideoCatalog {
	videoCatalogRegistryRWM.RLock()

	if videoCatalog, ok := videoCatalogRegistry[videoType]; ok {
		videoCatalogRegistryRWM.RUnlock()

		return videoCatalog
	}

	videoCatalogRegistryRWM.RUnlock()

	videoCatalogRegistryRWM.Lock()

	if videoCatalog, ok := videoCatalogRegistry[videoType]; ok {
		videoCatalogRegistryRWM.Unlock()

		return videoCatalog
	}

	context := appengine.NewContext(req)

	transport := &urlfetch.Transport{Context: context}

	videoCatalog, err := youtube.NewWithRoundTripper(transport)

	if err != nil {
		panic(err)
	}

	videoCatalogRegistry[videoType] = videoCatalog

	videoCatalogRegistryRWM.Unlock()

	return videoCatalog
}

func handledError(rw http.ResponseWriter, req *http.Request, err, apperr error, isAppErr bool) bool {
	if err == nil {
		return false
	}

	if status := http.StatusBadRequest; isAppErr {
		http.Error(rw, apperr.Error(), status)
	} else {
		status = http.StatusInternalServerError

		// TODO Generalize error handling
		http.Error(rw, "INTERNAL ERROR", status)
		log.Printf(err.Error())
	}

	return true
}

func longVideoUrl(videoType VideoType, videoId string, req *http.Request) (*LengthenedVideoUrl, error) {
	if videoType != YouTube {
		return nil, &UnsupportedVideoType{videoType}
	}

	videoCatalog := getVideoCatalog(videoType, req)

	videoRecord, err := videoCatalog.SearchByID(videoId)

	if err != nil {
		return nil, err
	}

	// TODO implement title converter here
	title := videoRecord.Title

	LengthenedVideoUrl := &LengthenedVideoUrl{
		ParsedVideoUrl{videoId, videoType},
		0, title,
	}

	// log.Println(title)

	return LengthenedVideoUrl, nil
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
