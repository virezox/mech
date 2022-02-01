package instagram

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

var logLevel format.LogLevel

type Info struct {
   Media_Type int
   Image_Versions2 struct {
      Candidates []struct {
         Width int
         Height int
         URL string
      }
   }
   Video_Versions []struct {
      Type int
      Width int
      Height int
      URL string
   }
}

func (i Info) URL() string {
   var addr string
   switch i.Media_Type {
   case 1:
      var height int
      for _, can := range i.Image_Versions2.Candidates {
         if can.Height > height {
            addr = can.URL
            height = can.Height
         }
      }
   case 2:
      for _, ver := range i.Video_Versions {
         // Type:101 Bitrate:372 kb/s
         // Type:102 Bitrate:567 kb/s
         if ver.Type == 102 {
            addr = ver.URL
         }
      }
   }
   return addr
}

type MediaItem struct {
   Info
   Carousel_Media []Info
   Like_Count int64
}

func MediaItems(shortcode string) ([]MediaItem, error) {
   var str strings.Builder
   str.WriteString("https://www.instagram.com/p/")
   str.WriteString(shortcode)
   str.WriteByte('/')
   req, err := http.NewRequest("GET", str.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {auth},
      "User-Agent": {"Instagram 214.1.0.29.120 Android"},
   }
   req.URL.RawQuery = "__a=1"
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var info struct {
      Items []MediaItem
   }
   if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
      return nil, err
   }
   return info.Items, nil
}

func (m MediaItem) Infos() []Info {
   if m.Media_Type == 8 {
      return m.Carousel_Media
   }
   return []Info{m.Info}
}
