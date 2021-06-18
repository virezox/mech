package main

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "strings"
)

func main() {
   req, err := http.NewRequest("GET", "http://example.com", nil)
   if err != nil {
      panic(err)
   }
   r := strings.NewReader("player=%7B%22month%22%3A12%2C%22day%22%3A31%7D")
   all, err := io.ReadAll(r)
   if err != nil {
      panic(err)
   }
   req.URL.RawQuery = string(all)
   play := req.URL.Query().Get("player")
   var date struct { Month, Day int }
   json.Unmarshal([]byte(play), &date)
   fmt.Printf("%+v\n", date)
}
