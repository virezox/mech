package youtube

import (
   "fmt"
   "sort"
)

const (
   WidthAutoHeightBlack = 0
   WidthAuto = 1
   WidthBlack = 2
   HeightCrop = 3
)

const (
   JPG = 1
   WebP = 0
)

var AdaptiveImages = Images{
   {90, WidthAutoHeightBlack, JPG, "default"},
   {90, WidthAutoHeightBlack, WebP, "default"},
   {90, WidthBlack, JPG, "1"},
   {90, WidthBlack, JPG, "2"},
   {90, WidthBlack, JPG, "3"},
   {90, WidthBlack, WebP, "1"},
   {90, WidthBlack, WebP, "2"},
   {90, WidthBlack, WebP, "3"},
   {180, HeightCrop, JPG, "mq1"},
   {180, HeightCrop, JPG, "mq2"},
   {180, HeightCrop, JPG, "mq3"},
   {180, HeightCrop, WebP, "mq1"},
   {180, HeightCrop, WebP, "mq2"},
   {180, HeightCrop, WebP, "mq3"},
   {180, WidthAuto, JPG, "mqdefault"},
   {180, WidthAuto, WebP, "mqdefault"},
   {360, WidthAutoHeightBlack, JPG, "0"},
   {360, WidthAutoHeightBlack, JPG, "hqdefault"},
   {360, WidthAutoHeightBlack, WebP, "0"},
   {360, WidthAutoHeightBlack, WebP, "hqdefault"},
   {360, WidthBlack, JPG, "hq1"},
   {360, WidthBlack, JPG, "hq2"},
   {360, WidthBlack, JPG, "hq3"},
   {360, WidthBlack, WebP, "hq1"},
   {360, WidthBlack, WebP, "hq2"},
   {360, WidthBlack, WebP, "hq3"},
   {480, WidthAutoHeightBlack, JPG, "sddefault"},
   {480, WidthAutoHeightBlack, WebP, "sddefault"},
   {480, WidthBlack, JPG, "sd1"},
   {480, WidthBlack, JPG, "sd2"},
   {480, WidthBlack, JPG, "sd3"},
   {480, WidthBlack, WebP, "sd1"},
   {480, WidthBlack, WebP, "sd2"},
   {480, WidthBlack, WebP, "sd3"},
   {720, WidthAuto, JPG, "hq720"},
   {720, WidthAuto, JPG, "maxresdefault"},
   {720, WidthAuto, WebP, "hq720"},
   {720, WidthAuto, WebP, "maxresdefault"},
   {720, WidthBlack, JPG, "maxres1"},
   {720, WidthBlack, JPG, "maxres2"},
   {720, WidthBlack, JPG, "maxres3"},
   {720, WidthBlack, WebP, "maxres1"},
   {720, WidthBlack, WebP, "maxres2"},
   {720, WidthBlack, WebP, "maxres3"},
}

type Image struct {
   Height int
   Frame int
   Format int
   Base string
}

func (i Image) Address(id string) string {
   dir := map[int]string{WebP: "vi_webp", JPG: "vi"}[i.Format]
   ext := map[int]string{WebP: "webp", JPG: "jpg"}[i.Format]
   return fmt.Sprintf("http://i.ytimg.com/%v/%v/%v.%v", dir, id, i.Base, ext)
}

type Images []Image

func (s Images) Filter(keep func(Image)bool) Images {
   var imgs Images
   for _, img := range s {
      if keep(img) {
         imgs = append(imgs, img)
      }
   }
   return imgs
}

func (s Images) Sort() {
   imageFns := []imageFn{
      func(a, b Image) bool {
         return b.Height < a.Height
      },
      func(a, b Image) bool {
         return a.Frame < b.Frame
      },
      func(a, b Image) bool {
         return a.Format < b.Format
      },
   }
   sort.SliceStable(s, func(a, b int) bool {
      sa, sb := s[a], s[b]
      for _, fn := range imageFns {
         if fn(sa, sb) {
            return true
         }
         if fn(sb, sa) {
            break
         }
      }
      return false
   })
}

type imageFn func(a, b Image) bool
