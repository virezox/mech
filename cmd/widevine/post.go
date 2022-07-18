package main

import (
   "github.com/89z/rosso/http"
   "strings"
)

type flags struct {
   address string
   client_ID string
   header string
   key_ID string
   private_key string
   verbose bool
}

func (f flags) Request_URL() string {
   return f.address
}

func (f flags) Request_Header() http.Header {
   head := make(http.Header)
   key, val, ok := strings.Cut(f.header, ":")
   if ok {
      head.Set(key, val)
   }
   return head
}

func (flags) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (flags) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
