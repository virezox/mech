package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/vimeo"
   "sort"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // c
   var clipID int64
   flag.Int64Var(&clipID, "c", 0, "clip ID")
   // d
   var downloadID int64
   flag.Int64Var(&downloadID, "d", 0, "download ID")
   // h
   var unlistedHash int64
   flag.Int64Var(&unlistedHash, "h", 0, "unlisted hash")
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
   clip, err := newClip(clipID, unlistedHash, address)
   if err != nil {
      flag.Usage()
   } else {
      web, err := vimeo.NewJsonWeb()
      if err != nil {
         panic(err)
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
   }
}
