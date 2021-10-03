package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "net/url"
   "os"
   "path"
)

func main() {
   var (
      auth, info bool
      username, password, shortcode string
   )
   flag.BoolVar(&auth, "a", false, "use authentication")
   flag.BoolVar(&info, "i", false, "info only")
   flag.StringVar(&username, "u", "", "username")
   flag.StringVar(&password, "p", "", "password")
   flag.StringVar(&shortcode, "s", "", "shortcode")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("instagram [flags]")
      flag.PrintDefaults()
      return
   }
   instagram.Verbose = true
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
   car, err := instagram.NewQuery(shortcode).Sidecar(&log)
   if err != nil {
      panic(err)
   }
   for _, edge := range car.Edges() {
      if info {
         fmt.Printf("%+v\n", edge)
      } else {
         err := download(edge.Node.Display_URL)
         if err != nil {
            panic(err)
         }
      }
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
