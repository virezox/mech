package widevine

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

var LogLevel format.LogLevel

type request struct {
   PSSH string `json:"pssh"`
   License string `json:"license"`
   Headers string `json:"headers"`
   Cache bool `json:"cache"`
}

func (r request) post() (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(r)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "http://getwvkeys.cc/api", buf)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
