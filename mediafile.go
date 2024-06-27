package main

import (
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
	fileType  MediaFileType
	title     string
	season    string
	episode   string
	year      string
	extension string
}

// Constructor for a MediaFile
func NewMediaFile(fileName string) (*MediaFile, error) {
	log.Logger.WithField(
		"component",
		"mediafile").Infof("processing file: %s\n", fileName)

	// Determine the type of media: TV or MOVIE and parse the file name
	segments := strings.Split(fileName, " - ")
	if len(segments) >= 3 {
		// Process TV files
		log.Logger.WithField(
			"component",
			"mediafile").Info("file appears to be a TV show")
		pattern := `([a-zA-Z0-9\s]+)\s-\s(\w+)\s-\s([a-zA-Z0-9\s\.]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)$`
		matches, err := _extractMatches(fileName, pattern)
		if err != nil {
			return &MediaFile{}, err
		}

		return &MediaFile{
			fileType:  TV,
			title:     matches[1],
			season:    matches[2],
			episode:   matches[3],
			year:      matches[4],
			extension: matches[5],
		}, nil

	} else if len(segments) < 3 && len(segments) > 0 {
		// Process Movie files
		log.Logger.WithField(
			"component",
			"mediafile").Info("file appears to be a movie")
		pattern := `([a-zA-Z0-9\s\-]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)`
		matches, err := _extractMatches(fileName, pattern)
		if err != nil {
			return &MediaFile{}, err
		}

		return &MediaFile{
			fileType:  MOVIE,
			title:     matches[1],
			season:    "", // not present on a Movie
			episode:   "", // not present on a Movie
			year:      matches[2],
			extension: matches[3],
		}, nil

	} else {
		// File name pattern is unknown
		log.Logger.WithField(
			"component",
			"mediafile").Errorf("unable to determine pattern of %s\n", fileName)
		return &MediaFile{}, errors.New("failed to create MediaFile")
	}
}

// extract matches from a file name using the regex pattern
func _extractMatches(fileName string, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Logger.Error("error compiling pattern")
		return nil, err
	}

	log.Logger.WithField(
		"component",
		"mediafile").Debugf("using regex pattern %s\n", re.String())

	matches := re.FindStringSubmatch(fileName)
	if len(matches) == 0 {
		return nil, errors.New("failed to find matches")
	}

	log.Logger.WithField(
		"component",
		"mediafile").Debugf("found %d field(s) in file name\n", len(matches))

	return matches, nil
}
