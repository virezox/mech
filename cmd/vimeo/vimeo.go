package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/vimeo"
   "net/http"
   "os"
   "path"
   "strings"
)

func download(vid vimeo.MasterVideo, con *vimeo.Config) error {
   addr, err := vid.URL(con)
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := con.Video.Owner.Name + "-" + con.Video.Title + path.Ext(addr)
   file, err := os.Create(strings.Map(format.Clean, name))
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
      info, verbose bool
      formatID string
   )
   flag.StringVar(&formatID, "f", "", "format")
   flag.BoolVar(&info, "i", false, "info only")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      format.Log.Level = 1
   }
   if flag.NArg() == 1 {
      id := flag.Arg(0)
      videoID, err := vimeo.Parse(id)
      if err != nil {
         panic(err)
      }
      con, err := vimeo.NewConfig(videoID)
      if err != nil {
         panic(err)
      }
      mas, err := con.Master()
      if err != nil {
         panic(err)
      }
      if info {
         fmt.Println(con.Video)
      }
      for _, vid := range mas.Video {
         if info {
            fmt.Print("ID:", vid.ID)
            fmt.Print(" Width:", vid.Width)
            fmt.Print(" Height:", vid.Height)
            fmt.Println()
         } else if vid.ID == formatID {
            err := download(vid, con)
            if err != nil {
               panic(err)
            }
         }
      }
   } else {
      fmt.Println("vimeo [flags] [video ID]")
      flag.PrintDefaults()
   }
}
