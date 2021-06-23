package main
import "net/http"

func main() {
   req, err := http.NewRequest("GET", "https://github.com/manifest.json", nil)
   if err != nil {
      panic(err)
   }
   req.Header.Set("Accept-Encoding", "gzip")
   res, err := new(http.Client).Do(req)
   if err != nil {
      panic(err)
   }
   println(res.ContentLength)
}
