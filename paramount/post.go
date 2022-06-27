package paramount

import (
   "net/http"
)

type Session struct {
   URL string
   LS_Session string
}

func (s Session) Request_URL() string {
   return s.URL
}

func (s Session) Request_Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + s.LS_Session)
   return head
}

func (s Session) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (s Session) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
