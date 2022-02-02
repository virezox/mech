package instagram

import (
   "encoding/json"
   "encoding/xml"
   "github.com/89z/format"
   "net/http"
   "strings"
)

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

func (i Info) URL() (string, error) {
   var (
      addr string
      max int
   )
   switch {
   case i.Media_Type == 1:
      for _, can := range i.Image_Versions2.Candidates {
         if can.Height > max {
            addr = can.URL
            max = can.Height
         }
      }
   case i.Video_DASH_Manifest != "":
      var manifest mpd
      err := xml.Unmarshal([]byte(i.Video_DASH_Manifest), &manifest)
      if err != nil {
         return "", err
      }
      for _, rep := range manifest.Period.AdaptationSet.Representation {
         if rep.Bandwidth > max {
            addr = rep.BaseURL
            max = rep.Bandwidth
         }
      }
   case i.Media_Type == 2:
      // Type:101 Bandwidth:211,754
      // Type:102 Bandwidth:541,145
      // Type:103 Bandwidth:541,145
      for _, ver := range i.Video_Versions {
         if ver.Type > max {
            addr = ver.URL
            max = ver.Type
         }
      }
   }
   return addr, nil
}

var logLevel format.LogLevel

type MediaItem struct {
   Info
   Carousel_Media []Info
   Like_Count int64
}

func (m MediaItem) Infos() []Info {
   if m.Media_Type == 8 {
      return m.Carousel_Media
   }
   return []Info{m.Info}
}

type Info struct {
   Image_Versions2 struct {
      Candidates []struct {
         Width int
         Height int
         URL string
      }
   }
   Media_Type int
   Video_DASH_Manifest string
   Video_Versions []struct {
      Type int
      Width int
      Height int
      URL string
   }
}

type mpd struct {
   Period struct {
      AdaptationSet struct {
         Representation []struct {
            Width int `xml:"width,attr"`
            Height int `xml:"height,attr"`
            Bandwidth int `xml:"bandwidth,attr"`
            BaseURL string
         }
      }
   }
}
