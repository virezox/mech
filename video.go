package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/url"
)

const API = "https://www.youtube.com/get_video_info"

// Client offers methods to download video metadata and video streams.
type Client struct {
   // decipherOpsCache cache decipher operations
   decipherOpsCache *simpleCache
}

// GetStream returns the url for a specific format
func (c *Client) GetStream(video *Video, format Format) (string, error) {
   if format.SignatureCipher == "" {
      return "", errors.New("cipher not found")
   }
   queryParams, err := url.ParseQuery(format.SignatureCipher)
   if err != nil { return "", err }
   if c.decipherOpsCache == nil {
      c.decipherOpsCache = new(simpleCache)
   }
   operations := c.decipherOpsCache.Get(video.ID)
   if operations == nil {
      operations, err = c.parseDecipherOps(video.ID)
      if err != nil { return "", err }
      c.decipherOpsCache.Set(video.ID, operations)
   }
   // apply operations
   bs := []byte(queryParams.Get("s"))
   for _, op := range operations {
      bs = op(bs)
   }
   return fmt.Sprintf(
      "%s&%s=%s", queryParams.Get("url"), queryParams.Get("sp"), string(bs),
   ), nil
}

// GetVideo fetches video metadata
func (c *Client) GetVideo(id string) (*Video, error) {
   req, err := http.NewRequest(http.MethodGet, API, nil)
   if err != nil { return nil, err }
   val := req.URL.Query()
   val.Set("video_id", id)
   val.Set("eurl", "https://youtube.googleapis.com/v/" + id)
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Client).Do(req)
   if err != nil { return nil, err }
   defer res.Body.Close()
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   req.URL.RawQuery = buf.String()
   play := req.URL.Query().Get("player_response")
   buf = bytes.NewBufferString(play)
   var vid Video
   err = json.NewDecoder(buf).Decode(&vid)
   if err != nil { return nil, err }
   vid.ID = id
   return &vid, nil
}

func (v Video) Description() string {
   return v.Microformat.PlayerMicroformatRenderer.Description.SimpleText
}

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string {
   return v.Microformat.PlayerMicroformatRenderer.Title.SimpleText
}

func (v Video) ViewCount() int {
   return v.Microformat.PlayerMicroformatRenderer.ViewCount
}

type Video struct {
   // FIXME kill this
   ID string
   StreamingData struct {
      AdaptiveFormats []Format
      ExpiresInSeconds string
   }
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
}

type Format struct {
   Bitrate int
   Height int
   Itag int
   MimeType string
   SignatureCipher string
}
