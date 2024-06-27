package media

/*
#cgo LDFLAGS: -framework CoreFoundation -L . -L/usr/local/lib -L/usr/local/Cellar/libmediainfo/24.05/lib -lmediainfo
#include <stdlib.h>
#include "../wrapper/mediainfo.c"
*/
import "C"

import (
	"errors"
	log "go-encoder/logger"
	"unsafe"
)

// Define a Go struct to hold the pointer to MediaInfoWrapper
type MediaInfoWrapper struct {
	pointer unsafe.Pointer
}

// Create and initialize an instance of the MediaInfoWrapper
func NewMediaInfoWrapper(fileName string) (*MediaInfoWrapper, error) {
	log.Logger.WithField(
		"component",
		"mediainfowrapper").Debug("Initializing MediaInfoWrapper")

	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	cmi := C.newMediaInfoWrapper(cFileName)
	if cmi == nil {
		errMsg := "failed to create MediaInfoWrapper"
		log.Logger.WithField(
			"component",
			"mediainfowrapper").Error(errMsg)
		return &MediaInfoWrapper{}, errors.New(errMsg)
	}

	log.Logger.WithField(
		"component",
		"mediainfowrapper").Debug("MediaInfoWrapper Initialized")
	return &MediaInfoWrapper{pointer: unsafe.Pointer(cmi)}, nil
}

// Prints all of the properties for a media file
func (mi *MediaInfoWrapper) PrintProperties() {
	// no need to free info, it points to a constant string (cont char*)
	info := C.readMediaFile((*C.struct_MediaInfoWrapper)(mi.pointer))
	log.Logger.WithField(
		"component",
		"mediainfowrapper").Debug(C.GoString(info))
}

// Generic function to retrieve a General property
func (mi *MediaInfoWrapper) GetGeneralProperty(property string) string {
	cProperty := C.CString(property)
	defer C.free(unsafe.Pointer(cProperty))

	// no need to free cProps, it points to a constant string (cont char*)
	cProp := C.getGeneralProperty((*C.struct_MediaInfoWrapper)(mi.pointer), cProperty)
	log.Logger.WithField(
		"component",
		"mediainfowrapper").Debugf("%s = %s", property, C.GoString(cProp))

	return C.GoString(cProp)
}

// Generic function to retrieve a property from a specific Video stream (index)
func (mi *MediaInfoWrapper) GetVideoProperty(property string, index int) string {
	cProperty := C.CString(property)
	defer C.free(unsafe.Pointer(cProperty))

	// no need to free cProps, it points to a constant string (cont char*)
	cProp := C.getVideoProperty((*C.struct_MediaInfoWrapper)(mi.pointer), cProperty)
	log.Logger.WithField(
		"component",
		"mediainfowrapper").Debugf("%s = %s", property, C.GoString(cProp))

	return C.GoString(cProp)
}

// Generic function to retrieve a property from a specific Audio stream (index)
func (mi *MediaInfoWrapper) GetAudioProperty(property string, index int) string {
	cProperty := C.CString(property)
	defer C.free(unsafe.Pointer(cProperty))

	// no need to free cProps, it points to a constant string (cont char*)
	cProp := C.getAudioProperty((*C.struct_MediaInfoWrapper)(mi.pointer), cProperty, C.int(index))
	log.Logger.WithField(
		"component",
		"mediainfowrapper").Debugf("%s = %s", property, C.GoString(cProp))

	return C.GoString(cProp)
}

// Generic function to retrieve a property from a specific Text stream (index)
func (mi *MediaInfoWrapper) GetTextProperty(property string, index int) string {
	cProperty := C.CString(property)
	defer C.free(unsafe.Pointer(cProperty))

	// no need to free cProps, it points to a constant string (cont char*)
	cProp := C.getTextProperty((*C.struct_MediaInfoWrapper)(mi.pointer), cProperty, C.int(index))
	log.Logger.WithField(
		"component",
		"mediainfowrapper").Debugf("%s = %s", property, C.GoString(cProp))

	return C.GoString(cProp)
}
