package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/reddit"
   "net/http"
   "os"
   "strings"
)


func main() {
   var dashFormat, hlsFormat bool
   flag.BoolVar(&dashFormat, "df", false, "DASH formats")
   flag.BoolVar(&hlsFormat, "hf", false, "HLS formats")
   dashIDs := make(map[string]bool)
   flag.Func("d", "DASH IDs", func(id string) error {
      dashIDs[id] = true
      return nil
   })
   hlsIDs := make(map[string]bool)
   flag.Func("h", "HLS IDs", func(id string) error {
      hlsIDs[id] = true
      return nil
   })
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("reddit [flags] [post ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   if ! reddit.Valid(id) {
      panic("invalid ID")
   }
   mech.Verbose = true
   post, err := reddit.NewPost(id)
   if err != nil {
      panic(err)
   }
   link, err := post.Link()
   if err != nil {
      panic(err)
   }
   if dashFormat || len(dashIDs) >= 1 {
      dash, err := link.DASH()
      if err != nil {
         panic(err)
      }
      for _, ada := range dash.Period.AdaptationSet {
         for _, rep := range ada.Representation {
            if dashFormat {
               fmt.Printf("%+v\n", rep)
            } else if dashIDs[rep.ID] {
               err := download(link, rep.BaseURL, rep.MimeType)
               if err != nil {
                  panic(err)
               }
            }
         }
      }
   }
   if hlsFormat || len(hlsIDs) >= 1 {
      hlss, err := link.HLS()
      if err != nil {
         panic(err)
      }
      for _, hls := range hlss {
         if hlsFormat {
            fmt.Printf("%+v\n", hls)
         } else {
         }
      }
   }
}

func download(link *reddit.Link, addr, typ string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   exts, err := mech.ExtensionsByType(typ)
   if exts == nil {
      return fmt.Errorf("exts %v, err %v", exts, err)
   }
   name := link.Subreddit + "-" + link.Title + exts[0]
   file, err := os.Create(strings.Map(mech.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
