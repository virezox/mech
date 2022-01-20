package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/twitter"
   "net/http"
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

func main() {
   var (
      info, space, verbose bool
      form int
      output string
   )
   flag.BoolVar(&info, "i", false, "info")
   flag.IntVar(&form, "f", 0, "format")
   flag.StringVar(&output, "o", "", "output")
   flag.BoolVar(&space, "s", false, "space")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      twitter.LogLevel = 1
   }
   if flag.NArg() == 1 {
      id := flag.Arg(0)
      if space {
         err := spacePath(id, info)
         if err != nil {
            panic(err)
         }
      } else {
         err := statusPath(id, output, info, form)
         if err != nil {
            panic(err)
         }
      }
   } else {
      fmt.Println("twitter [flags] [ID]")
      flag.PrintDefaults()
   }
}
