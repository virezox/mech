package pbs

import (
   "bytes"
   "github.com/89z/format/json"
   "html"
   "io"
   "net/http"
   "net/url"
)

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
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   // pbs.org/wgbh/masterpiece/episodes/downton-abbey-s6-e2
   buf = bytes.ReplaceAll(buf, []byte{'\n'}, nil)
   var (
      line = new(Frontline)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Unmarshal(buf, sep, line); err != nil {
      return nil, err
   }
   return line, nil
}

type Frontline struct {
   Graph []Object `json:"@graph"`
   Object
}

type Object struct {
   ContentURL string
   EmbedURL string
   Type string `json:"@type"`
}

func (f Frontline) VideoObject() Object {
   for _, graph := range f.Graph {
      if graph.Type == "VideoObject" {
         return graph
      }
   }
   return f.Object
}

func (o Object) VideoBridge() (*VideoBridge, error) {
   addr := o.ContentURL
   if addr == "" {
      addr = o.EmbedURL
   }
   parse, err := url.Parse(html.UnescapeString(addr))
   if err != nil {
      return nil, err
   }
   return NewVideoBridge(parse)
}
