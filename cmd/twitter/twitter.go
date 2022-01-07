package main

import (
   "fmt"
   "github.com/89z/mech/twitter"
   "net/http"
   "net/url"
   "os"
   "path"
   "strings"
)

func statusPath(id, output string, info bool, format int) error {
   nID, err := twitter.Parse(id)
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
      addr := variant.URL.String()
      switch {
      case info:
         fmt.Print("ID:", index)
         fmt.Print(" URL:", addr)
         fmt.Println()
      case format == index:
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         name := filename(output, stat.User.Name, id, addr)
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

func filename(output, name, id, addr string) string {
   var str strings.Builder
   if output != "" {
      str.WriteString(output)
      str.WriteByte('/')
   }
   str.WriteString(name)
   str.WriteByte('-')
   str.WriteString(id)
   str.WriteString(path.Ext(addr))
   return str.String()
}

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
