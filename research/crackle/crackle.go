package crackle

import (
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "net/http"
   "strconv"
)

var LogLevel format.LogLevel

type Media struct {
   MediaURLs []struct {
      Type string
      DrmPath string
   }
}

func NewMedia(id int64) (*Media, error) {
   buf := []byte("http://web-api-us.crackle.com/Service.svc/details/media/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, "/US"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "Authorization": {chunk.String()},
   }
   req.URL.RawQuery = "disableProtocols=true"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

type authorization struct {
   hmac string
   timestamp int64
   partnerID int64
}

var chunk = authorization{
   hmac: "FF",
   timestamp: 9,
   partnerID: 117,
}

func (a authorization) String() string {
   var buf []byte
   buf = append(buf, a.hmac...)
   buf = append(buf, '|')
   buf = strconv.AppendInt(buf, a.timestamp, 10)
   buf = append(buf, '|')
   buf = strconv.AppendInt(buf, a.partnerID, 10)
   return string(buf)
}
