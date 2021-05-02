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

const API = "https://www.youtube.com/get_video_info"

// httpGetBodyBytes reads the whole HTTP body and returns it
func httpGetBodyBytes(url string) ([]byte, error) {
   req, err := http.NewRequest(http.MethodGet, url, nil)
   if err != nil { return nil, err }
   println(req.Method, url)
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
   // Circumvent age restriction to pretend access through googleapis.com
   eurl := "https://youtube.googleapis.com/v/" + id
   body, err := httpGetBodyBytes("https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
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

////////////////////////////////////////////////////////////////////////////////


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

// newVideo fetches video metadata
func newVideo(id string) (oldVideo, error) {
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
