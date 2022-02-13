package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/tiktok"
)

func main() {
   // a
   var awemeID int64
   flag.Int64Var(&awemeID, "a", 0, "aweme ID")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      tiktok.LogLevel = 1
   }
   if awemeID >= 1 {
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
         err := get(det)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
