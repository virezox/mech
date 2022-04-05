package youtube

import (
   "fmt"
   "io"
   "net/http"
   "strings"
)

func WEB() {
   res, err := http.Get("https://www.youtube.com")
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}
