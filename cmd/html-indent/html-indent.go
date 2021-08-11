package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   var indent, output string
   flag.StringVar(&output, "o", "", "output file")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("html-indent [-o outfile] [infile]")
      flag.PrintDefaults()
      return
   }
   rd, err := os.Open(flag.Arg(0))
   if err != nil {
      panic(err)
   }
   defer rd.Close()
   var e mech.Encoder
   if output != "" {
      f, err := os.Create(output)
      if err != nil {
         panic(err)
      }
      defer f.Close()
      e = mech.NewEncoder(f)
   } else {
      e = mech.NewEncoder(os.Stdout)
   }
   e.SetIndent(indent)
   e.Encode(res.Body)
}
