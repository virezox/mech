package youtube

import (
   "encoding/json"
   "fmt"
   "github.com/corona10/goimagehash"
   "image"
   "image/jpeg"
   "net/http"
)

var Distance = make(map[string]int)

type CompactVideoRenderer struct {
   VideoID string
}

type Item struct {
   CompactVideoRenderer `json:"compactVideoRenderer"`
}

func (i Item) Distance(other *goimagehash.ImageHash) (int, error) {
   if d, ok := Distance[i.VideoID]; ok {
      return d, nil
   }
   p := Picture{480, 360, 270, "hqdefault", JPG}
   addr := p.Address(i.VideoID)
   if Verbose {
      fmt.Println("GET", addr)
   }
   res, err := http.Get(addr)
   if err != nil {
      return 0, err
   }
   defer res.Body.Close()
   img, err := jpeg.Decode(res.Body)
   if err != nil {
      return 0, err
   }
   x0 := (p.Width - p.SubHeight) / 2
   y0 := (p.Height - p.SubHeight) / 2
   rect := image.Rect(x0, y0, x0 + p.SubHeight, y0 + p.SubHeight)
   img = img.(*image.YCbCr).SubImage(rect)
   h, err := goimagehash.DifferenceHash(img)
   if err != nil {
      return 0, err
   }
   d, err := h.Distance(other)
   if err != nil {
      return 0, err
   }
   Distance[i.VideoID] = d
   return d, nil
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
   body.Context.Client = Mweb
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
         if item.VideoID != "" {
            items = append(items, item)
         }
      }
   }
   return items
}
