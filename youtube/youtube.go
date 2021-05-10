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

func getInfo(video youtube.Video) {
   for _, format := range video.StreamingData.AdaptiveFormats {
      fmt.Println(
         "itag:", format.Itag,
         "bitrate:", format.Bitrate,
         "height:", format.Height,
         "mimetype:", format.MimeType,
      )
   }
}

func clean(r rune) rune {
   switch {
      case strings.ContainsRune(`"*/:<>?\|`, r): return -1
      default: return r
   }
}

func download(video youtube.Video, itag int) error {
   format, err := video.NewFormat(itag)
   if err != nil { return err }
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
   begin := time.Now()
   err = format.Write(file)
   if err != nil { return err }
   fmt.Println(time.Since(begin))
   return nil
}

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
   // audio
   if atag > 0 {
      err := download(video, atag)
      if err != nil {
         panic(err)
      }
   }
   // video
   if vtag > 0 {
      err := download(video, vtag)
      if err != nil {
         panic(err)
      }
   }
}
