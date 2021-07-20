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
   "AIzaSyA64xQnVODx8qBOeSsrlfDc8gDEw_NLopk",
   "AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w",
   "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8",
   "AIzaSyA_n-CBlmsO1fOxFUZqRnQ9SX4Bh1jCjWg",
   "AIzaSyAxmTFlJLw9-uEJ1pFJUzw8LX7veGxGUoI",
   "AIzaSyBD1uN7sPOWjkZ3fNKv7xDlLqF7Rg_JLnk",
   "AIzaSyCChP9IaeaDS_LLGBI0P9CDQwTzCxn1kp8",
   "AIzaSyCTa7aViyHnB3GLIqhL5hQFZGb675SoCIA",
   "AIzaSyCV2I1gEhkJYkd51xG7MGaZGC85zylcS74",
   "AIzaSyCX7NVTCfWMK8eEUau8Scc2y6dZUpWfNd0",
   "AIzaSyCjc_pVEDi4qsv5MtC2dMXzpIaDoRFLsxw",
   "AIzaSyCqrNxCAJrrk_NQqIUp1-baqW05d3JYeOc",
   "AIzaSyCtkvNIR1HCEwzsqK6JuE6KqpyjusIRI30",
   "AIzaSyCymf5PAosq7hWs5DkgHy0-3uacHaY1SPE",
   "AIzaSyD5cCj3gK6IKFQCHRf1pYAt9nDKUzfxmPg",
   "AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8",
   "AIzaSyDHQ9ipnphqTzDqZsbtd8_Ru4_kiKVQe2k",
   "AIzaSyDil7P0s1hvamdVWsqFtySc1T5P1S9dHqk",
   "AIzaSyDjSMHkZSQWmcCKsNnvZcjRc2ZaJbAXpR4",
   "AIzaSyDtpXO8h8u8Z6N7asPxy6AczIICsqmkg64",
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
