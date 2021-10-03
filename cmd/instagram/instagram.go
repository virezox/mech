package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
   "os"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("instagram [-i] [shortcode]")
      flag.PrintDefaults()
      return
   }
   shortcode := flag.Arg(0)
   err := instagram.Valid(shortcode)
   if err != nil {
      panic(err)
   }
   instagram.Verbose = true
}
