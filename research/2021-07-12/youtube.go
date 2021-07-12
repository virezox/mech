package main

import (
   "net/http"
   "os"
   "strings"
)

const body = `
{
   "context": {
      "client": {
         "clientName": "TVHTML5",
         "clientVersion": "6.20180913",
         "originalUrl": "https://www.youtube.com/watch?v=bO7PgQ-DtZk"
      }
   },
   "videoId": "bO7PgQ-DtZk"
}
`

func main() {
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player",
      strings.NewReader(body),
   )
   if err != nil {
      panic(err)
   }
   q := req.URL.Query()
   q.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = q.Encode()
   req.Header.Set("X-Origin","https://www.youtube.com")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}


