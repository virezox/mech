// github.com/89z
package pbs

import (
   "github.com/89z/format/json"
   "html"
   "net/http"
   "net/url"
)

type Masterpiece string

func NewMasterpiece(addr string) (Masterpiece, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return "", err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   scan, err := json.NewScanner(res.Body)
   if err != nil {
      return "", err
   }
   scan.Split = []byte(`"https://video.`)
   scan.Scan()
   var master Masterpiece
   if err := scan.Decode(&master); err != nil {
      return "", err
   }
   return master, nil
}

func (m Masterpiece) Widget() (*Widget, error) {
   raw := html.UnescapeString(string(m))
   addr, err := url.Parse(raw)
   if err != nil {
      return nil, err
   }
   return NewWidget(addr)
}
