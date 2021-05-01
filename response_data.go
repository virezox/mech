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
      Status          string `json:"status"`
      Reason          string `json:"reason"`
      PlayableInEmbed bool   `json:"playableInEmbed"`
      ContextParams   string `json:"contextParams"`
   } `json:"playabilityStatus"`
   StreamingData struct {
      ExpiresInSeconds string   `json:"expiresInSeconds"`
      Formats          []Format `json:"formats"`
      AdaptiveFormats  []Format `json:"adaptiveFormats"`
      DashManifestURL  string   `json:"dashManifestUrl"`
      HlsManifestURL   string   `json:"hlsManifestUrl"`
   } `json:"streamingData"`
   VideoDetails struct {
      VideoID          string `json:"videoId"`
      Title            string `json:"title"`
      LengthSeconds    string `json:"lengthSeconds"`
      ChannelID        string `json:"channelId"`
      IsOwnerViewing   bool   `json:"isOwnerViewing"`
      ShortDescription string `json:"shortDescription"`
      IsCrawlable      bool   `json:"isCrawlable"`
      Thumbnail        struct {
         Thumbnails []Thumbnail `json:"thumbnails"`
      } `json:"thumbnail"`
      AverageRating     float64 `json:"averageRating"`
      AllowRatings      bool    `json:"allowRatings"`
      ViewCount         string  `json:"viewCount"`
      Author            string  `json:"author"`
      IsPrivate         bool    `json:"isPrivate"`
      IsUnpluggedCorpus bool    `json:"isUnpluggedCorpus"`
      IsLiveContent     bool    `json:"isLiveContent"`
   } `json:"videoDetails"`
   Microformat struct {
      PlayerMicroformatRenderer struct {
         Thumbnail struct {
            Thumbnails []struct {
               URL    string `json:"url"`
               Width  int    `json:"width"`
               Height int    `json:"height"`
            } `json:"thumbnails"`
         } `json:"thumbnail"`
         Embed struct {
            IframeURL      string `json:"iframeUrl"`
            FlashURL       string `json:"flashUrl"`
            Width          int    `json:"width"`
            Height         int    `json:"height"`
            FlashSecureURL string `json:"flashSecureUrl"`
         } `json:"embed"`
         Title struct {
            SimpleText string `json:"simpleText"`
         } `json:"title"`
         Description struct {
            SimpleText string `json:"simpleText"`
         } `json:"description"`
         LengthSeconds      string   `json:"lengthSeconds"`
         OwnerProfileURL    string   `json:"ownerProfileUrl"`
         ExternalChannelID  string   `json:"externalChannelId"`
         AvailableCountries []string `json:"availableCountries"`
         IsUnlisted         bool     `json:"isUnlisted"`
         HasYpcMetadata     bool     `json:"hasYpcMetadata"`
         ViewCount          string   `json:"viewCount"`
         Category           string   `json:"category"`
         PublishDate        string   `json:"publishDate"`
         OwnerChannelName   string   `json:"ownerChannelName"`
         UploadDate         string   `json:"uploadDate"`
      } `json:"playerMicroformatRenderer"`
   } `json:"microformat"`
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

type Thumbnails []Thumbnail

type Thumbnail struct {
	URL    string
	Width  uint
	Height uint
}
