package instagram

import (
   "encoding/xml"
)

type ImageVersion struct {
   Candidates []struct {
      Width int
      Height int
      URL string
   }
}

type VideoVersion struct {
   Type int
   Width int
   Height int
   URL string
}

type dashManifest struct {
   Period struct {
      AdaptationSet []struct { // one video one audio
         Representation []struct {
            Width int `xml:"width,attr"`
            Height int `xml:"height,attr"`
            Bandwidth int `xml:"bandwidth,attr"`
            BaseURL string
         }
      }
   }
}

type Item struct {
   Video_DASH_Manifest string
   Image_Versions2 ImageVersion
   Video_Versions []VideoVersion
   Carousel_Media []struct {
      Video_DASH_Manifest string
      Image_Versions2 ImageVersion
      Video_Versions []VideoVersion
   }
}

func (i Item) URLs() ([]string, error) {
   var (
      dst []string
      err error
   )
   if i.Video_DASH_Manifest != "" {
      dst, err = appendManifest(dst, i.Video_DASH_Manifest)
      if err != nil {
         return nil, err
      }
   } else if i.Video_Versions != nil {
      dst = appendVideo(dst, i.Video_Versions)
   } else {
      dst = appendImage(dst, i.Image_Versions2)
   }
   for _, med := range i.Carousel_Media {
      if med.Video_DASH_Manifest != "" {
         dst, err = appendManifest(dst, med.Video_DASH_Manifest)
         if err != nil {
            return nil, err
         }
      } else if med.Video_Versions != nil {
         dst = appendVideo(dst, med.Video_Versions)
      } else {
         dst = appendImage(dst, med.Image_Versions2)
      }
   }
   return dst, nil
}

func appendImage(dst []string, src ImageVersion) []string {
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
   return append(dst, addr)
}

func appendVideo(dst []string, src []VideoVersion) []string {
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
   return append(dst, addr)
}

func appendManifest(dst []string, src string) ([]string, error) {
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
   return dst, nil
}
