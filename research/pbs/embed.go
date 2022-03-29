package pbs

import (
   "bytes"
   "github.com/89z/format/json"
   "html"
   "io"
   "net/http"
   "net/url"
)

type Object struct {
   EmbedURL string
   Type string `json:"@type"`
}

type Embed struct {
   Graph []Object `json:"@graph"`
   Object
}

func NewEmbed(addr string) (*Embed, error) {
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
      embed = new(Embed)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Unmarshal(buf, sep, embed); err != nil {
      return nil, err
   }
   return embed, nil
}

func (e Embed) VideoObject() Object {
   for _, graph := range e.Graph {
      if graph.Type == "VideoObject" {
         return graph
      }
   }
   return e.Object
}

func (o Object) Bridge() (*Bridge, error) {
   parse, err := url.Parse(html.UnescapeString(o.EmbedURL))
   if err != nil {
      return nil, err
   }
   return NewBridge(parse)
}
