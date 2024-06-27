#include <stdio.h>
#include <stdlib.h>  // For malloc and free
#include <MediaInfoDLL/MediaInfoDLL_Static.h>
#include <stdbool.h>

struct MediaInfoWrapper {
  void* handle;
};

// Constructor
struct MediaInfoWrapper* newMediaInfoWrapper(const char* fileName) {
  struct MediaInfoWrapper* wrapper = malloc(sizeof(struct MediaInfoWrapper));
  if (wrapper) {
    wrapper->handle = MediaInfo_New();
    
    if (!wrapper->handle) {
      free(wrapper);
      return NULL;
    }

    // Set the character set to UTF-8 and open the media file for reading
    MediaInfo_Option(wrapper->handle, "CharSet", "UTF-8");
    MediaInfo_Open(wrapper->handle, fileName);
  }
  return wrapper;
}

// Function to read the media file
const char* readMediaFile(struct MediaInfoWrapper* self) {
  return MediaInfo_Inform(self->handle, 0);
}

// Function to return a general property
const char* getGeneralProperty(struct MediaInfoWrapper* self, const char* property) {
  return MediaInfo_Get(self->handle, MediaInfo_Stream_General, 0, property, MediaInfo_Info_Text, MediaInfo_Info_Name);
}

// Function to return a video property
const char* getVideoProperty(struct MediaInfoWrapper* self, const char* property) {
  return MediaInfo_Get(self->handle, MediaInfo_Stream_Video, 0, property, MediaInfo_Info_Text, MediaInfo_Info_Name);
}

// Function to return a property for a specific audio stream (audioIndex)
const char* getAudioProperty(struct MediaInfoWrapper* self, const char* property, int audioIndex) {
  return MediaInfo_Get(self->handle, MediaInfo_Stream_Audio, audioIndex, property, MediaInfo_Info_Text, MediaInfo_Info_Name);
}

// Function to return a property for a specific text stream (textIndex)
const char* getTextProperty(struct MediaInfoWrapper* self, const char* property, int textIndex) {
  return MediaInfo_Get(self->handle, MediaInfo_Stream_Text, textIndex, property, MediaInfo_Info_Text, MediaInfo_Info_Name);
}

// Destructor
void destroyMediaInfoWrapper(struct MediaInfoWrapper **self) {
  if (self && *self) {
    MediaInfo_Close((*self)->handle); // Make sure to close the handle
    MediaInfo_Delete((*self)->handle);
    free(*self);
    *self = NULL;
  }
}

/*
int main() {
  // Example usage
  struct MediaInfoWrapper* wrapper = newMediaInfoWrapper("example.mp4");
  if (wrapper) {
    const char* info = readMediaFile(wrapper);
    printf("Media Info: %s\n", info);

    const char* duration = getGeneralProperty(wrapper, "Duration");
    printf("Duration: %s\n", duration);

    destroyMediaInfoWrapper(&wrapper);
  } else {
    printf("Failed to open media file.\n");
  }

  return 0;
}
*/