package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/bbc"
   "github.com/89z/parse/m3u"
   "net/http"
   "os"
   "strconv"
   "strings"
)

func (c choice) HLS(con *bbc.Connection) error {
   forms, err := con.HLS()
   if err != nil {
      return err
   }
   for _, form := range forms {
      if c.format {
         if !strings.HasPrefix(form.URI.File, "keyframes/") {
            fmt.Printf("%+v\n", form)
         }
      } else {
         fmt.Println("GET", form.URI)
         res, err := http.Get(form.URI.String())
         if err != nil {
            return err
         }
         defer res.Body.Close()
         forms, err := m3u.Decode(res.Body, form.URI.Dir)
         if err != nil {
            return err
         }
         for _, form := range forms {
            fmt.Println("GET", form.URI)
            res, err := http.Get(form.URI.String())
            if err != nil {
               return err
            }
            defer res.Body.Close()
            file, err := os.Create(form.URI.File)
            if err != nil {
               return err
            }
            defer file.Close()
            if _, err := file.ReadFrom(res.Body); err != nil {
               return err
            }
         }
      }
   }
   return nil
}

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
   if !cHLS.format && len(cHLS.ids) == 0 {
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


