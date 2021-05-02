package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

// Client offers methods to download video metadata and video streams.
type Client struct {
   // decipherOpsCache cache decipher operations
   decipherOpsCache DecipherOperationsCache
}

// httpGetBodyBytes reads the whole HTTP body and returns it
func (c *Client) httpGetBodyBytes(url string) ([]byte, error) {
   req, err := http.NewRequest(http.MethodGet, url, nil)
   if err != nil { return nil, err }
   resp, err := new(http.Client).Do(req)
   if err != nil { return nil, err }
   defer resp.Body.Close()
   switch resp.StatusCode {
   case http.StatusOK, http.StatusPartialContent:
   default:
      return nil, fmt.Errorf("StatusCode %v", resp.StatusCode)
   }
   return io.ReadAll(resp.Body)
}

////////////////////////////////////////////////////////////////////////////////

// GetVideo fetches video metadata
func (c *Client) GetVideo(id string) (*Video, error) {
   // Circumvent age restriction to pretend access through googleapis.com
   eurl := "https://youtube.googleapis.com/v/" + id
   body, err := c.httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
   if err != nil { return nil, err }
   v := &Video{ID: id}
   query, err := url.ParseQuery(string(body))
   if err != nil { return nil, err }
   status := query.Get("status")
   if status != "ok" {
      return nil, fmt.Errorf("status: %q reason: %q", status, query.Get("reason"))
   }
   // read the streams map
   playerResponse := query.Get("player_response")
   if playerResponse == "" {
      return nil, errors.New("no player_response found in the server's answer")
   }
   err = json.Unmarshal([]byte(playerResponse), v)
   if err != nil {
      return nil, fmt.Errorf("unable to parse player response JSON %v", err)
   }
   return v, nil
}

// GetStream returns the url for a specific format
func (c *Client) GetStream(video *Video, format *Format) (string, error) {
   if format.URL != "" { return format.URL, nil }
   cipher := format.Cipher
   if cipher == "" { return "", ErrCipherNotFound }
   return c.decipherURL(video.ID, cipher)
}


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

func (v Video) Description() string {
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

type Video struct {
   // FIXME kill this
   ID string
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
   ApproxDurationMs string
   AudioChannels    int
   AudioQuality     string
   AudioSampleRate  string
   AverageBitrate   int
   Bitrate          int
   Cipher           string `json:"signatureCipher"`
   ContentLength    string
   FPS              int
   Height           int
   ItagNo           int    `json:"itag"`
   LastModified     string
   MimeType         string
   ProjectionType   string
   Quality          string
   QualityLabel     string
   URL              string
   Width            int
}

