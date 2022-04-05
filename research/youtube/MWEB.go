package youtube

import (
   "fmt"
   "io"
   "net/http"
   "strings"
)

func MWEB() {
   req, err := http.NewRequest("GET", "https://m.youtube.com", nil)
   if err != nil {
      panic(err)
   }
   req.Header.Set("User-Agent", "iPad")
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
