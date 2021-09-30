package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "os"
   "path"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("instagram [-i] [ID]")
      flag.PrintDefaults()
      return
   }
   id := flag.Arg(0)
   err := instagram.ValidID(id)
   if err != nil {
      panic(err)
   }
   instagram.Verbose = true
   car, err := instagram.NewSidecar(id)
   if err != nil {
      panic(err)
   }
   for _, edge := range car.Edges() {
      if info {
         fmt.Printf("%+v\n", edge.Node)
         continue
      }
      addr := edge.Node.Display_URL
      fmt.Println("GET", addr)
      res, err := http.Get(addr)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      f, err := os.Create(path.Base(addr))
      if err != nil {
         panic(err)
      }
      defer f.Close()
      if _, err := f.ReadFrom(res.Body); err != nil {
         panic(err)
      }
   }
}
