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
      Candidates []Version
   }
   Video_Versions []Version
}

func (i Info) Version() (*Version, error) {
   var dst Version
   switch i.Media_Type {
   case 1:
      for _, src := range i.Image_Versions2.Candidates {
         if src.Height > dst.Height {
            dst = src
         }
      }
   case 2:
      done := make(map[string]bool)
      var length int64
      for _, src := range i.Video_Versions {
         if !done[src.URL] {
            done[src.URL] = true
            if src.Height > dst.Height {
               dst = src
            } else if src.Height == dst.Height {
               req, err := http.NewRequest("HEAD", src.URL, nil)
               if err != nil {
                  return nil, err
               }
               logLevel.Dump(req)
               res, err := new(http.Transport).RoundTrip(req)
               if err != nil {
                  return nil, err
               }
               if res.ContentLength > length {
                  dst = src
               }
            }
         }
      }
   }
   return &dst, nil
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

type Version struct {
   Width int
   Height int
   URL string
}
