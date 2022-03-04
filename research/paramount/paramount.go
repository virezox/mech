package paramount

import (
   "net/http"
   "net/url"
   "strings"
   "github.com/89z/format"
)

const endpoint = "/s/dJ5BDC/media/guid/2198311517/"

var LogLevel format.LogLevel

// paramountplus.com/shows/star-trek-prodigy/video/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy/star-trek-prodigy-dreamcatcher
func VideoID(addr string) string {
   var prev string
   for _, split := range strings.Split(addr, "/") {
      if prev == "video" {
         return split
      }
      prev = split
   }
   return ""
}

func Media(videoID string) (*url.URL, error) {
   var buf strings.Builder
   buf.WriteString("https://link.theplatform.com")
   buf.WriteString(endpoint)
   buf.WriteString(videoID)
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   // We need "MPEG4", otherwise you get a "EXT-X-KEY" with "skd" scheme:
   req.URL.RawQuery = "formats=MPEG4,M3U"
   LogLevel.Dump(req)
   // This redirects, but we only care about the URL, so we dont need to follow:
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return res.Location()
}
