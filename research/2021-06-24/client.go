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

type client struct {
   http.Transport
}

func (c client) get(addr string) (*http.Response, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   if !c.DisableCompression {
      req.Header.Set("Accept-Encoding", "gzip")
   }
   res, err := c.RoundTrip(req)
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
   for _, gz := range []bool{false, true} {
      var c client
      c.DisableCompression = gz
      res, err := c.get("https://github.com/manifest.json")
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      fmt.Println(res.ContentLength)
      os.Stdout.ReadFrom(res.Body)
      fmt.Println()
   }
}
