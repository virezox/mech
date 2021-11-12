package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/vimeo"
   "net/http"
   "os"
   "path"
   "strings"
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
      fmt.Println("vimeo [flags] [video ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   if ! vimeo.Valid(id) {
      panic("invalid ID")
   }
   mech.Verbose = true
   cfg, err := vimeo.NewConfig(id)
   if err != nil {
      panic(err)
   }
   // info
   if info {
      for _, f := range cfg.Request.Files.Progressive {
         fmt.Printf("%+v\n", f)
      }
      return
   }
   // download
   for _, f := range cfg.Request.Files.Progressive {
      if f.Height == height {
         err := download(cfg, f.URL)
         if err != nil {
            panic(err)
         }
      }
   }
}

func download(cfg *vimeo.Config, addr string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := cfg.Video.Owner.Name + "-" + cfg.Video.Title + path.Ext(addr)
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
