package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/tiktok"
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
   if verbose {
      tiktok.LogLevel = 1
   }
   if flag.NArg() == 1 {
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
   } else {
      fmt.Println("tiktok [flags] [aweme ID]")
      flag.PrintDefaults()
   }
}
