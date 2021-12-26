package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/reddit"
   "os"
)

func main() {
   cDASH := choice{
      formats: make(map[string]bool),
   }
   flag.BoolVar(&cDASH.info, "di", false, "DASH info")
   flag.Func("d", "DASH formats", func(id string) error {
      cDASH.formats[id] = true
      return nil
   })
   cHLS := choice{
      formats: make(map[string]bool),
   }
   flag.BoolVar(&cHLS.info, "hi", false, "HLS info")
   flag.Func("h", "HLS formats", func(id string) error {
      cHLS.formats[id] = true
      return nil
   })
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("reddit [flags] [post ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   if !reddit.Valid(id) {
      panic("invalid ID")
   }
   post, err := reddit.NewPost(id)
   if err != nil {
      panic(err)
   }
   link, err := post.Link()
   if err != nil {
      panic(err)
   }
   if cDASH.info || len(cDASH.formats) >= 1 {
      err := cDASH.DASH(link)
      if err != nil {
         panic(err)
      }
   }
   if cHLS.info || len(cHLS.formats) >= 1 {
      err := cHLS.HLS(link)
      if err != nil {
         panic(err)
      }
   }
}
