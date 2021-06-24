package main

import (
   "compress/gzip"
   "fmt"
   "io"
   "net/http"
   "os"
)

type readCloser struct {
   io.Reader
   io.Closer
}

func get(addr string) (*http.Response, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Accept-Encoding", "gzip")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   if res.Header.Get("Content-Encoding") == "gzip" {
      gz, err := gzip.NewReader(res.Body)
      if err != nil {
         return nil, err
      }
      res.Body = readCloser{gz, res.Body}
   }
   return res, nil
}

func main() {
   res, err := get("https://github.com/manifest.json")
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Println(res.ContentLength)
   os.Stdout.ReadFrom(res.Body)
}
