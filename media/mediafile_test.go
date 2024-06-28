package media

import (
	"testing"
)

func TestNewMediaFile(t *testing.T) {
	t.Run("test movies", func(t *testing.T) {

		t.Run("file is a movie", func(t *testing.T) {
			mf, err := NewMediaFile("Example Movie Title (2024) Orig.mkv")
			if err != nil {
				t.Error(err)
				return
			}

			got := mf.toString()
			want := `{"Title":"Example Movie Title","Season":"","Episode":"","Year":"2024","Extension":"mkv"}`

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})

		t.Run("movie file missing orig suffix", func(t *testing.T) {
			_, err := NewMediaFile("Example Movie Title (2024).mkv")
			if err == nil { // should fail
				t.Error(err)
				return
			}
		})

		t.Run("movie file missing year", func(t *testing.T) {
			_, err := NewMediaFile("Example Movie Title Orig.mkv")
			if err == nil { // should fail
				t.Error(err)
				return
			}
		})

		t.Run("movie file has a hyphen", func(t *testing.T) {
			mf, err := NewMediaFile("Example - Movie Title (2024) Orig.mkv")
			if err != nil {
				t.Error(err)
				return
			}

			got := mf.toString()
			want := `{"Title":"Example - Movie Title","Season":"","Episode":"","Year":"2024","Extension":"mkv"}`

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})
	})

	/* Test TV Shows */
	t.Run("test tv shows", func(t *testing.T) {
		t.Run("file is a tv show", func(t *testing.T) {
			mf, err := NewMediaFile("TV Show - S01E01 - The Episode Name (2024) Orig.mkv")
			if err != nil {
				t.Error(err)
				return
			}

			got := mf.toString()
			want := `{"Title":"TV Show","Season":"S01E01","Episode":"The Episode Name","Year":"2024","Extension":"mkv"}`

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})

		t.Run("tv show name wrong season format", func(t *testing.T) {
			_, err := NewMediaFile("TV Show - S01 - The Episode Name (2024) Orig.mkv")
			if err == nil { // should fail
				t.Error(err)
				return
			}
		})

		t.Run("tv show name has an extra hyphen", func(t *testing.T) {
			_, err := NewMediaFile("TV - Show - S01 - The Episode Name (2024) Orig.mkv")
			if err == nil { // should fail
				t.Error(err)
				return
			}
		})

		t.Run("tv show name missing year", func(t *testing.T) {
			_, err := NewMediaFile("TV Show - S01 - The Episode Name Orig.mkv")
			if err == nil { // should fail
				t.Error(err)
				return
			}
		})
	})
}
