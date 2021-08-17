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
   f, err := os.Open(flag.Arg(0))
   if err != nil {
      panic(err)
   }
   defer f.Close()
   var w mech.HtmlWriter
   if output != "" {
      f, err := os.Create(output)
      if err != nil {
         panic(err)
      }
      defer f.Close()
      w = mech.NewHtmlWriter(f)
   } else {
      w = mech.NewHtmlWriter(os.Stdout)
   }
   w.SetIndent(indent)
   if err := w.ReadFrom(f); err != nil {
      panic(err)
   }
}
