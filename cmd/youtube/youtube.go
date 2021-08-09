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
      construct, exchange, refresh bool
      embed, info bool
   )
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   flag.BoolVar(&embed, "e", false, "use embedded player")
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.BoolVar(&info, "i", false, "info only")
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   flag.IntVar(&atag, "a", 0, "audio itag (-1 to skip)")
   flag.IntVar(&vtag, "v", 0, "video itag (-1 to skip)")
   flag.Parse()
   if len(os.Args) == 1 {
      // URL is not required if we are just printing help
      fmt.Println("youtube [flags] [video ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   youtube.Verbose = true
   // info
   if info {
      err := getInfo(id)
      if err != nil {
         panic(err)
      }
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
   p, err := youtube.NewPlayer(id, auth, client)
   if err != nil {
      panic(err)
   }
   // sort
   if p.StreamingData.DashManifestURL != "" {
      panic(p.StreamingData.DashManifestURL)
   }
   p.StreamingData.AdaptiveFormats.Sort()
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
      fmts := p.StreamingData.AdaptiveFormats.Filter(fn)
      if fmts == nil {
         ps := p.PlayabilityStatus
         fmt.Println(ps.Status, ps.Reason)
         return
      }
      err := download(p, fmts[0])
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

func getInfo(id string) error {
   p, err := youtube.NewPlayer(id, youtube.Key, youtube.Mweb)
   if err != nil {
      return err
   }
   fmt.Println("author:", p.Author())
   fmt.Println("title:", p.Title())
   fmt.Println("countries:", p.Countries())
   fmt.Println()
   for _, f := range p.StreamingData.AdaptiveFormats {
      fmt.Printf(
         "itag %v, height %v, %v, %v, %v\n",
         f.Itag,
         f.Height,
         numberFormat(f.Bitrate, "", " kb/s", " mb/s", " gb/s"),
         numberFormat(f.ContentLength, "", " kB", " MB", " GB"),
         f.MimeType,
      )
   }
   return nil
}

func download(p *youtube.Player, f youtube.Format) error {
   name := p.Author() + "-" + p.Title() + f.Ext()
   file, err := os.Create(strings.Map(clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   return f.Write(file)
}
