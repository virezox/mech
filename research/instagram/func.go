package instagram

import (
   "encoding/xml"
)

func appendImage() {
   /*
   var max int
   for _, can := range i.Image_Versions2.Candidates {
      if can.Height > max {
         addrs = []string{can.URL}
         max = can.Height
      }
   }
   */
}

func appendVideo() {
   /*
   var max int
   for _, ver := range i.Video_Versions {
      if ver.Type > max {
         addrs = []string{ver.URL}
         max = ver.Type
      }
   }
   */
}

func appendManifest(urls []string, manifest string) ([]string, error) {
   if manifest != "" {
      var video dashManifest
      err := xml.Unmarshal([]byte(manifest), &video)
      if err != nil {
         return nil, err
      }
      for _, ada := range video.Period.AdaptationSet {
         var (
            url string
            max int
         )
         for _, rep := range ada.Representation {
            if rep.Bandwidth > max {
               url = rep.BaseURL
               max = rep.Bandwidth
            }
         }
         urls = append(urls, url)
      }
   }
   return urls, nil
}


