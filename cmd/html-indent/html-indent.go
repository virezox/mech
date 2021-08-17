package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/html"
   "os"
)

func main() {
   var indent, output string
   flag.StringVar(&indent, "i", "", "indent")
   flag.StringVar(&output, "o", "", "output file")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("html-indent [flags] [input file]")
      flag.PrintDefaults()
      return
   }
   arg := flag.Arg(0)
   f, err := os.Open(arg)
   if err != nil {
      panic(err)
   }
   defer f.Close()
   var enc html.Encoder
   if output != "" {
      f, err := os.Create(output)
      if err != nil {
         panic(err)
      }
      defer f.Close()
      enc = html.NewEncoder(f)
   } else {
      enc = html.NewEncoder(os.Stdout)
   }
   enc.SetIndent(indent)
   if err := enc.Encode(f); err != nil {
      panic(err)
   }
}
