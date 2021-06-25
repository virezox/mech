package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
)

const VersionAndroid = "15.01"

type Android struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   VideoDetails `json:"videoDetails"`
}

func NewAndroid(id string) (*Android, error) {
   r, err := newPlayer(id, "ANDROID", VersionAndroid).post()
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   a := new(Android)
   json.NewDecoder(r.Body).Decode(a)
   return a, nil
}

func (a Android) NewFormat(itag int) (*Format, error) {
   for _, format := range a.StreamingData.AdaptiveFormats {
      if format.Itag == itag {
         return &format, nil
      }
   }
   return nil, fmt.Errorf("itag %v", itag)
}

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   URL string
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   ViewCount int `json:"viewCount,string"`
}

func (f Format) Write(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   var pos int64
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if res.StatusCode != http.StatusPartialContent {
         return fmt.Errorf("status %v", res.Status)
      }
      if _, err := io.Copy(w, res.Body); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}
