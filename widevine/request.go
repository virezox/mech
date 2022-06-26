package widevine

import (
   "net/http"
)

type Requester interface {
   Header() http.Header
   Key_ID() []byte
   URL() string
}
