package bandcamp

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/net"
   "html"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

var Images = []Image{
   {ID:0, Width:1500, Height:1500, Ext:".jpg"},
   {ID:1, Width:1500, Height:1500, Ext:".png"},
   {ID:2, Width:350, Height:350, Ext:".jpg"},
   {ID:3, Width:100, Height:100, Ext:".jpg"},
   {ID:4, Width:300, Height:300, Ext:".jpg"},
   {ID:5, Width:700, Height:700, Ext:".jpg"},
   {ID:6, Width:100, Height:100, Ext:".jpg"},
   {ID:7, Width:150, Height:150, Ext:".jpg"},
   {ID:8, Width:124, Height:124, Ext:".jpg"},
   {ID:9, Width:210, Height:210, Ext:".jpg"},
   {ID:10, Width:1200, Height:1200, Ext:".jpg"},
   {ID:11, Width:172, Height:172, Ext:".jpg"},
   {ID:12, Width:138, Height:138, Ext:".jpg"},
   {ID:13, Width:380, Height:380, Ext:".jpg"},
   {ID:14, Width:368, Height:368, Ext:".jpg"},
   {ID:15, Width:135, Height:135, Ext:".jpg"},
   {ID:16, Width:700, Height:700, Ext:".jpg"},
   {ID:20, Width:1024, Height:1024, Ext:".jpg"},
   {ID:21, Width:120, Height:120, Ext:".jpg"},
   {ID:22, Width:25, Height:25, Ext:".jpg"},
   {ID:23, Width:300, Height:300, Ext:".jpg"},
   {ID:24, Width:300, Height:300, Ext:".jpg"},
   {ID:25, Width:700, Height:700, Ext:".jpg"},
   {ID:26, Width:800, Height:600, Ext:".jpg", Crop:true},
   {ID:27, Width:715, Height:402, Ext:".jpg", Crop:true},
   {ID:28, Width:768, Height:432, Ext:".jpg", Crop:true},
   {ID:29, Width:100, Height:75, Ext:".jpg", Crop:true},
   {ID:31, Width:1024, Height:1024, Ext:".png"},
   {ID:32, Width:380, Height:285, Ext:".jpg", Crop:true},
   {ID:33, Width:368, Height:276, Ext:".jpg", Crop:true},
   {ID:36, Width:400, Height:300, Ext:".jpg", Crop:true},
   {ID:37, Width:168, Height:126, Ext:".jpg", Crop:true},
   {ID:38, Width:144, Height:108, Ext:".jpg", Crop:true},
   {ID:41, Width:210, Height:210, Ext:".jpg"},
   {ID:42, Width:50, Height:50, Ext:".jpg"},
   {ID:43, Width:100, Height:100, Ext:".jpg"},
   {ID:44, Width:200, Height:200, Ext:".jpg"},
   {ID:50, Width:140, Height:140, Ext:".jpg"},
   {ID:65, Width:700, Height:700, Ext:".jpg"},
   {ID:66, Width:1200, Height:1200, Ext:".jpg"},
   {ID:67, Width:350, Height:350, Ext:".jpg"},
   {ID:68, Width:210, Height:210, Ext:".jpg"},
   {ID:69, Width:700, Height:700, Ext:".jpg"},
}

var LogLevel format.LogLevel

// jonasmunk.bandcamp.com/track/altered-light
func (d DataTralbum) Date() (time.Time, error) {
   return time.Parse("02 Jan 2006 15:04:05 MST", d.Album_Release_Date)
}

type Image struct {
   ID int64
   Width int
   Height int
   Ext string
   Crop bool
}

// Extension is optional.
func (i Image) Format(artID int64) string {
   buf := []byte("http://f4.bcbits.com/img/a")
   buf = strconv.AppendInt(buf, artID, 10)
   buf = append(buf, '_')
   buf = strconv.AppendInt(buf, i.ID, 10)
   return string(buf)
}

func NewTralbum(typ byte, id int) (*Tralbum, error) {
   req, err := http.NewRequest(
      "GET", "http://bandcamp.com/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "band_id": {"1"},
      "tralbum_id": {strconv.Itoa(id)},
      "tralbum_type": {string(typ)},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tra := new(Tralbum)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}

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

func (t TrackInfo) Name(data *DataTralbum, head http.Header) (string, error) {
   ext, err := format.ExtensionByType(head.Get("Content-Type"))
   if err != nil {
      return "", err
   }
   return strings.Map(format.Clean, data.Artist + "-" + t.Title) + ext, nil
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

// All fields available with Track and Album
type Tralbum struct {
   Art_ID int
   Release_Date int64
   Title string
   Tracks []struct {
      Track_Num int
      Title string
      Streaming_URL map[string]string
   }
   Tralbum_Artist string
}

