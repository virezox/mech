package mech

import (
   "compress/gzip"
   "fmt"
   "io"
   "net/http"
   "strings"
)

const (
   StatusOK = http.StatusOK
   StatusPartialContent = http.StatusPartialContent
   StatusSeeOther = http.StatusSeeOther
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

var (
   ErrNoCookie = http.ErrNoCookie
   HandleFunc = http.HandleFunc
   NewRequest = http.NewRequest
   Redirect = http.Redirect
)

func Get(addr string) (*http.Response, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   return new(Transport).RoundTrip(req)
}

type (
   Cookie = http.Cookie
   Request = http.Request
   Response = http.Response
   ResponseWriter = http.ResponseWriter
   Server = http.Server
)

type Transport struct {
   http.Transport
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if !t.DisableCompression {
      req.Header.Set("Accept-Encoding", "gzip")
   }
   bytes := req.Header.Get("Range")
   if bytes == "" || strings.HasPrefix(bytes, "bytes=0-") {
      fmt.Println(invert, req.Method, reset, req.URL)
   }
   if bytes != "" {
      fmt.Println(bytes)
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
