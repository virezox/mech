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
   if len(os.Args) == 1 {
      fmt.Println("instagram [flags] [shortcode]")
      flag.PrintDefaults()
      return
   }
   if verbose {
      instagram.Log.Level = 1
   }
   if username != "" {
      login, err := instagram.NewLogin(username, password)
      if err != nil {
         panic(err)
      }
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      cache += "/mech"
      os.Mkdir(cache, os.ModeDir)
      file, err := os.Create(cache + "/instagram.json")
      if err != nil {
         panic(err)
      }
      defer file.Close()
      if err := login.Encode(file); err != nil {
         panic(err)
      }
      return
   }
   shortcode := flag.Arg(0)
   if !instagram.Valid(shortcode) {
      panic("invalid shortcode")
   }
   var login instagram.Login
   if auth {
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      file, err := os.Open(cache + "/mech/instagram.json")
      if err != nil {
         panic(err)
      }
      defer file.Close()
      if err := login.Decode(file); err != nil {
         panic(err)
      }
   }
   med, err := login.GraphQL(shortcode)
   if err != nil {
      panic(err)
   }
   if info {
      fmt.Printf("%+v", med.Shortcode_Media)
      return
   }
   // download video
   if med.Shortcode_Media.Video_URL != "" {
      err := download(med.Shortcode_Media.Video_URL)
      if err != nil {
         panic(err)
      }
      return
   }
   // download image
   if med.Edges() == nil {
      err := download(med.Shortcode_Media.Display_URL)
      if err != nil {
         panic(err)
      }
      return
   }
   // download image and video
   for _, edge := range med.Edges() {
      err := download(edge.URL())
      if err != nil {
         panic(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
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
