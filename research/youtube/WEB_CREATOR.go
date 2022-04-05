package main

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "net/http"
   "os"
   "strings"
)

type token struct {
   Access_Token string
}

func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   tok, err := format.Open[token](cache, "mech/youtube.json")
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest("GET", "https://studio.youtube.com", nil)
   if err != nil {
      panic(err)
   }
   req.URL.RawQuery = "approve_browser_access=true"
   req.Header.Set("Authorization", "Bearer " + tok.Access_Token)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}
