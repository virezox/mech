package youtube

import (
   "fmt"
   "sort"
)

var (
   webP = pictureFormat{0, "vi_webp", "webp"}
   jpg = pictureFormat{1, "vi", "jpg"}
)

var Pictures = PictureSlice{
   {120, 90, 68, "default", jpg},
   {120, 90, 90, "1", jpg},
   {120, 90, 90, "2", jpg},
   {120, 90, 90, "3", jpg},
   {120, 90, 68, "default", webP},
   {120, 90, 90, "1", webP},
   {120, 90, 90, "2", webP},
   {120, 90, 90, "3", webP},
   {320, 180, 180, "mqdefault", jpg},
   {320, 180, 320, "mq1", jpg},
   {320, 180, 320, "mq2", jpg},
   {320, 180, 320, "mq3", jpg},
   {320, 180, 180, "mqdefault", webP},
   {320, 180, 320, "mq1", webP},
   {320, 180, 320, "mq2", webP},
   {320, 180, 320, "mq3", webP},
   {480, 360, 270, "0", jpg},
   {480, 360, 270, "hqdefault", jpg},
   {480, 360, 360, "hq1", jpg},
   {480, 360, 360, "hq2", jpg},
   {480, 360, 360, "hq3", jpg},
   {480, 360, 270, "0", webP},
   {480, 360, 270, "hqdefault", webP},
   {480, 360, 360, "hq1", webP},
   {480, 360, 360, "hq2", webP},
   {480, 360, 360, "hq3", webP},
   {640, 480, 360, "sddefault", jpg},
   {640, 480, 480, "sd1", jpg},
   {640, 480, 480, "sd2", jpg},
   {640, 480, 480, "sd3", jpg},
   {640, 480, 360, "sddefault", webP},
   {640, 480, 480, "sd1", webP},
   {640, 480, 480, "sd2", webP},
   {640, 480, 480, "sd3", webP},
   {1280, 720, 720, "hq720", jpg},
   {1280, 720, 720, "maxres1", jpg},
   {1280, 720, 720, "maxres2", jpg},
   {1280, 720, 720, "maxres3", jpg},
   {1280, 720, 720, "maxresdefault", jpg},
   {1280, 720, 720, "hq720", webP},
   {1280, 720, 720, "maxres1", webP},
   {1280, 720, 720, "maxres2", webP},
   {1280, 720, 720, "maxres3", webP},
   {1280, 720, 720, "maxresdefault", webP},
}

type Picture struct {
   Width int
   Height int
   SubHeight int
   Base string
   Format pictureFormat
}

func (p Picture) Address(id string) string {
   return fmt.Sprintf(
      "http://i.ytimg.com/%v/%v/%v.%v", p.Format.Dir, id, p.Base, p.Format.Ext,
   )
}

type pictureFormat struct {
   Size int
   Dir string
   Ext string
}

type PictureSlice []Picture

func (p PictureSlice) Filter(keep func(Picture)bool) PictureSlice {
   var pics PictureSlice
   for _, pic := range p {
      if keep(pic) {
         pics = append(pics, pic)
      }
   }
   return pics
}

func (p PictureSlice) Sort(less ...func(a, b Picture) bool) {
   if less == nil {
      less = []func(a, b Picture) bool{
         func(a, b Picture) bool {
            return b.Height < a.Height
         },
         func(a, b Picture) bool {
            return a.SubHeight < b.SubHeight
         },
         func(a, b Picture) bool {
            return a.Base < b.Base
         },
         func(a, b Picture) bool {
            return a.Format.Size < b.Format.Size
         },
      }
   }
   sort.SliceStable(p, func(a, b int) bool {
      pa, pb := p[a], p[b]
      for _, fn := range less {
         if fn(pa, pb) {
            return true
         }
         if fn(pb, pa) {
            break
         }
      }
      return false
   })
}
