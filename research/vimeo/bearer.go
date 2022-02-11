package vimeo

import (
   "encoding/json"
   "net/http"
   "strings"
)

type Bearer struct {
   Access_Token string
}

// I am not sure if this is any better than JWT, but I didnt want to just
// delete it. So leaving it here for now.
func NewBearer() (*Bearer, error) {
   req, err := http.NewRequest(
      "POST", "https://api.vimeo.com/oauth/authorize/client",
      strings.NewReader("grant_type=client_credentials"),
   )
   if err != nil {
      return nil, err
   }
   var buf strings.Builder
   buf.WriteString("SHI3xbktBLL3+zxyu4J3Se/klYgQmrwHer7W/")
   buf.WriteString("II2QgdvBh0letIhxxDdlWvzLbBRl+e+")
   buf.WriteString("oE5SYTjTf7C3YV5lIoCrknxeex01dSsSVooIneJjLiwKDuc+")
   buf.WriteString("aLA5MKEkbahZ")
   req.SetBasicAuth("74fa89b811a1cbb750d8528d163f48af28a2dbe1", buf.String())
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   bear := new(Bearer)
   if err := json.NewDecoder(res.Body).Decode(bear); err != nil {
      return nil, err
   }
   return bear, nil
}

func (b Bearer) Video(path string) (*Video, error) {
   return newVideo(path, "Bearer " + b.Access_Token)
}
