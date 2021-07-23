package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
   "strings"
)

func main() {
   var (
      atag, vtag int
      embed, info bool
      construct, exchange, refresh bool
   )
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   flag.BoolVar(&embed, "e", false, "use embedded player")
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.BoolVar(&info, "i", false, "info only")
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   flag.IntVar(&atag, "a", 0, "audio (-1 to skip)")
   flag.IntVar(&vtag, "v", 0, "video (-1 to skip)")
   flag.Parse()
   if len(os.Args) == 1 {
      // URL is not required if we are just printing help
      fmt.Println("youtube [flags] [video ID]")
      flag.PrintDefaults()
      return
   }
   // exchange
   if exchange {
      err := authExchange()
      if err != nil {
         panic(err)
      }
      return
   }
   // refresh
   if refresh {
      err := authRefresh()
      if err != nil {
         panic(err)
      }
      return
   }
   // head
   auth := youtube.Key
   if construct {
      x, err := authConstruct()
      if err != nil {
         panic(err)
      }
      auth = youtube.Auth{"Authorization", "Bearer " + x.Access_Token}
   }
   // body
   client := youtube.Android
   if embed {
      client = youtube.Embed
   }
   play, err := youtube.NewPlayer(flag.Arg(0), auth, client)
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
         ps := play.PlayabilityStatus
         fmt.Println(ps.Status, ps.ReasonTitle)
         return
      }
      err := download(play, fmts[0])
      if err != nil {
         panic(err)
      }
   }
}

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
   return f.Write(file)
}
