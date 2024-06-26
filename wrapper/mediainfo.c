#include <stdio.h>
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
  }

  // Set the character set to UTF-8 and open the media file for reading
  MediaInfo_Option(wrapper->handle, "CharSet", "UTF-8");
  MediaInfo_Open(wrapper->handle, fileName);
  return wrapper;
}

// Function to read the media file
const char* readMediaFile(struct MediaInfoWrapper* self) {
  return MediaInfo_Inform(self->handle, 0);
}

// Function to get the total number of audio streams
int getAudioCount(struct MediaInfoWrapper* self) {
  const char* aCountStr = MediaInfo_Get(self->handle, MediaInfo_Stream_General, 0, "AudioCount", MediaInfo_Info_Text, MediaInfo_Info_Name);
  int aCount = aCountStr ? atoi(aCountStr) : 0;
  return aCount;
}

// Function to get the number of channels for an audio stream
int getAudioChannels(struct MediaInfoWrapper* self, int audioIndex) {
  const char* channelCountStr = MediaInfo_Get(self->handle, MediaInfo_Stream_Audio, audioIndex, "Channels", MediaInfo_Info_Text, MediaInfo_Info_Name);
  int channelCount = channelCountStr ? atoi(channelCountStr) : 0;
  return channelCount;
}

// Function to get the number of channels for an audio stream
int getAudioBitRate(struct MediaInfoWrapper* self, int audioIndex) {
  const char* bitrateStr = MediaInfo_Get(self->handle, MediaInfo_Stream_Audio, audioIndex, "BitRate", MediaInfo_Info_Text, MediaInfo_Info_Name);
  int bitrate = bitrateStr ? atoi(bitrateStr) : 0;
  return bitrate;
}

// Destructor
void destroyMediaInfoWrapper(struct MediaInfoWrapper **self) {
  free(*self);
  *self = NULL;
}