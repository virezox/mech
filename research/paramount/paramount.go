package paramount

import (
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
   "text/scanner"
)

// paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
// paramountplus.com/shows/bull/video/TUT_4UVB87huHEOfPCjMkxOW_Xe1hNWw/bull-gone
func VideoID(addr string) string {
   var buf scanner.Scanner
   buf.Init(strings.NewReader(addr))
   buf.IsIdentRune = func(r rune, i int) bool {
      return r == '-' || r > '/'
   }
   for buf.Scan() != scanner.EOF {
      switch buf.TokenText() {
      case "video":
         buf.Scan(); buf.Scan()
         return buf.TokenText()
      case "movies":
         buf.Scan(); buf.Scan(); buf.Scan(); buf.Scan()
         return buf.TokenText()
      }
   }
   return ""
}

const endpoint = "/s/dJ5BDC/media/guid/2198311517/"

var LogLevel format.LogLevel

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
