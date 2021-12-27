package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/twitter"
   "net/http"
   "net/url"
   "os"
   "path"
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
      dst, err := os.Create(space.Admins() + "-" + space.Title() + ".aac")
      if err != nil {
         return err
      }
      defer dst.Close()
      for key, src := range srcs {
         loc, err := url.Parse(src["URI"])
         if err != nil {
            return err
         }
         fmt.Printf("%v/%v %v\n", key, len(srcs), loc.Path)
         res, err := http.Get(loc.String())
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

func statusPath(id string, info bool, format int) error {
   nID, err := mech.Parse(id)
   if err != nil {
      return err
   }
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   stat, err := twitter.NewStatus(guest, nID)
   if err != nil {
      return err
   }
   for index, variant := range stat.Variants() {
      loc := variant.URL.String()
      switch {
      case info:
         fmt.Print("ID:", index)
         fmt.Print(" URL:", loc)
         fmt.Println()
      case format == index:
         fmt.Println("GET", loc)
         res, err := http.Get(loc)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         name := stat.User.Name + "-" + id + path.Ext(loc)
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
