package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/vimeo"
   "sort"
)

func main() {
   clip := new(vimeo.Clip)
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // c
   flag.Int64Var(&clip.ID, "c", 0, "clip ID")
   // d
   var downloadID int64
   flag.Int64Var(&downloadID, "d", 0, "download ID")
   // h
   flag.Int64Var(&clip.UnlistedHash, "h", 0, "unlisted hash")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      vimeo.LogLevel = 1
   }
   if clip.ID >= 1 || address != "" {
      web, err := vimeo.NewJsonWeb()
      if err != nil {
         panic(err)
      }
      if address != "" {
         var err error
         clip, err = vimeo.NewClip(address)
         if err != nil {
            panic(err)
         }
      }
      video, err := web.Video(clip)
      if err != nil {
         panic(err)
      }
      sort.Slice(video.Download, func(a, b int) bool {
         return video.Download[a].Height < video.Download[b].Height
      })
      if info {
         form := video.Format(false)
         fmt.Println(form)
      } else {
         for _, down := range video.Download {
            if down.Video_File_ID == downloadID {
               err := download(down)
               if err != nil {
                  panic(err)
               }
            }
         }
      }
   } else {
      flag.Usage()
   }
}
