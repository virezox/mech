package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/youtube"
   "net/url"
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

func getInfo(video youtube.Player) {
   fmt.Println("author:", video.Author())
   fmt.Println("title:", video.Title())
   fmt.Println("countries:", video.Countries())
   if len(video.Formats()) > 0 {
      fmt.Println()
   }
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
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}

func download(video youtube.Player, format youtube.Format) error {
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


func main() {
   var (
      atag, vtag int
      info bool
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.IntVar(&atag, "a", 251, "audio (0 to skip)")
   flag.IntVar(&vtag, "v", 247, "video (0 to skip)")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("youtube [flags] [URL]")
      flag.PrintDefaults()
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
   video, err := youtube.NewPlayer(id)
   if err != nil {
      panic(err)
   }
   // info
   if info {
      getInfo(video)
      return
   }
   // check formats
   var forms []youtube.Format
   if atag > 0 {
      format, err := video.NewFormat(atag)
      if err != nil {
         panic(err)
      }
      forms = append(forms, format)
   }
   if vtag > 0 {
      format, err := video.NewFormat(vtag)
      if err != nil {
         panic(err)
      }
      forms = append(forms, format)
   }
   // download
   for _, form := range forms {
      err := download(video, form)
      if err != nil {
         panic(err)
      }
   }
}
