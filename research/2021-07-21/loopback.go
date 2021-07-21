package main
 
import (
   "fmt"
   "net/http"
   "net/http/httputil"
)
 
func main() {
   http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
      b, err := httputil.DumpRequest(req, true)
      if err != nil {
         panic(err)
      }
      w.Write(b)
   })
   port := ":50342"
   fmt.Println("localhost" + port)
   http.ListenAndServe(port, nil)
}
