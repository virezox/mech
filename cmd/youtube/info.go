package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
)

func infoPath(id string) error {
   p, err := youtube.NewPlayer(id, youtube.Key, youtube.Mweb)
   if err != nil {
      return err
   }
   fmt.Println("author:", p.Author())
   fmt.Println("title:", p.Title())
   fmt.Println("countries:", p.Countries())
   fmt.Println()
   for _, f := range p.StreamingData.AdaptiveFormats {
      fmt.Printf(
         "itag %v, height %v, %v, %v, %v\n",
         f.Itag, f.Height, f.Bitrate, f.ContentLength, f.MimeType,
      )
   }
   return nil
}
