package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/tiktok"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if flag.NArg() == 0 {
      fmt.Println("tiktok [flags] [URL]")
      flag.PrintDefaults()
      return
   }
   addr := flag.Arg(0)
   vid, err := tiktok.NewVideo(addr)
   if err != nil {
      panic(err)
   }
   req, err := tiktok.Request(vid)
   if err != nil {
      panic(err)
   }
   if info {
      buf, err := httputil.DumpRequest(req, false)
      if err != nil {
         panic(err)
      }
      os.Stdout.Write(buf)
   } else {
      err := get(req, vid)
      if err != nil {
         panic(err)
      }
   }
}

func get(req *http.Request, vid tiktok.Video) error {
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   ext, err := mech.ExtensionByType(res.Header.Get("Content-Type"))
   if err != nil {
      return err
   }
   name := vid.Author() + "-" + vid.ID() + ext
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
