package main

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "os"
   "strings"
)

func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   file, err := os.Open(cache + "/mech/youtube.json")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   var token struct {
      Access_Token string
   }
   json.NewDecoder(file).Decode(&token)
   req, err := http.NewRequest("GET", "https://studio.youtube.com", nil)
   if err != nil {
      panic(err)
   }
   req.URL.RawQuery = "approve_browser_access=true"
   req.Header.Set("Authorization", "Bearer " + token.Access_Token)
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
