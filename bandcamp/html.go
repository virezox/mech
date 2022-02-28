package bandcamp

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/net"
   "html"
   "net/http"
   "strings"
   "time"
)

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

// jonasmunk.bandcamp.com/track/altered-light
func (d DataTralbum) Date() (time.Time, error) {
   return time.Parse("02 Jan 2006 15:04:05 MST", d.Album_Release_Date)
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

func (t TrackInfo) Name(data *DataTralbum, head http.Header) (string, error) {
   ext, err := format.ExtensionByType(head.Get("Content-Type"))
   if err != nil {
      return "", err
   }
   return strings.Map(format.Clean, data.Artist + "-" + t.Title) + ext, nil
}

type TrackInfo struct {
   Title string
   File map[string]string
}

// some tracks cannot be streamed:
// schnaussandmunk.bandcamp.com/album/passage-2
func (t TrackInfo) MP3_128() (string, bool) {
   mp3, ok := t.File["mp3-128"]
   if !ok {
      return "", false
   }
   return mp3, true
}
