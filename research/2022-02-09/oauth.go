package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   val := url.Values{
      "grant_type":[]string{"client_credentials"},
      "scope":[]string{"private public create edit delete interact upload purchased stats"},
   }
   req, err := http.NewRequest(
      "POST", "https://api.vimeo.com/oauth/authorize/client",
      strings.NewReader(val.Encode()),
   )
   if err != nil {
      panic(err)
   }
   req.SetBasicAuth(
      "74fa89b811a1cbb750d8528d163f48af28a2dbe1",
      "SHI3xbktBLL3+zxyu4J3Se/klYgQmrwHer7W/II2QgdvBh0letIhxxDdlWvzLbBRl+e+oE5SYTjTf7C3YV5lIoCrknxeex01dSsSVooIneJjLiwKDuc+aLA5MKEkbahZ",
   )
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
