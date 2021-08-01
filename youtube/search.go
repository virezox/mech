package youtube

import (
   "encoding/json"
   "fmt"
   "github.com/corona10/goimagehash"
   "image"
   "image/jpeg"
   "net/http"
   "time"
)

// cache is public so user can clear
var Hash = make(map[string]*goimagehash.ImageHash)

type Item struct {
   TvMusicVideoRenderer *struct {
      LengthText struct {
         SimpleText string
      }
      NavigationEndpoint struct {
         WatchEndpoint struct {
            VideoID string
         }
      }
      PrimaryText struct {
         SimpleText string
      }
   }
}

func (i Item) Duration() (time.Duration, error) {
   l := "00:00:00"
   r := i.TvMusicVideoRenderer.LengthText.SimpleText
   l = l[:len(l) - len(r)]
   u, err := time.Parse("15:04:05", l + r)
   if err != nil {
      return 0, err
   }
   var t time.Time
   return u.AddDate(1, 0, 0).Sub(t), nil
}

func (i Item) Hash() (*goimagehash.ImageHash, error) {
   id := i.VideoID()
   if h, ok := Hash[id]; ok {
      return h, nil
   }
   p := Picture{480, 360, 270, "hqdefault", JPG}
   addr := p.Address(id)
   if Verbose {
      fmt.Println("GET", addr)
   }
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   img, err := jpeg.Decode(res.Body)
   if err != nil {
      return nil, err
   }
   x0 := (p.Width - p.SubHeight) / 2
   y0 := (p.Height - p.SubHeight) / 2
   rect := image.Rect(x0, y0, x0 + p.SubHeight, y0 + p.SubHeight)
   img = img.(*image.YCbCr).SubImage(rect)
   h, err := goimagehash.DifferenceHash(img)
   if err != nil {
      return nil, err
   }
   Hash[id] = h
   return h, nil
}

func (i Item) Title() string {
   return i.TvMusicVideoRenderer.PrimaryText.SimpleText
}

func (i Item) VideoID() string {
   return i.TvMusicVideoRenderer.NavigationEndpoint.WatchEndpoint.VideoID
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents []Item
            }
         }
      }
   }
}

func NewSearch(query string) (*Search, error) {
   var body youTubeI
   body.Context.Client = TV
   body.Params = "EgIQAQ" // type video
   body.Query = query
   res, err := post(origin + "/youtubei/v1/search", Key, body)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s := new(Search)
   if err := json.NewDecoder(res.Body).Decode(s); err != nil {
      return nil, err
   }
   return s, nil
}

func (s Search) Items() []Item {
   var items []Item
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         if item.TvMusicVideoRenderer != nil {
            items = append(items, item)
         }
      }
   }
   return items
}
