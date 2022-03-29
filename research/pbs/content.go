package pbs

import (
   "github.com/89z/format/json"
   "html"
   "io"
   "net/http"
   "net/url"
)

type Content struct {
   Graph []Object `json:"@graph"`
   Object
}

func NewContent(addr string) (*Content, error) {
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
   var (
      line = new(Content)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Unmarshal(buf, sep, line); err != nil {
      return nil, err
   }
   return line, nil
}

func (c Content) VideoObject() Object {
   for _, graph := range c.Graph {
      if graph.Type == "VideoObject" {
         return graph
      }
   }
   return c.Object
}

type Object struct {
   ContentURL string
   EmbedURL string
   Type string `json:"@type"`
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
