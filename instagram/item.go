package instagram

import (
   "encoding/json"
   "encoding/xml"
   "net/http"
   "strings"
   "time"
)

func (i Item) Format() (string, error) {
   var buf []byte
   buf = append(buf, "Taken: "...)
   buf = append(buf, i.Time().String()...)
   buf = append(buf, "\nUser: "...)
   buf = append(buf, i.User.Username...)
   buf = append(buf, "\nCaption: "...)
   buf = append(buf, i.Caption.Text...)
   for _, med := range i.GetItemMedia() {
      addrs, err := med.URLs()
      if err != nil {
         return "", err
      }
      for _, addr := range addrs {
         buf = append(buf, "\nURL: "...)
         buf = append(buf, addr...)
      }
   }
   return string(buf), nil
}

func (i Item) Time() time.Time {
   return time.Unix(i.Taken_At, 0)
}

func (i Item) GetItemMedia() []ItemMedia {
   if i.Media_Type == 8 {
      return i.Carousel_Media
   }
   return []ItemMedia{i.ItemMedia}
}

func (l Login) Items(shortcode string) ([]Item, error) {
   var buf strings.Builder
   buf.WriteString("https://www.instagram.com/p/")
   buf.WriteString(shortcode)
   buf.WriteByte('/')
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {l.Authorization},
      "User-Agent": {Android.String()},
   }
   req.URL.RawQuery = "__a=1"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var post struct {
      Items []Item
   }
   if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
      return nil, err
   }
   return post.Items, nil
}

func (i ItemMedia) URLs() ([]string, error) {
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

type Item struct {
   Caption struct {
      Text string
   }
   Carousel_Media []ItemMedia
   ItemMedia
   Taken_At int64
   User struct {
      Username string
   }
}

type ItemMedia struct {
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
