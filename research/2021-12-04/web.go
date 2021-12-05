package nbc

import (
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
)

const platform = "http://link.theplatform.com"

func media(guid int) (*http.Response, error) {
   req, err := http.NewRequest(
      "GET",
      platform + "/s/NnzsPC/media/guid/2410887629/" + strconv.Itoa(guid),
      nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "format": {"SMIL"}, // can kill
      "manifest": {"m3u"}, // maybe can kill?
      "mbr": {"true"}, // can kill
   }.Encode()
   mech.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
