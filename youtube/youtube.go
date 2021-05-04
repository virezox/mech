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
      down bool
      itag int
   )
   flag.BoolVar(&down, "d", false, "download")
   flag.IntVar(&itag, "i", 251, "itag")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("youtube [flags] URL")
      flag.PrintDefaults()
      os.Exit(1)
   }
   arg := flag.Arg(0)
   watch, e := url.Parse(arg)
   if e != nil {
      panic(e)
   }
   id := watch.Query().Get("v")
   video, e := youtube.NewVideo(id)
   if e != nil {
      panic(e)
   }
   if down {
      e := download(video, itag)
      if e != nil {
         panic(e)
      }
   } else {
      info(video)
   }
}

func download(video youtube.Video, itag int) error {
   stream, e := video.GetStream(itag)
   if e != nil { return e }
   fmt.Println("Get", stream)
   res, e := http.Get(stream)
   if e != nil { return e }
   defer res.Body.Close()
   /*
   audio/webm; codecs="opus"
   */
}

func info(video youtube.Video) {
   for _, f := range video.StreamingData.AdaptiveFormats {
      fmt.Println(
         "itag:", f.Itag,
         "bitrate:", f.Bitrate,
         "height:", f.Height,
         "mimetype:", f.MimeType,
      )
   }
}
