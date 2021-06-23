package main
import "net/http"

func main() {
   {
      r, err := http.Head("https://github.com/manifest.json")
      if err != nil {
         panic(err)
      }
      println(r.Uncompressed)
   }
   {
      r, err := http.Get("https://github.com/manifest.json")
      if err != nil {
         panic(err)
      }
      println(r.Uncompressed)
   }
}
