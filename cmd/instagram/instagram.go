package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "net/url"
   "os"
   "path"
   "time"
)

func newMedia(shortcode string, auth bool) (*instagram.Media, error) {
   if auth {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      log, err := instagram.OpenLogin(cache + "/mech/instagram.json")
      if err != nil {
         return nil, err
      }
      return log.Media(shortcode)
   }
   return instagram.NewMedia(shortcode)
}

func download(addr string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   par, err := url.Parse(addr)
   if err != nil {
      return err
   }
   file, err := os.Create(path.Base(par.Path))
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
      if info {
         fmt.Println(med)
      } else {
         for _, addr := range med.URLs() {
            err := download(addr)
            if err != nil {
               panic(err)
            }
            time.Sleep(99 * time.Millisecond)
         }
      }
   }
}
