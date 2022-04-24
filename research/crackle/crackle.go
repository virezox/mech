package crackle

import (
   "github.com/89z/format"
   "net/http"
   "strconv"
)

var LogLevel format.LogLevel

func media(id int64) (*http.Response, error) {
   buf := []byte("http://web-api-us.crackle.com/Service.svc/details/media/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, "/US"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "Authorization": {"155A5C14B97EB77F13C66D081B518A6B3E13722F|202204240234|117|1"},
   }
   req.URL.RawQuery = "disableProtocols=true"
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
