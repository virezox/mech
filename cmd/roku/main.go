package main

import (
   "flag"
   "github.com/89z/mech/roku"
)

func main() {
   // b
   var id string
   flag.StringVar(&id, "b", "", "ID")
   // dash
   var isDASH bool
   flag.BoolVar(&isDASH, "dash", false, "DASH download")
   // f
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   var video int64
   flag.Int64Var(&video, "f", 1920832, "video bandwidth")
   // g
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   var audio int64
   flag.Int64Var(&audio, "g", 128000, "audio bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      roku.LogLevel = 1
   }
   if id != "" {
      content, err := roku.NewContent(id)
      if err != nil {
         panic(err)
      }
      if isDASH {
         down := downloader{Content: content, info: info}
         err := down.DASH(audio, video)
         if err != nil {
            panic(err)
         }
      } else {
         err := doHLS(content, video, info)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
