package pbs

import (
   "github.com/89z/format/json"
   "html"
   "net/http"
   "net/url"
)

type Frontline struct {
   Graph []struct {
      Type string `json:"@type"`
      EmbedURL string
   } `json:"@graph"`
}

func NewFrontline(addr string) (*Frontline, error) {
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
   var (
      line = new(Frontline)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Decode(res.Body, sep, line); err != nil {
      return nil, err
   }
   return line, nil
}

func (f Frontline) Widget() (*Widget, error) {
   for _, graph := range f.Graph {
      if graph.Type == "VideoObject" {
         addr, err := url.Parse(html.UnescapeString(graph.EmbedURL))
         if err != nil {
            return nil, err
         }
         return NewWidget(addr)
      }
   }
   return nil, notFound{"VideoObject"}
}
