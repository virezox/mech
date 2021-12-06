package nbc

import (
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "strconv"
)

const accountID = 2410887629

type media struct {
   location string
}

func (m media) video() ([]m3u.Format, error) {
   req, err := http.NewRequest("GET", m.location, nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return m3u.Decode(res.Body, "")
}

func newMedia(guid int64) (*media, error) {
   addr := []byte("http://link.theplatform.com/s/NnzsPC/media/guid/")
   addr = strconv.AppendInt(addr, accountID, 10)
   addr = append(addr, '/')
   addr = strconv.AppendInt(addr, guid, 10)
   addr = append(addr, "?manifest=m3u"...)
   req, err := http.NewRequest(
      "HEAD", string(addr), nil,
   )
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   var med media
   med.location = res.Header.Get("Location")
   return &med, nil
}

