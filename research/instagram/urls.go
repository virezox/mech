package instagram

import (
   "net/http"
)

func (i item) URLs() ([]string, error) {
   switch i.Media_Type {
   case 1:
      ver := getImage(i.Image_Versions2.Candidates)
      return []string{ver.URL}, nil
   case 2:
      ver, err := getVideo(i.Video_Versions)
      if err != nil {
         return nil, err
      }
      return []string{ver.URL}, nil
   }
   var dst []string
   for _, car := range i.Carousel_Media {
      var (
         err error
         ver *version
      )
      switch car.Media_Type {
      case 1:
         ver = getImage(i.Image_Versions2.Candidates)
      case 2:
         ver, err = getVideo(i.Video_Versions)
         if err != nil {
            return nil, err
         }
      }
      dst = append(dst, ver.URL)
   }
   return dst, nil
}

func getImage(srcs []version) *version {
   var dst version
   for _, src := range srcs {
      if src.Height > dst.Height {
         dst = src
      }
   }
   return &dst
}

func getVideo(srcs []version) (*version, error) {
   var (
      dst version
      length int64
   )
   for _, src := range srcs {
      res, err := http.Head(src.URL)
      if err != nil {
         return nil, err
      }
      if res.ContentLength > length {
         dst = src
      }
   }
   return &dst, nil
}
