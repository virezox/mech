package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/bbc"
   "strconv"
)

type choice struct {
   format bool
   ids map[string]bool
}

func main() {
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
      fmt.Println("bbc [flags] [ID]")
      flag.PrintDefaults()
      return
   }
   sID := flag.Arg(0)
   if !bbc.Valid(sID) {
      panic("invalid ID")
   }
   id, err := strconv.Atoi(sID)
   if err != nil {
      panic(err)
   }
   video, err := bbc.NewNewsVideo(id)
   if err != nil {
      panic(err)
   }
   con, err := video.Connection()
   if err != nil {
      panic(err)
   }
   if err := cHLS.HLS(con); err != nil {
      panic(err)
   }
}

func (c choice) HLS(con *bbc.Connection) error {
   hlss, err := con.HLS()
   if err != nil {
      return err
   }
   for _, hls := range hlss {
      if c.format {
         fmt.Printf("%+v\n", hls)
      } else {
         // FIXME
         fmt.Println("GET", hls.URI)
         res, err := http.Get(hls.URI)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         prefix, _ := path.Split(hls.URI)
         for key := range m3u.NewByteRange(res.Body) {
            fmt.Println("GET", prefix + key)
            res, err := http.Get(prefix + key)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            file, err := os.Create(key)
            if err != nil {
               return err
            }
            defer file.Close()
            file.ReadFrom(res.Body)
         }
      }
   }
   return nil
}
