package youtube

import (
   "fmt"
   "sort"
)

var (
   WebP = PictureFormat{0, "vi_webp", "webp"}
   JPG = PictureFormat{1, "vi", "jpg"}
)

var Pictures = PictureSlice{
   {120, 90, 68, "default", JPG},
   {120, 90, 90, "1", JPG},
   {120, 90, 90, "2", JPG},
   {120, 90, 90, "3", JPG},
   {120, 90, 68, "default", WebP},
   {120, 90, 90, "1", WebP},
   {120, 90, 90, "2", WebP},
   {120, 90, 90, "3", WebP},
   {320, 180, 180, "mqdefault", JPG},
   {320, 180, 320, "mq1", JPG},
   {320, 180, 320, "mq2", JPG},
   {320, 180, 320, "mq3", JPG},
   {320, 180, 180, "mqdefault", WebP},
   {320, 180, 320, "mq1", WebP},
   {320, 180, 320, "mq2", WebP},
   {320, 180, 320, "mq3", WebP},
   {480, 360, 270, "0", JPG},
   {480, 360, 270, "hqdefault", JPG},
   {480, 360, 360, "hq1", JPG},
   {480, 360, 360, "hq2", JPG},
   {480, 360, 360, "hq3", JPG},
   {480, 360, 270, "0", WebP},
   {480, 360, 270, "hqdefault", WebP},
   {480, 360, 360, "hq1", WebP},
   {480, 360, 360, "hq2", WebP},
   {480, 360, 360, "hq3", WebP},
   {640, 480, 360, "sddefault", JPG},
   {640, 480, 480, "sd1", JPG},
   {640, 480, 480, "sd2", JPG},
   {640, 480, 480, "sd3", JPG},
   {640, 480, 360, "sddefault", WebP},
   {640, 480, 480, "sd1", WebP},
   {640, 480, 480, "sd2", WebP},
   {640, 480, 480, "sd3", WebP},
   {1280, 720, 720, "hq720", JPG},
   {1280, 720, 720, "maxres1", JPG},
   {1280, 720, 720, "maxres2", JPG},
   {1280, 720, 720, "maxres3", JPG},
   {1280, 720, 720, "maxresdefault", JPG},
   {1280, 720, 720, "hq720", WebP},
   {1280, 720, 720, "maxres1", WebP},
   {1280, 720, 720, "maxres2", WebP},
   {1280, 720, 720, "maxres3", WebP},
   {1280, 720, 720, "maxresdefault", WebP},
}

type Picture struct {
   Width int
   Height int
   SubHeight int
   Base string
   Format PictureFormat
}

func (p Picture) Address(id string) string {
   return fmt.Sprintf(
      "http://i.ytimg.com/%v/%v/%v.%v", p.Format.Dir, id, p.Base, p.Format.Ext,
   )
}

type PictureFormat struct {
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
