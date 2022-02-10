package vimeo

import (
   "encoding/json"
   "net/http"
   "strings"
)

func (b bearer) video(path string) (*video, error) {
   return newVideo(path, "Bearer " + b.Access_Token)
}

type bearer struct {
   Access_Token string
}

func newBearer() (*bearer, error) {
   req, err := http.NewRequest(
      "POST", "https://api.vimeo.com/oauth/authorize/client",
      strings.NewReader("grant_type=client_credentials"),
   )
   if err != nil {
      return nil, err
   }
   var str strings.Builder
   str.WriteString("SHI3xbktBLL3+zxyu4J3Se/klYgQmrwHer7W/")
   str.WriteString("II2QgdvBh0letIhxxDdlWvzLbBRl+e+")
   str.WriteString("oE5SYTjTf7C3YV5lIoCrknxeex01dSsSVooIneJjLiwKDuc+")
   str.WriteString("aLA5MKEkbahZ")
   req.SetBasicAuth("74fa89b811a1cbb750d8528d163f48af28a2dbe1", str.String())
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   bear := new(bearer)
   if err := json.NewDecoder(res.Body).Decode(bear); err != nil {
      return nil, err
   }
   return bear, nil
}
