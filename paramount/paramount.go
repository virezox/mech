package paramount

import (
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "text/scanner"
)

const (
   aid = 2198311517
   sid = "dJ5BDC"
)

var LogLevel format.LogLevel

func GUID(addr string) string {
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

func Media(guid string) (*url.URL, error) {
   buf := []byte("https://link.theplatform.com/s/")
   buf = append(buf, sid...)
   buf = append(buf, "/media/guid/"...)
   buf = strconv.AppendInt(buf, aid, 10)
   buf = append(buf, '/')
   buf = append(buf, guid...)
   req, err := http.NewRequest("GET", string(buf), nil)
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
