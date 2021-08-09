package main

import (
   "fmt"
   "net/http"
)

func main() {
   req, err := http.NewRequest(
      "HEAD",
      "https://api.spotify.com/v1/playlists/6rZ28nCpmG5Wo1Ik64EoDm/tracks",
      nil,
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("Authorization", "Bearer BQC1D86LnOO67vOcMVg0nBzMo3Sr5J6yS7SqSm610E-uLPonxBZ2Ava5QuZ2siQpB7iSnlx25Y0kgVHVPAE")
   q := req.URL.Query()
   q.Set("additional_types", "track,episode")
   q.Set("limit", "100")
   q.Set("market", "US")
   q.Set("offset", "0")
   req.URL.RawQuery = q.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", res)
}
