package crackle

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
)

var Client format.Client

type Media struct {
   MediaURLs []struct {
      Type string
      DrmPath string
   }
}

func New_Media(id int64) (*Media, error) {
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
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

type authorization struct {
   hmac string
   timestamp int64
   partner_id int64
}

var chunk = authorization{
   hmac: "FF",
   timestamp: 9,
   partner_id: 117,
}

func (a authorization) String() string {
   var buf []byte
   buf = append(buf, a.hmac...)
   buf = append(buf, '|')
   buf = strconv.AppendInt(buf, a.timestamp, 10)
   buf = append(buf, '|')
   buf = strconv.AppendInt(buf, a.partner_id, 10)
   return string(buf)
}
