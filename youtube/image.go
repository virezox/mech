package youtube

import (
   "fmt"
   "image"
   "sort"
)

var (
   WebP = ImageFormat{0, "vi_webp", "webp"}
   JPG = ImageFormat{1, "vi", "jpg"}
)

var Images = ImageSlice{
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

type Image struct {
   Width int
   Height int
   SubHeight int
   Base string
   Format ImageFormat
}

func (i Image) Address(id string) string {
   return fmt.Sprintf(
      "http://i.ytimg.com/%v/%v/%v.%v", i.Format.Dir, id, i.Base, i.Format.Ext,
   )
}

func (i Image) Rect() image.Rectangle {
   x0 := (i.Width - i.SubHeight) / 2
   y0 := (i.Height - i.SubHeight) / 2
   return image.Rect(x0, y0, x0 + i.SubHeight, y0 + i.SubHeight)
}

type ImageFormat struct {
   Sort int
   Dir string
   Ext string
}

type ImageSlice []Image

func (i ImageSlice) Filter(keep func(Image)bool) ImageSlice {
   var imgs ImageSlice
   for _, img := range i {
      if keep(img) {
         imgs = append(imgs, img)
      }
   }
   return imgs
}

func (i ImageSlice) Sort(less ...func(a, b Image) bool) {
   if less == nil {
      less = []func(a, b Image) bool{
         func(a, b Image) bool {
            return b.Height < a.Height
         },
         func(a, b Image) bool {
            return a.SubHeight < b.SubHeight
         },
         func(a, b Image) bool {
            return a.Base < b.Base
         },
         func(a, b Image) bool {
            return a.Format.Sort < b.Format.Sort
         },
      }
   }
   sort.SliceStable(i, func(a, b int) bool {
      ia, ib := i[a], i[b]
      for _, fn := range less {
         if fn(ia, ib) {
            return true
         }
         if fn(ib, ia) {
            break
         }
      }
      return false
   })
}
