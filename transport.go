package mech

import (
   "compress/gzip"
   "fmt"
   "io"
   "net/http"
   "strings"
)

type Transport struct {
   http.Transport
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if !t.DisableCompression {
      req.Header.Set("Accept-Encoding", "gzip")
   }
   res, err := t.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   if strings.EqualFold(res.Header.Get("Content-Encoding"), "gzip") {
      gz, err := gzip.NewReader(res.Body)
      if err != nil {
         return nil, err
      }
      res.Body = readCloser{gz, res.Body}
   }
   return res, nil
}

type readCloser struct {
   io.Reader
   io.Closer
}
