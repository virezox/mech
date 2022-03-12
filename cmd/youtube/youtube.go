package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

func main() {
   var choose choice
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var videoID string
   flag.StringVar(&videoID, "b", "", "video ID")
   // c
   var construct bool
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   // e
   var embed bool
   flag.BoolVar(&embed, "e", false, "use embedded player")
   // f
   choose.itags = make(map[string]bool)
   flag.Func("f", "formats", func(itag string) error {
      choose.itags[itag] = true
      return nil
   })
   // i
   flag.BoolVar(&choose.info, "i", false, "information")
   // r
   var refresh bool
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   // x
   var exchange bool
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.Parse()
   if verbose {
      youtube.LogLevel = 1
   }
   if exchange {
      oauth, err := youtube.NewOAuth()
      if err != nil {
         panic(err)
      }
      fmt.Println(oauth)
      fmt.Scanln()
      exc, err := oauth.Exchange()
      if err != nil {
         panic(err)
      }
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      if err := exc.Create(cache + "/mech/youtube.json"); err != nil {
         panic(err)
      }
   } else if refresh {
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      exc, err := youtube.OpenExchange(cache + "/mech/youtube.json")
      if err != nil {
         panic(err)
      }
      if err := exc.Refresh(); err != nil {
         panic(err)
      }
      if err := exc.Create(cache + "/mech/youtube.json"); err != nil {
         panic(err)
      }
   } else if videoID != "" || address != "" {
      if videoID == "" {
         var err error
         videoID, err = youtube.VideoID(address)
         if err != nil {
            panic(err)
         }
      }
      play, err := player(construct, embed, videoID)
      if err != nil {
         panic(err)
      }
      if err := choose.adaptiveFormats(play); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type choice struct {
   itags map[string]bool
   info bool
}

// Videos can support both AdaptiveFormats and DASH: zgJT91LA9gA
func (c choice) adaptiveFormats(play *youtube.Player) error {
   fmt.Println(play.Status())
   if c.info {
      fmt.Println(play.Details())
   }
   if len(c.itags) == 0 {
      c.itags = map[string]bool{
         "247": true, // youtube.com/watch?v=Leq8J0E2TQ0
         "251": true,
         "302": true, // youtube.com/watch?v=kVNl1P9StSU
      }
   }
   for _, ada := range play.StreamingData.AdaptiveFormats {
      switch {
      case c.info:
         ada.URL = ""
         form, err := ada.Format()
         if err != nil {
            return err
         }
         fmt.Println(form)
      case c.itags[fmt.Sprint(ada.Itag)]:
         name, err := ada.Name(play)
         if err != nil {
            return err
         }
         file, err := os.Create(name)
         if err != nil {
            return err
         }
         defer file.Close()
         if err := ada.Write(file); err != nil {
            return err
         }
      }
   }
   return nil
}

func player(construct, embed bool, id string) (*youtube.Player, error) {
   client := youtube.Android
   if embed {
      client = youtube.Embed
   }
   if construct {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      exc, err := youtube.OpenExchange(cache, "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
      return client.PlayerHeader(exc.Header(), id)
   }
   return client.Player(id)
}

