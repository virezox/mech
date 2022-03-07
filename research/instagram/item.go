package instagram

import (
   "encoding/json"
   "encoding/xml"
   "net/http"
   "strings"
   "time"
)

type VideoVersion struct {
   Type int
   Width int
   Height int
   URL string
}

type ImageVersion struct {
   Candidates []struct {
      Width int
      Height int
      URL string
   }
}

type Item struct {
   Caption struct {
      Text string
   }
   Carousel_Media []struct {
      Media_Type int
      Video_DASH_Manifest string
      Video_Versions []VideoVersion
      Image_Versions2 ImageVersion
   }
   Image_Versions2 ImageVersion
   Media_Type int
   Taken_At int64
   User struct {
      Username string
   }
   Video_DASH_Manifest string
   Video_Versions []VideoVersion
}

func (i Item) GetItemMedia() []ItemMedia {
   if i.Media_Type == 8 {
      return i.Carousel_Media
   }
   return []ItemMedia{i.ItemMedia}
}

func (i Item) URLs() ([]string, error) {
   var addrs []string
   switch i.Media_Type {
   case 1:
      var max int
      for _, can := range i.Image_Versions2.Candidates {
         if can.Height > max {
            addrs = []string{can.URL}
            max = can.Height
         }
      }
   case 2:
      if i.Video_DASH_Manifest != "" {
         var manifest mpd
         err := xml.Unmarshal([]byte(i.Video_DASH_Manifest), &manifest)
         if err != nil {
            return nil, err
         }
         for _, ada := range manifest.Period.AdaptationSet {
            var (
               addr string
               max int
            )
            for _, rep := range ada.Representation {
               if rep.Bandwidth > max {
                  addr = rep.BaseURL
                  max = rep.Bandwidth
               }
            }
            addrs = append(addrs, addr)
         }
      } else {
         var max int
         for _, ver := range i.Video_Versions {
            if ver.Type > max {
               addrs = []string{ver.URL}
               max = ver.Type
            }
         }
      }
   }
   return addrs, nil
}
