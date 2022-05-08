package bandcamp

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/xml"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

const (
   JPEG = iota
   PNG
)

var LogLevel format.LogLevel

type Image struct {
   Crop bool
   Format int
   Height int
   ID int64
   Width int
}

var Images = []Image{
   {ID:0, Width:1500, Height:1500, Format:JPEG},
   {ID:1, Width:1500, Height:1500, Format:PNG},
   {ID:2, Width:350, Height:350, Format:JPEG},
   {ID:3, Width:100, Height:100, Format:JPEG},
   {ID:4, Width:300, Height:300, Format:JPEG},
   {ID:5, Width:700, Height:700, Format:JPEG},
   {ID:6, Width:100, Height:100, Format:JPEG},
   {ID:7, Width:150, Height:150, Format:JPEG},
   {ID:8, Width:124, Height:124, Format:JPEG},
   {ID:9, Width:210, Height:210, Format:JPEG},
   {ID:10, Width:1200, Height:1200, Format:JPEG},
   {ID:11, Width:172, Height:172, Format:JPEG},
   {ID:12, Width:138, Height:138, Format:JPEG},
   {ID:13, Width:380, Height:380, Format:JPEG},
   {ID:14, Width:368, Height:368, Format:JPEG},
   {ID:15, Width:135, Height:135, Format:JPEG},
   {ID:16, Width:700, Height:700, Format:JPEG},
   {ID:20, Width:1024, Height:1024, Format:JPEG},
   {ID:21, Width:120, Height:120, Format:JPEG},
   {ID:22, Width:25, Height:25, Format:JPEG},
   {ID:23, Width:300, Height:300, Format:JPEG},
   {ID:24, Width:300, Height:300, Format:JPEG},
   {ID:25, Width:700, Height:700, Format:JPEG},
   {ID:26, Width:800, Height:600, Format:JPEG, Crop:true},
   {ID:27, Width:715, Height:402, Format:JPEG, Crop:true},
   {ID:28, Width:768, Height:432, Format:JPEG, Crop:true},
   {ID:29, Width:100, Height:75, Format:JPEG, Crop:true},
   {ID:31, Width:1024, Height:1024, Format:PNG},
   {ID:32, Width:380, Height:285, Format:JPEG, Crop:true},
   {ID:33, Width:368, Height:276, Format:JPEG, Crop:true},
   {ID:36, Width:400, Height:300, Format:JPEG, Crop:true},
   {ID:37, Width:168, Height:126, Format:JPEG, Crop:true},
   {ID:38, Width:144, Height:108, Format:JPEG, Crop:true},
   {ID:41, Width:210, Height:210, Format:JPEG},
   {ID:42, Width:50, Height:50, Format:JPEG},
   {ID:43, Width:100, Height:100, Format:JPEG},
   {ID:44, Width:200, Height:200, Format:JPEG},
   {ID:50, Width:140, Height:140, Format:JPEG},
   {ID:65, Width:700, Height:700, Format:JPEG},
   {ID:66, Width:1200, Height:1200, Format:JPEG},
   {ID:67, Width:350, Height:350, Format:JPEG},
   {ID:68, Width:210, Height:210, Format:JPEG},
   {ID:69, Width:700, Height:700, Format:JPEG},
}

// Extension is optional.
func (i Image) URL(artID int64) string {
   buf := []byte("http://f4.bcbits.com/img/a")
   buf = strconv.AppendInt(buf, artID, 10)
   buf = append(buf, '_')
   buf = strconv.AppendInt(buf, i.ID, 10)
   return string(buf)
}

func (t Track) Base() string {
   return mech.Clean(t.Band_Name + "-" + t.Title)
}

func (t Track) String() string {
   var buf []byte
   buf = append(buf, "Num:"...)
   buf = strconv.AppendInt(buf, t.Track_Num, 10)
   buf = append(buf, " Title:"...)
   buf = append(buf, t.Title...)
   buf = append(buf, " Band:"...)
   buf = append(buf, t.Band_Name...)
   if t.Streaming_URL != nil {
      buf = append(buf, " URL:"...)
      buf = append(buf, t.Streaming_URL.MP3_128...)
   }
   return string(buf)
}

func (t Tralbum) Date() time.Time {
   return time.Unix(t.Release_Date, 0)
}

type Tralbum struct {
   Art_ID int64
   Release_Date int64
   Title string
   Tralbum_Artist string
   Tracks []Track
}

type Track struct {
   Track_Num int64
   Title string
   Band_Name string
   Streaming_URL *struct {
      MP3_128 string `json:"mp3-128"`
   }
}

type Band struct {
   Name string
   Discography []Item
}

type invalidType struct {
   value string
}

func (i invalidType) Error() string {
   return "invalid type " + strconv.Quote(i.value)
}

func (i Item) Tralbum() (*Tralbum, error) {
   switch i.Item_Type {
   case "album":
      return newTralbum('a', i.Item_ID)
   case "track":
      return newTralbum('t', i.Item_ID)
   }
   return nil, invalidType{i.Item_Type}
}

func newTralbum(typ byte, id int) (*Tralbum, error) {
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
   tralb := new(Tralbum)
   if err := json.NewDecoder(res.Body).Decode(tralb); err != nil {
      return nil, err
   }
   return tralb, nil
}

type Item struct {
   Item_ID int
   Item_Type string
}

func NewBand(id int) (*Band, error) {
   req, err := http.NewRequest(
      "GET", "http://bandcamp.com/api/mobile/24/band_details", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "band_id=" + strconv.Itoa(id)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   band := new(Band)
   if err := json.NewDecoder(res.Body).Decode(band); err != nil {
      return nil, err
   }
   return band, nil
}

type Params struct {
   A_ID int
   I_ID int
   I_Type string
}

func (p Params) Tralbum() (*Tralbum, error) {
   switch p.I_Type {
   case "a":
      return newTralbum('a', p.I_ID)
   case "t":
      return newTralbum('t', p.I_ID)
   }
   return nil, invalidType{p.I_Type}
}

func NewParams(addr string) (*Params, error) {
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
   scan, err := xml.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte(`<p id="report-account-vm"`)
   scan.Scan()
   var p struct {
      DataTouReportParams []byte `xml:"data-tou-report-params,attr"`
   }
   if err := scan.Decode(&p); err != nil {
      return nil, err
   }
   param := new(Params)
   if err := json.Unmarshal(p.DataTouReportParams, param); err != nil {
      return nil, err
   }
   return param, nil
}
