package deezer

import (
   "bytes"
   "encoding/json"
   "net/http"
)

const gatewayWWW = "http://www.deezer.com/ajax/gw-light.php"

type song struct {
   Results struct {
      MD5_Origin string
      Track_Token string
   }
}

func newSong(apiToken, sid string, sngId int) (song, error) {
   in, out := map[string]int{"SNG_ID": sngId}, new(bytes.Buffer)
   json.NewEncoder(out).Encode(in)
   req, err := http.NewRequest("POST", gatewayWWW, out)
   if err != nil {
      return song{}, err
   }
   val := req.URL.Query()
   val.Set("api_version", "1.0")
   val.Set("input", "3")
   val.Set("method", "song.getData")
   val.Set("api_token", apiToken)
   req.URL.RawQuery = val.Encode()
   req.AddCookie(&http.Cookie{Name: "sid", Value: sid})
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return song{}, err
   }
   var data song
   json.NewDecoder(res.Body).Decode(&data)
   return data, nil
}
