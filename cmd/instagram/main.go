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
      auth, info, verbose bool
      username, password string
   )
   flag.BoolVar(&auth, "a", false, "use authentication")
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
      login, err := instagram.NewLogin(username, password)
      if err != nil {
         panic(err)
      }
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      if err := login.Create(cache + "/mech/instagram.json"); err != nil {
         panic(err)
      }
   default:
      shortcode := flag.Arg(0)
      if !instagram.Valid(shortcode) {
         panic("invalid shortcode")
      }
      med, err := newMedia(shortcode, auth)
      if err != nil {
         panic(err)
      }
      switch {
      case info:
         fmt.Println(med)
      case med.Video_URL != "":
         err := download(med.Video_URL)
         if err != nil {
            panic(err)
         }
      case med.Sidecar() == nil:
         err := download(med.Display_URL)
         if err != nil {
            panic(err)
         }
      default:
         for _, edge := range med.Sidecar() {
            err := download(edge.URL())
            if err != nil {
               panic(err)
            }
            time.Sleep(99 * time.Millisecond)
         }
      }
   }
}
