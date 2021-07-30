package youtube

import (
   "encoding/json"
   "fmt"
   "github.com/corona10/goimagehash"
   "image"
   "image/jpeg"
   "net/http"
)

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents []struct {
                  CompactVideoRenderer struct {
                     VideoID string
                  }
               }
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

////////////////////////////////////////////////////////////////////////////////

type Result struct {
   VideoID string
   Distance int
}

func (s Search) Results() []Result {
   var ress []Result
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         if item.CompactVideoRenderer.VideoID != "" {
            res := Result{VideoID: item.CompactVideoRenderer.VideoID}
            ress = append(ress, res)
         }
      }
   }
   return ress
}

func (r *Result) SetDistance(other *goimagehash.ImageHash) error {
   p := Picture{480, 360, 270, "hqdefault", JPG}
   addr := p.Address(r.VideoID)
   if Verbose {
      fmt.Println("GET", addr)
   }
   // BAD
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   img, err := jpeg.Decode(res.Body)
   if err != nil {
      return err
   }
   x0 := (p.Width - p.SubHeight) / 2
   y0 := (p.Height - p.SubHeight) / 2
   rect := image.Rect(x0, y0, x0 + p.SubHeight, y0 + p.SubHeight)
   img = img.(*image.YCbCr).SubImage(rect)
   h, err := goimagehash.DifferenceHash(img)
   if err != nil {
      return err
   }
   d, err := h.Distance(other)
   if err != nil {
      return err
   }
   r.Distance = d
   return nil
}
