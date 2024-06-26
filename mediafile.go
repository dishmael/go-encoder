package main

import (
	"errors"
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
	// Determine the type of media: TV or MOVIE and parse the file name
	segments := strings.Split(fileName, " - ")
	if len(segments) >= 3 {
		// Process TV files
		Logger.Infof("%s appears to be a TV show\n", fileName)
		fileType := TV
		pattern := `^([a-zA-Z0-9\s]+)\s-\s(\w+)\s-\s([a-zA-Z0-9\s\.]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)$`
		re, err := regexp.Compile(pattern)
		if err != nil {
			Logger.Error("Error compiling pattern")
			return &MediaFile{}, err
		}

		matches := re.FindStringSubmatch(fileName)
		return &MediaFile{
			fileType:  fileType,
			title:     matches[1],
			season:    matches[2],
			episode:   matches[3],
			year:      matches[4],
			extension: matches[5],
		}, nil

	} else if len(segments) < 3 && len(segments) > 0 {
		// Process Movie files
		Logger.Infof("%s appears to be a movie\n", fileName)
		fileType := MOVIE
		pattern := `([a-zA-Z0-9\s\-]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)`
		re, err := regexp.Compile(pattern)
		if err != nil {
			Logger.Error("Error compiling pattern")
			return &MediaFile{}, err
		}

		matches := re.FindStringSubmatch(fileName)
		return &MediaFile{
			fileType:  fileType,
			title:     matches[1],
			season:    "", // not present on a Movie
			episode:   "", // not present on a Movie
			year:      matches[2],
			extension: matches[3],
		}, nil

	} else {
		// We should never get here
		Logger.Errorf("Unable to determine pattern of file %s\n", fileName)
		return &MediaFile{}, errors.New("failed to create MediaFile")
	}
}
