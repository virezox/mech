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
   all := new(strings.Builder)
   if _, err := io.Copy(all, r); err != nil {
      panic(err)
   }
   req.URL.RawQuery = all.String()
   play := req.URL.Query().Get("player")
   var date struct { Month, Day int }
   json.Unmarshal([]byte(play), &date)
   fmt.Printf("%+v\n", date)
}
