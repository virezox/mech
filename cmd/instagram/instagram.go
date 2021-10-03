package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
   "os"
)

func main() {
   var (
      auth bool
      username, password, shortcode string
   )
   flag.BoolVar(&auth, "a", false, "use authentication")
   flag.StringVar(&username, "u", "", "username")
   flag.StringVar(&password, "p", "", "password")
   flag.StringVar(&shortcode, "s", "", "shortcode")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("instagram [flags]")
      flag.PrintDefaults()
      return
   }
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
      fmt.Printf("%+v\n", edge)
   }
}
