package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for the correct number of arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	mf, err := NewMediaFile(os.Args[1])
	if err != nil {
		Logger.Errorf("%v\n", err)
		os.Exit(-1)
	}
	Logger.Debugf("%+v\n", mf)

	/*
		mi := GetInstance(os.Args[1])
		br := mi.getAudioBitRate(0)
		Logger.Infof("BitRate for Audio 0: %d\n", br)
	*/
}
