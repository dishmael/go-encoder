package media

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMediaFile(t *testing.T) {
	tests := []struct {
		name        string
		fileName    string
		expected    *MediaFile
		expectedErr error
	}{
		{
			name:     "TV show file",
			fileName: "Show Name - S01E01 - Episode Name (2022) Orig.mkv",
			expected: &MediaFile{
				Title:     "Show Name",
				Season:    "S01E01",
				Episode:   "Episode Name",
				Year:      "2022",
				Extension: "mkv",
			},
			expectedErr: nil,
		},
		{
			name:     "Movie file",
			fileName: "Movie Name (2022) Orig.mkv",
			expected: &MediaFile{
				Title:     "Movie Name",
				Year:      "2022",
				Extension: "mkv",
			},
			expectedErr: nil,
		},
		{
			name:        "Unknown file type",
			fileName:    "Unknown File Type.mkv",
			expected:    nil,
			expectedErr: errors.New("failed to find matches"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf, err := NewMediaFile(tt.fileName)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, mf)
			}
		})
	}
}

func TestGetSegments(t *testing.T) {
	tests := []struct {
		name      string
		fileName  string
		expected  int
		expectErr bool
	}{
		{"TV show file", "Show Name - S01E01 - Episode Name (2022) Orig.mkv", 3, false},
		{"Movie file", "Movie Name (2022) Orig.mkv", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			segments := getSegments(tt.fileName)
			assert.Equal(t, tt.expected, segments)
		})
	}
}

func TestGetMediaFileType(t *testing.T) {
	tests := []struct {
		name     string
		segments int
		expected MediaFileType
	}{
		{"TV show", 5, TV},
		{"Movie", 2, MOVIE},
		{"Unknown", 0, UNKNOWN},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mft := getMediaFileType(tt.segments)
			assert.Equal(t, tt.expected, mft)
		})
	}
}

func TestGetRegExPattern(t *testing.T) {
	tests := []struct {
		name     string
		mft      MediaFileType
		expected string
	}{
		{"TV show pattern", TV, `([a-zA-Z0-9\s]+)\s-\s(S[0-9]+E[0-9]+)\s-\s([a-zA-Z0-9\s\.]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)$`},
		{"Movie pattern", MOVIE, `([a-zA-Z0-9\s\-]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)`},
		{"Unknown pattern", UNKNOWN, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pattern := getRegExPattern(tt.mft)
			assert.Equal(t, tt.expected, pattern)
		})
	}
}

func TestExtractMatches(t *testing.T) {
	tests := []struct {
		name      string
		fileName  string
		pattern   string
		expected  []string
		expectErr bool
	}{
		{
			name:     "TV show match",
			fileName: "Show Name - S01E01 - Episode Name (2022) Orig.mkv",
			pattern:  `([a-zA-Z0-9\s]+)\s-\s(\w+)\s-\s([a-zA-Z0-9\s\.]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)$`,
			expected: []string{"Show Name - S01E01 - Episode Name (2022) Orig.mkv", "Show Name", "S01E01", "Episode Name", "2022", "mkv"},
		},
		{
			name:     "Movie match",
			fileName: "Movie Name (2022) Orig.mkv",
			pattern:  `([a-zA-Z0-9\s\-]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)`,
			expected: []string{"Movie Name (2022) Orig.mkv", "Movie Name", "2022", "mkv"},
		},
		{
			name:      "No match",
			fileName:  "No Match File.mkv",
			pattern:   `([a-zA-Z0-9\s]+)\s-\s(\w+)\s-\s([a-zA-Z0-9\s\.]+)\s\(([0-9]+)\)\sOrig\.([mpkv4]+)$`,
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches, err := extractMatches(tt.fileName, tt.pattern)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, matches)
			}
		})
	}
}

func TestPopulateMediaFile(t *testing.T) {
	tests := []struct {
		name     string
		fields   []string
		mft      MediaFileType
		expected *MediaFile
	}{
		{
			name:     "Populate TV show",
			fields:   []string{"Show Name - S01E01 - Episode Name (2022) Orig.mkv", "Show Name", "S01E01", "Episode Name", "2022", "mkv"},
			mft:      TV,
			expected: &MediaFile{Title: "Show Name", Season: "S01E01", Episode: "Episode Name", Year: "2022", Extension: "mkv"},
		},
		{
			name:     "Populate Movie",
			fields:   []string{"Movie Name (2022) Orig.mkv", "Movie Name", "2022", "mkv"},
			mft:      MOVIE,
			expected: &MediaFile{Title: "Movie Name", Year: "2022", Extension: "mkv"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := populateMediaFile(tt.fields, tt.mft)
			assert.Equal(t, tt.expected, mf)
		})
	}
}
