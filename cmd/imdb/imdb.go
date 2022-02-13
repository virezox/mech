package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/imdb"
   "net/http"
   "os"
   "path"
   "time"
)

type choice struct {
   info bool
   sleep time.Duration
   width int64
}

func (c choice) gallery(rgconst string) error {
   cred, err := imdb.NewCredential()
   if err != nil {
      return err
   }
   gallery, err := cred.Gallery(rgconst)
   if err != nil {
      return err
   }
   for _, img := range gallery.Images {
      addr := img.Format(c.width)
      if c.info {
         fmt.Println(addr)
      } else {
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         file, err := os.Create(path.Base(addr))
         if err != nil {
            return err
         }
         pro := format.NewProgress(res)
         if _, err := file.ReadFrom(pro); err != nil {
            return err
         }
         // We want to close manually so we dont have to wait for everything to
         // download:
         file.Close()
         time.Sleep(c.sleep)
      }
   }
   return nil
}

func main() {
   var c choice
   // i
   flag.BoolVar(&c.info, "i", false, "info only")
   // r
   var rgconst string
   flag.StringVar(&rgconst, "r", "", "runway gallery ID")
   // s
   flag.DurationVar(&c.sleep, "s", time.Second, "sleep")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   // w
   flag.Int64Var(&c.width, "w", 800, "width")
   flag.Parse()
   if verbose {
      imdb.LogLevel = 1
   }
   if rgconst != "" {
      err := c.gallery(rgconst)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
