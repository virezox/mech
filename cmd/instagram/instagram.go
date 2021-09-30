package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "net/url"
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
      fmt.Println("GET", edge.Node.Display_URL)
      res, err := http.Get(edge.Node.Display_URL)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      addr, err := url.Parse(edge.Node.Display_URL)
      if err != nil {
         panic(err)
      }
      file, err := os.Create(path.Base(addr.Path))
      if err != nil {
         panic(err)
      }
      defer file.Close()
      if _, err := file.ReadFrom(res.Body); err != nil {
         panic(err)
      }
   }
}
