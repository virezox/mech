package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "time"
)

var els = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/Hexer10/youtube_explode_dart
   "embedded",
}

func request(req *http.Request) error {
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf("status %v", res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   date := []byte("publishDate")
   if ! bytes.Contains(body, date) {
      return fmt.Errorf("missing %q", date)
   }
   return nil
}

func values() error {
   req, err := http.NewRequest(
      "GET", "https://www.youtube.com/get_video_info", nil,
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("video_id", "NMYIVsdGfoo")
   // stackoverflow.com/a/67629882
   val.Set("eurl", "https://youtube.googleapis.com/v/NMYIVsdGfoo")
   // stackoverflow.com/a/67629882
   val.Set("html5", "1")
   // stackoverflow.com/a/67629882
   val.Set("c", "TVHTML5")
   // stackoverflow.com/a/67629882
   val.Set("cver", "6.20180913")
   for _, el := range els {
      if el != "" {
         val.Set("el", el)
      }
      req.URL.RawQuery = val.Encode()
      err := request(req)
      fmt.Println(val, err)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}

func main() {
   err := values()
   if err != nil {
      panic(err)
   }
}
