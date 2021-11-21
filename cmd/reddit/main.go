package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/reddit"
   "os"
)

func main() {
   cDASH := choice{
      ids: make(map[string]bool),
   }
   flag.BoolVar(&cDASH.format, "df", false, "DASH formats")
   flag.Func("d", "DASH IDs", func(id string) error {
      cDASH.ids[id] = true
      return nil
   })
   cHLS := choice{
      ids: make(map[string]bool),
   }
   flag.BoolVar(&cHLS.format, "hf", false, "HLS formats")
   flag.Func("h", "HLS IDs", func(id string) error {
      cHLS.ids[id] = true
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
   if cDASH.format || len(cDASH.ids) >= 1 {
      err := cDASH.DASH(link)
      if err != nil {
         panic(err)
      }
   }
   if cHLS.format || len(cHLS.ids) >= 1 {
      err := cHLS.HLS(link)
      if err != nil {
         panic(err)
      }
   }
}
