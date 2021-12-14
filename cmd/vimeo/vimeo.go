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
   videoID, err := mech.Parse(id)
   if err != nil {
      panic(err)
   }
   con, err := vimeo.NewConfig(videoID)
   if err != nil {
      panic(err)
   }
   // info
   if info {
      fmt.Printf("%+v\n", con)
      return
   }
   // download
   for _, f := range con.Request.Files.Progressive {
      if f.Height == height {
         err := download(con, f.URL)
         if err != nil {
            panic(err)
         }
      }
   }
}

func download(con *vimeo.Config, addr string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := con.Video.Owner.Name + "-" + con.Video.Title + path.Ext(addr)
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
