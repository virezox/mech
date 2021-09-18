package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/reddit"
   "os"
   "strings"
   //"net/http"
   //"path"
)

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
      fmt.Println(t3.Data.URL)
      for _, set := range mpd.Period.AdaptationSet {
         for _, rep := range set.Representation {
            fmt.Printf("%+v\n", rep)
         }
      }
      return
   }
   for _, set := range mpd.Period.AdaptationSet {
      for _, rep := range set.Representation {
         if rep.Height == 0 || rep.Height == height {
            /*
            err := download(cfg, f.URL)
            if err != nil {
               panic(err)
            }
            */
         }
      }
   }
}

/*
func download(cfg *reddit.Config, addr string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := cfg.Video.Owner.Name + "-" + cfg.Video.Title + path.Ext(addr)
   file, err := os.Create(strings.Map(clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
*/

func clean(r rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}
