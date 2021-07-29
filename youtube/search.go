package youtube

import (
   "encoding/json"
   "github.com/corona10/goimagehash"
   "image/jpeg"
   "net/http"
   "sort"
)

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct{
            ItemSectionRenderer struct {
               Contents []struct {
                  CompactVideoRenderer Video
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

func (s Search) Videos() VideoSlice {
   var v VideoSlice
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         v = append(v, item.CompactVideoRenderer)
      }
   }
   return v
}

type Video struct {
   VideoID string
}

type VideoSlice []Video

func hash(addr string, img *Image) (*goimagehash.ImageHash, error) {
   r, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   i, err := jpeg.Decode(r.Body)
   if err != nil {
      return nil, err
   }
   if img != nil {
      i = img.SubImage(i)
   }
   return goimagehash.DifferenceHash(i)
}

func distance(a *goimagehash.ImageHash, v Video) (int, error) {
   img := Image{480, 360, 270, "hqdefault", JPG}
   addr := img.Address(v.VideoID)
   b, err := hash(addr, &img)
   if err != nil {
      return 0, err
   }
   return a.Distance(b)
}

func (v VideoSlice) Sort(musicbrainz string) error {
   mb, err := hash(musicbrainz, nil)
   if err != nil {
      return err
   }
   sort.Slice(v, func(a, b int) bool {
      da, err := distance(mb, v[a])
      if err != nil {
         panic(err)
      }
      db, err := distance(mb, v[b])
      if err != nil {
         panic(err)
      }
      return da < db
   })
   return nil
}
