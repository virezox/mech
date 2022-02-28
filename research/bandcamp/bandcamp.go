package bandcamp

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/net"
   "net/http"
)

var LogLevel format.LogLevel

func NewDataTralbum(addr string) error {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   for _, node := range net.ReadHTML(res.Body, "script") {
      fmt.Print(node, "\n\n")
   }
   return nil
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
