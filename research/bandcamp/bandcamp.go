package bandcamp

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/net"
   "html"
   "net/http"
)

var LogLevel format.LogLevel

func NewDataTralbum(addr string) (*DataTralbum, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data := new(DataTralbum)
   for _, node := range net.ReadHTML(res.Body, "script") {
      buf, ok := node.Attr["data-tralbum"]
      if ok {
         buf = html.UnescapeString(buf)
         err := json.Unmarshal([]byte(buf), data)
         if err != nil {
            return nil, err
         }
         break
      }
   }
   return data, nil
}

type DataTralbum struct {
   AlbumRelease []struct {
      MusicReleaseFormat string
   }
   Album_Release_Date string // 20 Jan 2017 00:00:00 GMT
   Art_ID int
   Artist string
   Current struct {
      Title string
   }
   ID int
   TrackInfo []TrackInfo
}

type TrackInfo struct {
   Title string
   File map[string]string
}
