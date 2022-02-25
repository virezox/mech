package bandcamp

import (
   "encoding/json"
   "github.com/89z/format/net"
   "html"
   "net/http"
   "strconv"
   "strings"
)

func NewDataTralbum(addr string) (*DataTralbum, error) {
   contains := func(s string) bool {
      return strings.Contains(addr, s)
   }
   if !contains("/album/") && !contains("/track/") {
      return nil, notFound{"/album/,/track/"}
   }
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   for _, node := range net.ReadHTML(res.Body, "script") {
      data, ok := node.Attr["data-tralbum"]
      if ok {
         data = html.UnescapeString(data)
         tra := new(DataTralbum)
         err := json.Unmarshal([]byte(data), tra)
         if err != nil {
            return nil, err
         }
         return tra, nil
      }
   }
   return nil, notFound{"data-tralbum"}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " not found"
}
