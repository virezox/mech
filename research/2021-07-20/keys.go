package main

import (
   "fmt"
   "net/http"
   "os"
   "strings"
   "time"
)

const body = `
{
  "videoId": "Cr381pDsSsA",
  "context": {
    "client": {
      "clientName": "TVHTML5",
      "clientVersion": "7.20210713.10.00"
    }
  }
}
`

var keys = []string{
   "AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w",
   "AIzaSyC8UYZpvA2eknNex0Pjid0_eTLJoDu6los",
   "AIzaSyCjc_pVEDi4qsv5MtC2dMXzpIaDoRFLsxw",
   "AIzaSyCtkvNIR1HCEwzsqK6JuE6KqpyjusIRI30",
   "AIzaSyDHQ9ipnphqTzDqZsbtd8_Ru4_kiKVQe2k",
}

func main() {
   for _, key := range keys {
      req, err := http.NewRequest(
         "POST", "https://www.youtube.com/youtubei/v1/player",
         strings.NewReader(body),
      )
      if err != nil {
         panic(err)
      }
      q := req.URL.Query()
      q.Set("key", key)
      req.URL.RawQuery = q.Encode()
      fmt.Println(req.Method, req.URL)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      f, err := os.Create(key + ".json")
      if err != nil {
         panic(err)
      }
      defer f.Close()
      f.ReadFrom(res.Body)
      time.Sleep(100 *time.Millisecond)
   }
}
