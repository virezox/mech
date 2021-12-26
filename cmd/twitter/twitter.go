package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
   "path"
)

type choice struct {
   info bool
   format int
}

func main() {
   var (
      tweet choice
      verbose bool
   )
   flag.IntVar(&tweet.format, "f", 0, "format")
   flag.BoolVar(&tweet.info, "i", false, "info")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("twitter [flags] [ID]")
      flag.PrintDefaults()
      return
   }
   if verbose {
      twitter.LogLevel = 1
   }
   id := flag.Arg(0)
   if err := tweet.choose(id); err != nil {
      panic(err)
   }
}

func (c choice) choose(id string) error {
   nID, err := mech.Parse(id)
   if err != nil {
      return err
   }
   act, err := twitter.NewActivate()
   if err != nil {
      return err
   }
   stat, err := act.Status(nID)
   if err != nil {
      return err
   }
   for format, vari := range stat.Variants() {
      addr := vari.URL.String()
      switch {
      case c.info:
         fmt.Print("ID:", format)
         fmt.Print(" URL:", addr)
         fmt.Println()
      case c.format == format:
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         name := stat.User.Name + "-" + id + path.Ext(addr)
         dst, err := os.Create(name)
         if err != nil {
            return err
         }
         defer dst.Close()
         if _, err := dst.ReadFrom(res.Body); err != nil {
            return err
         }
      }
   }
   return nil
}
