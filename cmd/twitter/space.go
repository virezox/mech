package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/twitter"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func spacePath(id string, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   space, err := twitter.NewSpace(guest, id)
   if err != nil {
      return err
   }
   stream, err := space.Stream(guest)
   if err != nil {
      return err
   }
   if info {
      fmt.Println("Admins:", space.Admins())
      fmt.Println("Title:", space.Title())
      fmt.Println("Duration:", space.Duration())
      fmt.Println("Location:", stream.Source.Location)
   } else {
      srcs, err := stream.Chunks()
      if err != nil {
         return err
      }
      name := space.Admins() + "-" + space.Title() + ".aac"
      dst, err := os.Create(strings.Map(format.Clean, name))
      if err != nil {
         return err
      }
      defer dst.Close()
      for key, src := range srcs {
         addr, err := url.Parse(src["URI"])
         if err != nil {
            return err
         }
         fmt.Printf("%v/%v %v\n", key, len(srcs), addr.Path)
         res, err := http.Get(addr.String())
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if _, err := dst.ReadFrom(res.Body); err != nil {
            return err
         }
      }
   }
   return nil
}
