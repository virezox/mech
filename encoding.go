package mech

import (
   "fmt"
   "net/http"
)

var AcceptEncoding = []string{
   "identity",
   // https://github.com/manifest.json
   "gzip",
   // https://serverpilot.io
   "br",
}

type Content struct {
   Encoding string
   Length string
}

func NewContent(req *http.Request) (map[string]Content, error) {
   con := make(map[string]Content)
   for _, enc := range AcceptEncoding {
      req.Header.Set("Accept-Encoding", enc)
      fmt.Println(req.Method, req.URL)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         return nil, err
      }
      defer res.Body.Close()
      con[enc] = Content{
         res.Header.Get("Content-Encoding"), res.Header.Get("Content-Length"),
      }
   }
   return con, nil
}
