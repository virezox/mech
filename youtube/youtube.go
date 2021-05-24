package main

import (
   "fmt"
   "github.com/89z/youtube"
   "os"
   "strings"
   "time"
)

func numberFormat(d int64, a ...string) string {
   var (
      e = float64(d)
      f int
   )
   for e >= 1000 {
      e /= 1000
      f++
   }
   return fmt.Sprintf("%.1f", e) + a[f]
}

func getInfo(video youtube.Video) {
   fmt.Println("Author:", video.Author())
   fmt.Println("Title:", video.Title())
   fmt.Println()
   for _, f := range video.Formats() {
      fmt.Printf(
         "itag %v, height %v, %v, %v, %v\n",
         f.Itag,
         f.Height,
         numberFormat(f.Bitrate, "", " kb/s", " mb/s", " gb/s"),
         numberFormat(f.ContentLength, "", " kB", " MB", " GB"),
         f.MimeType,
      )
   }
}

func clean(r rune) rune {
   switch {
      case strings.ContainsRune(`"*/:<>?\|`, r): return -1
      default: return r
   }
}

func download(video youtube.Video, format youtube.Format) error {
   ext := map[string]string{
      "audio/mp4;": ".m4a",
      "audio/webm": ".weba",
      "video/mp4;": ".mp4",
      "video/webm": ".webm",
   }[format.MimeType[:10]]
   create := video.Author() + "-" + video.Title() + ext
   file, err := os.Create(strings.Map(clean, create))
   if err != nil { return err }
   defer func() {
      file.Close()
      if err != nil {
         os.Remove(file.Name())
      }
   }()
   begin := time.Now()
   if err := format.Write(file)
   err != nil { return err }
   fmt.Println(time.Since(begin))
   return nil
}
