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
   // We need "MPEG4", otherwise you get a "EXT-X-KEY" with "skd" scheme:
   req.URL.RawQuery = "formats=MPEG4,M3U"
   // This redirects:
   return new(http.Client).Do(req)
}
