package youtube

import (
   "encoding/json"
   "github.com/corona10/goimagehash"
   "image"
   "image/jpeg"
   "net/http"
   "sort"
)

type Item struct {
   CompactVideoRenderer struct {
      VideoID string
   }
}

func (i Item) Distance(other *goimagehash.ImageHash) (int, error) {
   p := Picture{480, 360, 270, "hqdefault", JPG}
   addr := p.Address(i.CompactVideoRenderer.VideoID)
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

type ItemSlice []Item

func (i ItemSlice) Sort(img image.Image) error {
   other, err := goimagehash.DifferenceHash(img)
   if err != nil {
      return err
   }
   sort.Slice(i, func(a, b int) bool {
      da, err := i[a].Distance(other)
      if err != nil {
         panic(err)
      }
      db, err := i[b].Distance(other)
      if err != nil {
         panic(err)
      }
      return da < db
   })
   return nil
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents ItemSlice
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

func (s Search) Items() ItemSlice {
   var items ItemSlice
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         if item.CompactVideoRenderer.VideoID != "" {
            items = append(items, item)
         }
      }
   }
   return items
}
