package main

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

const partLength = 1_000_000

func (f format) download() error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   var (
      begin = time.Now()
      content int64
   )
   for content < f.ContentLength {
      req.Header.Set(
         "Range", fmt.Sprintf("bytes=%v-%v", content, content+partLength-1),
      )
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.ReadAll(res.Body); err != nil {
         return err
      }
      content += partLength
   }
}

func main() {
   play, err := newPlayer()
   if err != nil {
      panic(err)
   }
   for _, form := range play.StreamingData.AdaptiveFormats {
      if form.ContentLength >= 20_000_000 {
         err := form.download()
         if err != nil {
            panic(err)
         }
         break
      }
   }
}

func newPlayer() (*player, error) {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Host"] = []string{"www.youtube.com"}
   req.Header["X-Goog-Api-Key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type format struct {
   ContentLength int `json:"contentLength,string"`
   MimeType string
   URL string
}

type player struct {
   StreamingData struct {
      AdaptiveFormats []format
   }
}

var body = strings.NewReader(`{
 "videoId": "HY6kFlstcFE",
 "context": {
  "client": {
   "clientName": "TVHTML5_SIMPLY_EMBEDDED_PLAYER",
   "clientVersion": "2.0"
  }
 }
}
`)
