package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/tiktok"
   "net/http"
   "os"
   "strings"
)

func main() {
   var (
      info, verbose bool
      output string
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.StringVar(&output, "o", "", "output")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("tiktok [flags] [aweme ID]")
      flag.PrintDefaults()
      return
   }
   if verbose {
      format.Log.Level = 1
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
      for _, addr := range det.Video.Play_Addr.URL_List {
         fmt.Println("-", addr)
      }
   } else {
      err := get(det, output)
      if err != nil {
         panic(err)
      }
   }
}

func get(det *tiktok.AwemeDetail, output string) error {
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
   if output == "" {
      ext, err := format.ExtensionByType(res.Header.Get("Content-Type"))
      if err != nil {
         return err
      }
      name := det.Author.Unique_ID + "-" + det.Aweme_ID + ext
      output = strings.Map(format.Clean, name)
   }
   file, err := os.Create(output)
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
