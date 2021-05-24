package main

import (
   "flag"
   "fmt"
   "github.com/89z/youtube"
   "net/url"
   "os"
)

func main() {
   var (
      atag, vtag int
      info, update bool
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.BoolVar(&update, "u", false, "update base.js")
   flag.IntVar(&atag, "a", 251, "audio (0 to skip)")
   flag.IntVar(&vtag, "v", 247, "video (0 to skip)")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("youtube [flags] [URL]")
      flag.PrintDefaults()
      return
   }
   // update
   if update {
      base, err := youtube.NewBaseJS()
      if err != nil {
         panic(err)
      }
      base.Get()
      return
   }
   // check URL
   if flag.NArg() != 1 {
      panic("missing URL")
   }
   arg := flag.Arg(0)
   watch, err := url.Parse(arg)
   if err != nil {
      panic(err)
   }
   id := watch.Query().Get("v")
   video, err := youtube.NewVideo(id)
   if err != nil {
      panic(err)
   }
   // info
   if info {
      getInfo(video)
      return
   }
   // check formats
   var afmt, vfmt youtube.Format
   if atag > 0 {
      if afmt, err = video.NewFormat(atag); err != nil {
         panic(err)
      }
   }
   if vtag > 0 {
      if vfmt, err = video.NewFormat(vtag); err != nil {
         panic(err)
      }
   }
   // download
   if atag > 0 {
      if err := download(video, afmt); err != nil {
         panic(err)
      }
   }
   if vtag > 0 {
      if err := download(video, vfmt); err != nil {
         panic(err)
      }
   }
}
