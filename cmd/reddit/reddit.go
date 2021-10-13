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

func download(link *reddit.Link, base, typ string) error {
   addr := link.URL + "/" + base
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   ext, err := mech.Ext(typ)
   if err != nil {
      return err
   }
   name := link.Subreddit + "-" + link.Title + ext
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

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("reddit [-i] [post ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   err := reddit.Valid(id)
   if err != nil {
      panic(err)
   }
   mech.Verbose(true)
   post, err := reddit.NewPost(id)
   if err != nil {
      panic(err)
   }
   link, err := post.Link()
   if err != nil {
      panic(err)
   }
   mpd, err := link.MPD()
   if err != nil {
      panic(err)
   }
   for _, ada := range mpd.Period.AdaptationSet {
      ada.Representation.Sort()
      for _, rep := range ada.Representation {
         if ! info {
            err := download(link, rep.BaseURL, ada.MimeType + rep.MimeType)
            if err != nil {
               panic(err)
            }
            break
         }
         fmt.Printf("%+v\n", rep)
      }
   }
}
