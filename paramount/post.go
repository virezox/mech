package paramount

import (
   "net/http"
)

type Session struct {
   URL string
   LS_Session string
}

func (s Session) License_URL() string {
   return s.URL
}

func (s Session) Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + s.LS_Session)
   return head
}

func (s Session) Body(buf []byte) []byte {
   return buf
}
