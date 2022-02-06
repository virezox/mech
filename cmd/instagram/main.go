package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
   "os"
   "time"
)

func main() {
   var (
      info, verbose bool
      username, password string
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.StringVar(&password, "p", "", "password")
   flag.StringVar(&username, "u", "", "username")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      instagram.LogLevel = 1
   }
   switch {
   case len(os.Args) == 1:
      fmt.Println("instagram [flags] [shortcode]")
      flag.PrintDefaults()
   case username != "":
      err := saveLogin(username, password)
      if err != nil {
         panic(err)
      }
   default:
      shortcode := flag.Arg(0)
      if !instagram.Valid(shortcode) {
         panic("invalid shortcode")
      }
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      login, err := instagram.OpenLogin(cache + "/mech/instagram.json")
      if err != nil {
         panic(err)
      }
      items, err := login.MediaItems(shortcode)
      if err != nil {
         panic(err)
      }
      for _, item := range items {
         if info {
            form, err := item.Format()
            if err != nil {
               panic(err)
            }
            fmt.Println(form)
         } else {
            for _, info := range item.Infos() {
               addrs, err := info.URLs()
               if err != nil {
                  panic(err)
               }
               for _, addr := range addrs {
                  err := download(addr)
                  if err != nil {
                     panic(err)
                  }
                  time.Sleep(99 * time.Millisecond)
               }
            }
         }
      }
   }
}
