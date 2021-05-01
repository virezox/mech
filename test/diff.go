package main

import (
   "github.com/89z/youtube"
   "net/url"
   "strings"
)

func split(s string) ([]string, error) {
   v, e := youtube.NewVideo("BnEn7X3Pr7o")
   if e != nil { return nil, e }
   println(v.StreamingData.DashManifestURL)
   p, e := url.Parse(v.StreamingData.DashManifestURL)
   if e != nil { return nil, e }
   return strings.Split(p.Path, "/")[1:], nil
}

func main() {
   s := "BnEn7X3Pr7o"
   a, e := split(s)
   if e != nil {
      panic(e)
   }
   b, e := split(s)
   if e != nil {
      panic(e)
   }
   for n := range a {
      if a[n] != b[n] {
         switch a[n-1] {
         case "ei", "expire", "initcwndbps", "lsig", "mm", "mn", "ms", "sig":
         case "mt": continue
         }
         println(a[n-1])
         println(a[n])
         println(b[n])
         println()
      }
   }
}
