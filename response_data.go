package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const API = "https://www.youtube.com/get_video_info"

type oldVideo struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         Description struct {
            SimpleText string
         }
         PublishDate string
         Title struct {
            SimpleText string
         }
         ViewCount int `json:",string"`
      }
   }
   StreamingData struct {
      DashManifestURL string
   }
}

// NewVideo fetches video metadata
func NewVideo(id string) (oldVideo, error) {
   req, err := http.NewRequest(http.MethodGet, API, nil)
   if err != nil {
      return oldVideo{}, err
   }
   val := req.URL.Query()
   val.Set("video_id", id)
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Range", "bytes=0-")
   res, err := new(http.Client).Do(req)
   if err != nil {
      return oldVideo{}, err
   }
   defer res.Body.Close()
   switch res.StatusCode {
   case http.StatusOK, http.StatusPartialContent:
   default:
      return oldVideo{}, fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   req.URL.RawQuery = buf.String()
   play := req.URL.Query().Get("player_response")
   buf = bytes.NewBufferString(play)
   var vid oldVideo
   err = json.NewDecoder(buf).Decode(&vid)
   if err != nil {
      return oldVideo{}, err
   }
   return vid, nil
}

func (v oldVideo) Description() string {
   return v.Microformat.PlayerMicroformatRenderer.Description.SimpleText
}

func (v oldVideo) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v oldVideo) Title() string {
   return v.Microformat.PlayerMicroformatRenderer.Title.SimpleText
}

func (v oldVideo) ViewCount() int {
   return v.Microformat.PlayerMicroformatRenderer.ViewCount
}

type playerResponseData struct {
   PlayabilityStatus struct {
      Status          string
      Reason          string
      PlayableInEmbed bool
      ContextParams   string
   }
   StreamingData struct {
      ExpiresInSeconds string
      Formats          []Format
      AdaptiveFormats  []Format
   }
   VideoDetails struct {
      VideoID          string
      Title            string
      LengthSeconds    string
      ChannelID        string
      IsOwnerViewing   bool
      ShortDescription string
      IsCrawlable      bool
      Thumbnail        struct {
         Thumbnails []Thumbnail
      }
      AverageRating     float64
      AllowRatings      bool
      ViewCount         string
      Author            string
      IsPrivate         bool
      IsUnpluggedCorpus bool
      IsLiveContent     bool
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         Thumbnail struct {
            Thumbnails []struct {
               URL    string
               Width  int
               Height int
            }
         }
         Title struct {
            SimpleText string
         }
         Description struct {
            SimpleText string
         }
         LengthSeconds      string
         OwnerProfileURL    string
         ExternalChannelID  string
         AvailableCountries []string
         IsUnlisted         bool
         HasYpcMetadata     bool
         ViewCount          string
         Category           string
         PublishDate        string
         OwnerChannelName   string
         UploadDate         string
      }
   }
}

type Format struct {
   ItagNo           int    `json:"itag"`
   URL              string `json:"url"`
   MimeType         string `json:"mimeType"`
   Quality          string `json:"quality"`
   Cipher           string `json:"signatureCipher"`
   Bitrate          int    `json:"bitrate"`
   FPS              int    `json:"fps"`
   Width            int    `json:"width"`
   Height           int    `json:"height"`
   LastModified     string `json:"lastModified"`
   ContentLength    string `json:"contentLength"`
   QualityLabel     string `json:"qualityLabel"`
   ProjectionType   string `json:"projectionType"`
   AverageBitrate   int    `json:"averageBitrate"`
   AudioQuality     string `json:"audioQuality"`
   ApproxDurationMs string `json:"approxDurationMs"`
   AudioSampleRate  string `json:"audioSampleRate"`
   AudioChannels    int    `json:"audioChannels"`
   // InitRange is only available for adaptive formats
   InitRange *struct {
      Start string `json:"start"`
      End   string `json:"end"`
   } `json:"initRange"`
   // IndexRange is only available for adaptive formats
   IndexRange *struct {
      Start string `json:"start"`
      End   string `json:"end"`
   } `json:"indexRange"`
}

type Thumbnail struct {
	URL    string
	Width  uint
	Height uint
}
