package bandcamp

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

const (
   MobileBand = "http://bandcamp.com/api/mobile/24/band_details"
   MobileTralbum = "http://bandcamp.com/api/mobile/24/tralbum_details"
)

type Image struct {
   ID int64
   Width int
   Height int
   Ext string
   Crop bool
}

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

// Extension is optional.
func (i Image) Format(artID int64) string {
   buf := []byte("http://f4.bcbits.com/img/a")
   buf = strconv.AppendInt(buf, artID, 10)
   buf = append(buf, '_')
   buf = strconv.AppendInt(buf, i.ID, 10)
   return string(buf)
}

type Band struct {
   Artists []struct {
      ID int
      Name string
   }
   Bandcamp_URL string
   Discography []Item
}

// ID to Band. Request is anonymous.
func NewBand(id int) (*Band, error) {
   req, err := http.NewRequest("GET", MobileBand, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "band_id=" + strconv.Itoa(id)
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   ban := new(Band)
   if err := json.NewDecoder(res.Body).Decode(ban); err != nil {
      return nil, err
   }
   return ban, nil
}

type Streams map[string]string

// some tracks cannot be streamed:
// schnaussandmunk.bandcamp.com/album/passage-2
func (s Streams) MP3_128() (string, bool) {
   mp3, ok := s["mp3-128"]
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
      Streaming_URL Streams
   }
   Tralbum_Artist string
}

func (t Tralbum) Unix() time.Time {
   return time.Unix(t.Release_Date, 0)
}

type Item struct {
   Item_Type string
   Item_ID int
}

func (i Item) Tralbum() (*Tralbum, error) {
   if i.Item_Type == "" {
      return nil, mech.Invalid{"tralbum_type"}
   }
   req, err := http.NewRequest("GET", MobileTralbum, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "band_id": {"1"},
      "tralbum_id": {strconv.Itoa(i.Item_ID)},
      "tralbum_type": {i.Item_Type[:1]},
   }.Encode()
   mech.Dump(req)
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
