package youtube

import (
   "encoding/json"
   "github.com/corona10/goimagehash"
   "image"
   "image/jpeg"
   "net/http"
   "sort"
)

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents VideoSlice
            }
         }
      }
   }
}

func NewSearch(query string) (*Search, error) {
   var body youTubeI
   body.Context.Client = Mweb
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

func (s Search) Videos() VideoSlice {
   var v VideoSlice
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      v = append(v, sect.ItemSectionRenderer.Contents...)
   }
   return v
}

type Video struct {
   CompactVideoRenderer struct {
      VideoID string
   }
}

func (v Video) Distance(other *goimagehash.ImageHash) (int, error) {
   p := Picture{480, 360, 270, "hqdefault", JPG}
   addr := p.Address(v.CompactVideoRenderer.VideoID)
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
   return h.Distance(other)
}

type VideoSlice []Video

func (v VideoSlice) Sort(img image.Image) error {
   _, err := goimagehash.DifferenceHash(img)
   if err != nil {
      return err
   }
   sort.Slice(v, func(a, b int) bool {
      return true
   })
   return nil
}
