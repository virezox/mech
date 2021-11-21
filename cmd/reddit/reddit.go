package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/reddit"
   "net/http"
   "os"
   "sort"
   "strings"
)

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
   // DASH
   dash, err := link.DASH()
   if err != nil {
      panic(err)
   }
   for _, ada := range dash.Period.AdaptationSet {
      reps := ada.Representation
      sort.Slice(reps, func(a, b int) bool {
         return reps[a].ID < reps[b].ID
      })
      for _, rep := range reps {
         if rep.MimeType == "" {
            rep.MimeType = ada.MimeType
         }
         if info {
            fmt.Printf("%+v\n", rep)
         } else {
            err := download(link, rep.BaseURL, rep.MimeType)
            if err != nil {
               panic(err)
            }
            break
         }
      }
   }
   // HLS
   hlss, err := link.HLS()
   if err != nil {
      panic(err)
   }
   if info {
      for _, hls := range hlss {
         fmt.Printf("%+v\n", hls)
      }
   }
}
