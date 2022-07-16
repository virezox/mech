package paramount

import (
   "net/http"
)

type Session struct {
   URL string
   LS_Session string
}

func (self Session) Request_URL() string {
   return self.URL
}

func (self Session) Request_Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + self.LS_Session)
   return head
}

func (Session) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Session) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
