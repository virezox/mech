package main

import (
   "flag"
   "fmt"
   "github.com/89z/youtube"
   "net/url"
   "os"
   "strings"
   "time"
)

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

func clean(r rune) rune {
   switch {
      case strings.ContainsRune(`"*/:<>?\|`, r): return -1
      default: return r
   }
}

func download(video youtube.Video, itag int, update bool) error {
   format, err := video.NewFormat(itag)
   if err != nil { return err }
   // destination
   semi := strings.IndexByte(format.MimeType, ';')
   name := fmt.Sprint(
      video.VideoDetails.Author, "-",
      video.VideoDetails.Title, "-",
      itag, ".",
      format.MimeType[6:semi], // audio/webm; codecs="opus"
   )
   file, err := os.Create(strings.Map(clean, name))
   if err != nil { return err }
   defer file.Close()
   // source
   begin := time.Now()
   err = format.Write(file, update)
   if err != nil { return err }
   fmt.Println(time.Since(begin))
   return nil
}

func main() {
   var (
      atag, vtag int
      down, update bool
   )
   flag.BoolVar(&down, "d", false, "download")
   flag.BoolVar(&update, "u", false, "update base.js")
   flag.IntVar(&atag, "a", 251, "audio (0 to skip)")
   flag.IntVar(&vtag, "v", 247, "video (0 to skip)")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("youtube [flags] URL")
      flag.PrintDefaults()
      os.Exit(1)
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
   if down {
      if atag > 0 {
         err := download(video, atag, update)
         if err != nil {
            panic(err)
         }
      }
      if vtag > 0 {
         err := download(video, vtag, update)
         if err != nil {
            panic(err)
         }
      }
   } else {
      info(video)
   }
}
