package pandora

import (
   "github.com/89z/format"
   "github.com/89z/format/net"
   "net/http"
   "net/url"
   "strconv"
)

type musicRecording struct {
   ID string `json:"@id"`
   Name string
   ByArtist struct {
      Name string
   }
}

func newMusicRecording(addr string) (*musicRecording, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   for _, node := range net.ReadHTML(res.Body, "script") {
      if node.Attr["type"] == "application/ld+json" {
         con := node.Attr["content"]
         addr, err := url.Parse(con)
         if err != nil {
            return "", err
         }
         return addr.Query().Get("pandoraId"), nil
      }
   }
   return "", notFound{"application/ld+json"}
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
