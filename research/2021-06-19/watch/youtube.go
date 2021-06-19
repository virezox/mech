package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
)

func request() error {
   req, err := http.NewRequest("GET", "https://www.youtube.com/watch", nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("v", "NMYIVsdGfoo")
   req.URL.RawQuery = val.Encode()
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
   count := bytes.Count(body, []byte("videoplayback"))
   println(count)
   return nil
}

func main() {
   err := request()
   if err != nil {
      panic(err)
   }
}
