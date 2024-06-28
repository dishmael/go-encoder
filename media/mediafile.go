package media

import (
	"encoding/json"
	"errors"
	log "go-encoder/logger"
	"regexp"
	"strings"
)

// Enum
type MediaFileType int

const (
	UNKNOWN MediaFileType = iota
	MOVIE
	TV
)

type MediaFile struct {
	Title     string
	Season    string
	Episode   string
	Year      string
	Extension string
}

// Constructor for a MediaFile
func NewMediaFile(fileName string) (*MediaFile, error) {
	log.Logger.
		WithField("component", "mediafile").
		Infof("processing file: %s", fileName)

	// Initialization Logic Flow
	// 1. determine the number of segments (getSegments)
	// 2. determine the media type (getMediaFileType)
	// 3. determine the regexp pattern (getRegExpPattern)
	// 4. extract the fields (extractMatches)
	// 5. populate the MediaFile with the matches (populateMediaFile)

	segments := getSegments(fileName)
	if segments == 0 {
		return nil, errors.New("failed to find any segments in file name")
	}

	mft := getMediaFileType(segments)
	if mft == UNKNOWN {
		return nil, errors.New("unable to determine media type based on file name")
	}

	pattern := getRegExPattern(mft)
	if pattern == "" {
		return nil, errors.New("unable to determine pattern of file name")
	}

	fields, err := extractMatches(fileName, pattern)
	if err != nil {
		return nil, err
	}

	mf := populateMediaFile(fields, mft)
	log.Logger.
		WithField("component", "mediafile").
		Infof("%+v", mf)

	return mf, nil
}

// getSegments
func getSegments(fileName string) int {
	segments := len(strings.Split(fileName, " - "))
	log.Logger.
		WithField("component", "mediafile").
		Debugf("found %d segments in file name", segments)
	return segments
}

// getMediaFileType
func getMediaFileType(segments int) MediaFileType {
	switch {
	case segments >= 3:
		log.Logger.
			WithField("component", "mediafile").
			Info("file appears to be a tv show")
		return TV
	case segments > 0:
		log.Logger.
			WithField("component", "mediafile").
			Info("file appears to be a movie")
		return MOVIE
	default:
		log.Logger.
			WithField("component", "mediafile").
			Error("unable to determine media type based on file name")
		return UNKNOWN
	}
}

// getRegExpPattern
func getRegExPattern(mft MediaFileType) string {
	switch mft {
	case TV:
		return `([a-zA-Z0-9\s]+)\s-\s(S[0-9]+E[0-9]+)\s-\s([a-zA-Z0-9\s\.]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)$`
	case MOVIE:
		return `([a-zA-Z0-9\s\-]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)`
	default:
		log.Logger.
			WithField("component", "mediafile").
			Error("unable to determine pattern of file name")
		return ""
	}
}

// extrctMatches
func extractMatches(fileName, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Logger.
			WithField("component", "mediafile").
			Error("error compiling pattern")
		return nil, err
	}

	log.Logger.
		WithField("component", "mediafile").
		Debugf("using regex pattern %s", re.String())

	matches := re.FindStringSubmatch(fileName)
	if len(matches) == 0 {
		log.Logger.
			WithField("component", "mediafile").
			Error("no matches found")
		return nil, errors.New("failed to find matches")
	}

	log.Logger.
		WithField("component", "mediafile").
		Debugf("found %d field(s) in file name", len(matches))
	return matches, nil
}

// populateMediaFile
func populateMediaFile(fields []string, mft MediaFileType) *MediaFile {
	mf := &MediaFile{}

	if len(fields) != 4 && len(fields) != 6 {
		log.Logger.
			WithField("component", "mediafile").
			Errorf("wrong number of fields: got %d want %d or %d", len(fields), 4, 6)
		return mf
	}

	switch mft {
	case TV:
		mf.Title = fields[1]
		mf.Season = fields[2]
		mf.Episode = fields[3]
		mf.Year = fields[4]
		mf.Extension = fields[5]
	case MOVIE:
		mf.Title = fields[1]
		mf.Year = fields[2]
		mf.Extension = fields[3]
	default:
		// NOOP: We should never get here
	}

	return mf
}

// toString
func (mf *MediaFile) toString() string {
	out, err := json.Marshal(mf)
	if err != nil {
		log.Logger.
			WithField("component", "mediafile").
			Error("failed to stringify MediaInfo")
		return ""
	}

	return string(out)
}
