package instagram

import (
   "encoding/xml"
)

func (i Item) URLs() ([]string, error) {
   var (
      dst []string
      err error
   )
   dst, err = appendManifest(dst, i.Video_DASH_Manifest)
   if err != nil {
      return nil, err
   }
   dst = appendImage(dst, i.Image_Versions2)
   dst = appendVideo(dst, i.Video_Versions)
   for _, media := range i.Carousel_Media {
      dst, err = appendManifest(dst, media.Video_DASH_Manifest)
      if err != nil {
         return nil, err
      }
      dst = appendImage(dst, media.Image_Versions2)
      dst = appendVideo(dst, media.Video_Versions)
   }
   return dst, nil
}

func appendImage(dst []string, src ImageVersion) []string {
   if src.Candidates != nil {
      var (
         addr string
         max int
      )
      for _, can := range src.Candidates {
         if can.Height > max {
            addr = can.URL
            max = can.Height
         }
      }
      dst = append(dst, addr)
   }
   return dst
}

func appendVideo(dst []string, src []VideoVersion) []string {
   if src != nil {
      var (
         addr string
         max int
      )
      for _, ver := range src {
         if ver.Type > max {
            addr = ver.URL
            max = ver.Type
         }
      }
      dst = append(dst, addr)
   }
   return dst
}

func appendManifest(dst []string, src string) ([]string, error) {
   if src != "" {
      var video dashManifest
      err := xml.Unmarshal([]byte(src), &video)
      if err != nil {
         return nil, err
      }
      for _, ada := range video.Period.AdaptationSet {
         var (
            addr string
            max int
         )
         for _, rep := range ada.Representation {
            if rep.Bandwidth > max {
               addr = rep.BaseURL
               max = rep.Bandwidth
            }
         }
         dst = append(dst, addr)
      }
   }
   return dst, nil
}
