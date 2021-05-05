package main

import (
   "flag"
   "fmt"
   "github.com/89z/youtube"
   "net/http"
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

func download(video youtube.Video, itag int) error {
   format, e := video.NewFormat(itag)
   if e != nil { return e }
   // source
   req, e := format.NewRequest()
   if e != nil { return e }
   fmt.Println("Get", req.URL)
   res, e := new(http.Client).Do(req)
   if e != nil { return e }
   switch res.StatusCode {
   case http.StatusOK, http.StatusPartialContent:
   default:
      return fmt.Errorf("StatusCode %v", res.StatusCode)
   }
   defer res.Body.Close()
   // destination
   semi := strings.IndexByte(format.MimeType, ';')
   name := fmt.Sprintf(
      "%v - %v.%v",
      video.VideoDetails.Author,
      video.VideoDetails.Title,
      // audio/webm; codecs="opus"
      format.MimeType[6:semi],
   )
   file, e := os.Create(strings.Map(clean, name))
   if e != nil { return e }
   defer file.Close()
   // copy
   begin := time.Now()
   _, e = file.ReadFrom(res.Body)
   fmt.Println(time.Since(begin))
   return e
}

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
