package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/tiktok"
   "net/http"
   "os"
   "strings"
)

func main() {
   var info, verbose bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("tiktok [flags] [aweme ID]")
      flag.PrintDefaults()
      return
   }
   if verbose {
      tiktok.LogLevel = 1
   }
   id := flag.Arg(0)
   awemeID, err := tiktok.Parse(id)
   if err != nil {
      panic(err)
   }
   det, err := tiktok.NewAwemeDetail(awemeID)
   if err != nil {
      panic(err)
   }
   if info {
      fmt.Println("Author:", det.Author.Unique_ID)
      fmt.Println("Create_Time:", det.Time())
      fmt.Println("Duration:", det.Duration())
      fmt.Println("Width:", det.Video.Play_Addr.Width)
      fmt.Println("Height:", det.Video.Play_Addr.Height)
      fmt.Println("URL_List:")
      for _, loc := range det.Video.Play_Addr.URL_List {
         fmt.Println("-", loc)
      }
   } else {
      err := get(det)
      if err != nil {
         panic(err)
      }
   }
}

func get(det *tiktok.AwemeDetail) error {
   addr, err := det.URL()
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   ext, err := mech.ExtensionByType(res.Header.Get("Content-Type"))
   if err != nil {
      return err
   }
   name := det.Author.Unique_ID + "-" + det.Aweme_ID + ext
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
