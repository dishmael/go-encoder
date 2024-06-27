package main

import (
	"fmt"
	log "go-encoder/logger"
	"go-encoder/test"
	"os"
)

/*
func encode() {
	cmd := exec.Command(
		"ffmpeg", "-y",
		"-i", os.Args[1],
		"-map", "0",
		"-c", "copy",
		"test.mkv",
	)

	// Create strings.Builder for combined stdout and stderr
	//var outBuilder strings.Builder
	//multiWriter := io.MultiWriter(&outBuilder, os.Stdout)
	multiWriter := io.MultiWriter(os.Stdout)

	cmd.Stdout = multiWriter
	cmd.Stderr = multiWriter

	// Run the command
	cmd.Run()

	// Print the combined output and error messages
	//fmt.Println(outBuilder.String())
}
*/

func main() {
	// Check for the correct number of arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	log.Logger.Debug("Testing!")

	/*
		mf, err := NewMediaFile(os.Args[1])
		if err != nil {
			Logger.Errorf("%v\n", err)
			os.Exit(-1)
		}
		Logger.Debugf("%+v\n", mf)
	*/

	/*
		mi := GetInstance(os.Args[1])
		br := mi.getAudioBitRate(0)
		Logger.Infof("BitRate for Audio 0: %d\n", br)
	*/

	t := test.Test{}
	t.SayHello()
}
