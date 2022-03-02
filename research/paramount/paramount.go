package paramount

import (
   "net/http"
   "net/url"
   "strings"
)

func master() (*http.Response, error) {
   var buf strings.Builder
   buf.WriteString("http://link.theplatform.com/s/dJ5BDC/media/guid/2198311517")
   buf.WriteString("/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "format": {"SMIL"},
      "formats": {"MPEG4,M3U"},
   }.Encode()
   return new(http.Client).Do(req) // this redirects
}
