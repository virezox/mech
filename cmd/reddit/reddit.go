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

func download(t3 *reddit.T3, base, typ string) error {
   addr := t3.URL + "/" + base
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
   name := t3.Subreddit + "-" + t3.Title + ext
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
   var (
      height int
      info bool
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.IntVar(&height, "h", 720, "height")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("reddit [flags] [post ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   err := reddit.ValidID(id)
   if err != nil {
      panic(err)
   }
   reddit.Verbose = true
   post, err := reddit.NewPost(id)
   if err != nil {
      panic(err)
   }
   t3, err := post.T3()
   if err != nil {
      panic(err)
   }
   mpd, err := t3.MPD()
   if err != nil {
      panic(err)
   }
   // info
   if info {
      fmt.Println(t3.URL)
      for _, ada := range mpd.Period.AdaptationSet {
         for _, rep := range ada.Representation {
            fmt.Printf("%+v\n", rep)
         }
      }
      return
   }
   for _, ada := range mpd.Period.AdaptationSet {
      for _, rep := range ada.Representation {
         if rep.Height == 0 || rep.Height == height {
            err := download(t3, rep.BaseURL, ada.MimeType)
            if err != nil {
               panic(err)
            }
         }
      }
   }
}
