package main

/*
#cgo LDFLAGS: -framework CoreFoundation -L/usr/local/lib -L/usr/local/Cellar/libmediainfo/24.05/lib -lmediainfo
#include <stdlib.h>
#include "wrapper/mediainfo.c"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// Define a Go struct to hold the pointer to MediaInfoWrapper
type MediaInfoWrapper struct {
	pointer unsafe.Pointer
}

// Create and initialize an instance of the MediaInfoWrapper
func NewMediaInfoWrapper(fileName string) (*MediaInfoWrapper, error) {
	Logger.WithField(
		"component",
		"mediainfowrapper").Debug("Initializing MediaInfoWrapper")

	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	cmi := C.newMediaInfoWrapper(cFileName)
	if cmi == nil {
		errMsg := "failed to create MediaInfoWrapper"
		Logger.WithField(
			"component",
			"mediainfowrapper").Error(errMsg)
		return &MediaInfoWrapper{}, errors.New(errMsg)
	}

	Logger.WithField(
		"component",
		"mediainfowrapper").Debug("MediaInfoWrapper Initialized")
	return &MediaInfoWrapper{pointer: unsafe.Pointer(cmi)}, nil
}

// Prints all of the properties for a media file
func (mi *MediaInfoWrapper) listProperties() {
	info := C.readMediaFile((*C.struct_MediaInfoWrapper)(mi.pointer))
	Logger.WithField(
		"component",
		"mediainfowrapper").Debug(C.GoString(info))
}

// Get the number of audio streams in a media file
func (mi *MediaInfoWrapper) getAudioCount() int {
	audioCount := C.getAudioCount((*C.struct_MediaInfoWrapper)(mi.pointer))
	return int(audioCount)
}

// Get the number of channels for a specific audio index
func (mi *MediaInfoWrapper) getAudioChannels(index int) int {
	audioChannels := C.getAudioChannels((*C.struct_MediaInfoWrapper)(mi.pointer), C.int(index))
	return int(audioChannels)
}

// Get the bitrate for a specific audio index
func (mi *MediaInfoWrapper) getAudioBitRate(index int) int {
	audioBitRate := C.getAudioBitRate((*C.struct_MediaInfoWrapper)(mi.pointer), C.int(index))
	return int(audioBitRate)
}
