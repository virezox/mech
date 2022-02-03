package instagram

import (
   "encoding/xml"
   "strconv"
)

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

func (i Info) URLs() ([]string, error) {
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
         // Type:101 Bandwidth:211,754
         // Type:102 Bandwidth:541,145
         // Type:103 Bandwidth:541,145
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

type MediaItem struct {
   Info
   Carousel_Media []Info
   Like_Count int64
}

// Anonymous request
func MediaItems(shortcode string) ([]MediaItem, error) {
   return Login{}.MediaItems(shortcode)
}

func (m MediaItem) Format() (string, error) {
   buf := []byte("Like_Count: ")
   buf = strconv.AppendInt(buf, m.Like_Count, 10)
   buf = append(buf, "\nURLs: "...)
   for i, info := range m.Infos() {
      addrs, err := info.URLs()
      if err != nil {
         return "", err
      }
      if i >= 1 {
         buf = append(buf, "\n---\n"...)
      }
      for j, addr := range addrs {
         if j >= 1 {
            buf = append(buf, "\n---\n"...)
         }
         buf = append(buf, addr...)
      }
   }
   return string(buf), nil
}

func (m MediaItem) Infos() []Info {
   if m.Media_Type == 8 {
      return m.Carousel_Media
   }
   return []Info{m.Info}
}

// I noticed that even with the posts that have `video_dash_manifest`, you have
// to request with a correct User-Agent. If you use wrong agent, you will get a
// normal response, but the `video_dash_manifest` will be missing.
type UserAgent struct {
   API int64
   Brand string
   Density string
   Device string
   Instagram string
   Model string
   Platform string
   Release int64
   Resolution string
}

var Android = UserAgent{
   API: 99,
   Brand: "brand",
   Density: "density",
   Device: "device",
   Instagram: "220.0.0.16.115",
   Model: "model",
   Platform: "platform",
   Release: 9,
   Resolution: "9999x9999",
}

func (u UserAgent) String() string {
   buf := []byte("Instagram ")
   buf = append(buf, u.Instagram...)
   buf = append(buf, " Android ("...)
   buf = strconv.AppendInt(buf, u.API, 10)
   buf = append(buf, '/')
   buf = strconv.AppendInt(buf, u.Release, 10)
   buf = append(buf, "; "...)
   buf = append(buf, u.Density...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Resolution...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Brand...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Model...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Device...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Platform...)
   return string(buf)
}

type mpd struct {
   Period struct {
      AdaptationSet []struct {
         Representation []struct {
            Width int `xml:"width,attr"`
            Height int `xml:"height,attr"`
            Bandwidth int `xml:"bandwidth,attr"`
            BaseURL string
         }
      }
   }
}
