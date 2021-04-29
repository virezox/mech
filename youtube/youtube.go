package main

import (
   "encoding/xml"
   "net/http"
   "os"
)

type mpd struct {
   Period struct {
      AdaptationSet []struct {
         Representation []struct {
            BaseURL string
            SegmentList struct {
               Initialization struct {
                  SourceURL string `xml:"sourceURL,attr"`
               }
               SegmentURL []struct {
                  Media string `xml:"media,attr"`
               }
            }
         }
      }
   }
}

func main() {
   f, e := os.Open("manifest.xml")
   if e != nil {
      panic(e)
   }
   defer f.Close()
   var m mpd
   xml.NewDecoder(f).Decode(&m)
   f, e = os.Create("file.webm")
   if e != nil {
      panic(e)
   }
   defer f.Close()
   rep := m.Period.AdaptationSet[3].Representation[0]
   // get init
   println(rep.SegmentList.Initialization.SourceURL)
   r, e := http.Get(rep.BaseURL + rep.SegmentList.Initialization.SourceURL)
   if e != nil {
      panic(e)
   }
   defer r.Body.Close()
   f.ReadFrom(r.Body)
   // get segments
   for _, segment := range rep.SegmentList.SegmentURL {
      println(segment.Media)
      r, e := http.Get(rep.BaseURL + segment.Media)
      if e != nil {
         panic(e)
      }
      defer r.Body.Close()
      f.ReadFrom(r.Body)
   }
}
