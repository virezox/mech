package mech

import (
   "compress/gzip"
   "io"
   "net/http"
   "strings"
)

const (
   StatusOK = http.StatusOK
   StatusPartialContent = http.StatusPartialContent
)

var (
   ErrNoCookie = http.ErrNoCookie
   NewRequest = http.NewRequest
)

type (
   Cookie = http.Cookie
   Response = http.Response
)

type Transport struct {
   http.Transport
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if !t.DisableCompression {
      req.Header.Set("Accept-Encoding", "gzip")
   }
   res, err := t.Transport.RoundTrip(req)
   if err != nil {
      return nil, err
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
