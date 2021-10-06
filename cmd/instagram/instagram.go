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
      auth, info bool
      username, password string
   )
   flag.BoolVar(&auth, "a", false, "use authentication")
   flag.BoolVar(&info, "i", false, "info only")
   flag.StringVar(&username, "u", "", "username")
   flag.StringVar(&password, "p", "", "password")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("instagram [flags] [shortcode]")
      flag.PrintDefaults()
      return
   }
   instagram.Verbose(true)
   if username != "" {
      log, err := instagram.NewLogin(username, password)
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
      if err := log.Encode(file); err != nil {
         panic(err)
      }
      return
   }
   shortcode := flag.Arg(0)
   err := instagram.Valid(shortcode)
   if err != nil {
      panic(err)
   }
   var log instagram.Login
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
      if err := log.Decode(file); err != nil {
         panic(err)
      }
   }
   car, err := instagram.GraphQL(shortcode, &log)
   if err != nil {
      panic(err)
   }
   if info {
      fmt.Printf("%+v", car.Shortcode_Media)
      return
   }
   // download video
   if car.Shortcode_Media.Video_URL != "" {
      err := download(car.Shortcode_Media.Video_URL)
      if err != nil {
         panic(err)
      }
      return
   }
   // download image
   if car.Edges() == nil {
      err := download(car.Shortcode_Media.Display_URL)
      if err != nil {
         panic(err)
      }
      return
   }
   // download images
   for _, edge := range car.Edges() {
      err := download(edge.Node.Display_URL)
      if err != nil {
         panic(err)
      }
      time.Sleep(100 * time.Millisecond)
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
