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

func clean(r rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}

func main() {
   var (
      atag, vtag int
      detailPage, info bool
   )
   flag.BoolVar(&detailPage, "d", false, "playback on other websites has been disabled")
   flag.BoolVar(&info, "i", false, "info only")
   flag.IntVar(&atag, "a", 0, "audio (-1 to skip)")
   flag.IntVar(&vtag, "v", 0, "video (-1 to skip)")
   flag.Parse()
   if flag.NArg() != 1 {
      // URL is not required if we are just printing help
      fmt.Println("youtube [flags] [URL]")
      flag.PrintDefaults()
      return
   }
   // check URL
   watch, err := url.Parse(flag.Arg(0))
   if err != nil {
      panic(err)
   }
   id := watch.Query().Get("v")
   play, err := youtube.GetVideoInfo(id, detailPage)
   if err != nil {
      panic(err)
   }
   // info
   if info {
      getInfo(play)
      return
   }
   // sort
   play.AdaptiveFormats.Sort()
   formats := []youtube.Format{
      {Itag: atag}, {Itag: vtag, Height: 720},
   }
   for _, a := range formats {
      var fn func(youtube.Format)bool
      switch a.Itag {
      case -1:
         continue
      case 0:
         fn = func(b youtube.Format) bool {
            return b.Height <= a.Height
         }
      default:
         fn = func(b youtube.Format) bool {
            return b.Itag == a.Itag
         }
      }
      fmts := play.AdaptiveFormats.Filter(fn)
      if fmts == nil {
         fmt.Println(play.PlayabilityStatus.Reason)
         return
      }
      err := download(play, fmts[0])
      if err != nil {
         panic(err)
      }
   }
}

func getInfo(play *youtube.Player) {
   fmt.Println("author:", play.Author)
   fmt.Println("title:", play.Title)
   fmt.Println()
   for _, f := range play.AdaptiveFormats {
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

func download(p *youtube.Player, f youtube.Format) error {
   create := strings.Map(clean, p.Author + "-" + p.Title + f.Ext())
   file, err := os.Create(create)
   if err != nil {
      return err
   }
   defer file.Close()
   begin := time.Now()
   if err := f.Write(file); err != nil {
      return err
   }
   fmt.Println(time.Since(begin))
   return nil
}
