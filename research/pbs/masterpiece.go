package pbs

import (
   "bytes"
   "github.com/89z/format/json"
   "html"
   "io"
   "net/http"
   "net/url"
)

type Masterpiece struct {
   EmbedURL string
}

func NewMasterpiece(addr string) (*Masterpiece, error) {
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
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   // pbs.org/wgbh/masterpiece/episodes/downton-abbey-s6-e2
   buf = bytes.ReplaceAll(buf, []byte{'\n'}, nil)
   var (
      master = new(Masterpiece)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Unmarshal(buf, sep, master); err != nil {
      return nil, err
   }
   return master, nil
}

func (m Masterpiece) Widget() (*Widget, error) {
   addr, err := url.Parse(html.UnescapeString(m.EmbedURL))
   if err != nil {
      return nil, err
   }
   return NewWidget(addr)
}
