package paramount

import (
   "fmt"
   "net/http"
   "strings"
   "github.com/89z/format"
   "github.com/89z/format/hls"
)

var logLevel format.LogLevel

func master() (*hls.Master, error) {
   var buf strings.Builder
   buf.WriteString("https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517")
   buf.WriteString("/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   // We need "MPEG4", otherwise you get a "EXT-X-KEY" with "skd" scheme:
   req.URL.RawQuery = "formats=MPEG4,M3U"
   logLevel.Dump(req)
   // This redirects:
   res, err := new(http.Client).Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewMaster(res.Request.URL, res.Body)
}

func segment() (*hls.Segment, error) {
   mas, err := master()
   if err != nil {
      return nil, err
   }
   fmt.Println(mas.Stream[0].URI)
   res, err := http.Get(mas.Stream[0].URI)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewSegment(res.Request.URL, res.Body)
}
