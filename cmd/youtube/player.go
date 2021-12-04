package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
   "strings"
)

type choice map[string]bool

func (c choice) download(play *youtube.Player, id string) error {
   if len(c) == 0 {
      c["247"] = true
      c["251"] = true
   }
   // Videos can support both AdaptiveFormats and DASH:
   // zgJT91LA9gA
   for _, form := range play.StreamingData.AdaptiveFormats {
      if c[fmt.Sprint(form.Itag)] {
         ext, err := mech.ExtensionByType(form.MimeType)
         if err != nil {
            return err
         }
         name := play.Author() + "-" + play.Title() + ext
         file, err := os.Create(strings.Map(mech.Clean, name))
         if err != nil {
            return err
         }
         defer file.Close()
         if err := form.Write(file); err != nil {
            return err
         }
      }
   }
   return nil
}

func infoPath(play *youtube.Player, id string) error {
   fmt.Println("author:", play.Author())
   fmt.Println("title:", play.Title())
   fmt.Println()
   for _, f := range play.StreamingData.AdaptiveFormats {
      fmt.Printf(
         "itag %v, height %v, %v, %v, %v\n",
         f.Itag, f.Height, f.Bitrate, f.ContentLength, f.MimeType,
      )
   }
   return nil
}
