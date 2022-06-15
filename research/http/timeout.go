package main

import (
   "errors"
   "fmt"
   "net/http"
   "net/http/httptest"
   "os"
   "time"
)

func Three(s string) error {
   req, err := http.NewRequest("GET", s, nil)
   if err != nil {
      return err
   }
   tr := http.Transport{IdleConnTimeout: time.Second}
   res, err := tr.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   if _, err := os.Stdout.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

func main() {
   handler := func(w http.ResponseWriter, r *http.Request) {
      for range [9]bool{} {
         fmt.Fprintln(w, "fail")
         w.(http.Flusher).Flush()
         time.Sleep(9*time.Second)
      }
   }
   ts := httptest.NewServer(http.HandlerFunc(handler))
   err := Three(ts.URL)
   if err != nil {
      panic(err)
   }
}
