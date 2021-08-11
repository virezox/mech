package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
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
   input := flag.Arg(0)
   f, err := os.Open(input)
   if err != nil {
      panic(err)
   }
   defer f.Close()
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
   if err := e.Encode(f); err != nil {
      panic(err)
   }
}
