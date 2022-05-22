package paramount

import (
   "net/http"
   "net/url"
   "strconv"
)

type Address struct {
   sid string
   aid int64
   guid string
}

func NewAddress(guid string) Address {
   return Address{sid: "dJ5BDC", aid: 2198311517, guid: guid}
}

func (a Address) String() string {
   var buf []byte
   buf = append(buf, "http://link.theplatform.com/s/"...)
   buf = append(buf, a.sid...)
   buf = append(buf, "/media/guid/"...)
   buf = strconv.AppendInt(buf, a.aid, 10)
   buf = append(buf, '/')
   buf = append(buf, a.guid...)
   return string(buf)
}

func (a Address) media(guid, formats, asset string) (*url.URL, error) {
   req, err := http.NewRequest("HEAD", a.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "assetTypes": {asset},
      "formats": {formats},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   return res.Location()
}
