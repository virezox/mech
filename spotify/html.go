package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   req, err := http.NewRequest(
      "GET", "https://open.spotify.com/playlist/6rZ28nCpmG5Wo1Ik64EoDm", nil,
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("User-Agent", "Firefox/60")
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   b, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(b, []byte(`"accessToken"`)) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
