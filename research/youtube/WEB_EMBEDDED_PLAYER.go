package main

import (
   "fmt"
   "io"
   "net/http"
   "strings"
)

func main() {
   res, err := http.Get("https://www.youtube.com/embed/MIchMEqVwvg")
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}
