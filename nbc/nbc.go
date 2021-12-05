package nbc

import (
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "net/url"
   "strconv"
)

const accountID = 2410887629

func Media(guid int64) ([]m3u.Format, error) {
   var addr []byte
   addr = append(addr, "http://link.theplatform.com/s/NnzsPC/media/guid/"...)
   addr = strconv.AppendInt(addr, accountID, 10)
   addr = append(addr, '/')
   addr = strconv.AppendInt(addr, guid, 10)
   req, err := http.NewRequest(
      "GET", string(addr), nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "manifest": {"m3u"},
   }.Encode()
   mech.Dump(req)
   // this redirects, so cannot use RoundTrip
   res, err := new(http.Client).Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return m3u.Decode(res.Body, "")
}
