package pbs

import (
   "bytes"
   "github.com/89z/format/json"
   "io"
   "net/http"
)

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
      line = new(Embed)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Unmarshal(buf, sep, line); err != nil {
      return nil, err
   }
   return line, nil
}

func (e Embed) VideoObject() Object {
   for _, graph := range e.Graph {
      if graph.Type == "VideoObject" {
         return graph
      }
   }
   return e.Object
}
