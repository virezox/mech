package facebook

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

var LogLevel format.LogLevel

type Image struct {
   Image struct {
      URI string
   }
}

func (i *Image) Parse() error {
   addr, err := url.Parse(i.Image.URI)
   if err != nil {
      return err
   }
   vals := addr.Query()
   for key := range vals {
      switch key {
      case
      "_nc_ht",
      "_nc_oc",
      "_nc_ohc",
      "oe",
      "oh",
      "stp":
      default:
         vals.Del(key)
      }
   }
   addr.RawQuery = vals.Encode()
   i.Image.URI = addr.String()
   return nil
}

type Video struct {
   DateCreated string
   Playable_URL_Quality_HD string
   Preferred_Thumbnail Image
   Text string
}

func NewVideo(id int64) (*Video, error) {
   req, err := http.NewRequest("GET", "https://www.facebook.com/video.php", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Accept", "text/html")
   req.URL.RawQuery = "v=" + strconv.FormatInt(id, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   scan, err := json.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   var vid Video
   scan.Split = []byte(`{"\u0040context"`)
   scan.Scan()
   if err := scan.Decode(&vid); err != nil {
      return nil, err
   }
   scan.Split = []byte(`{"__typename"`)
   scan.Scan()
   if err := scan.Decode(&vid); err != nil {
      return nil, err
   }
   scan.Split = []byte(`{"delight_ranges"`)
   scan.Scan()
   if err := scan.Decode(&vid); err != nil {
      return nil, err
   }
   return &vid, nil
}

func (v Video) String() string {
   var buf strings.Builder
   buf.WriteString("Text: ")
   buf.WriteString(v.Text)
   buf.WriteString("\nDate: ")
   buf.WriteString(v.DateCreated)
   buf.WriteString("\nImage: ")
   buf.WriteString(v.Preferred_Thumbnail.Image.URI)
   buf.WriteString("\nVideo: ")
   buf.WriteString(v.Playable_URL_Quality_HD)
   return buf.String()
}
