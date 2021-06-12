package deezer

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type songList struct {
   Results struct {
      Data []struct {
         MD5_Origin string
         Track_Token string
      }
   }
}

func newSongList(apiToken, sid string, sngIds ...int) (songList, error) {
   in, out := map[string][]int{"SNG_IDS": sngIds}, new(bytes.Buffer)
   json.NewEncoder(out).Encode(in)
   req, err := http.NewRequest("POST", gatewayWWW, out)
   if err != nil {
      return songList{}, err
   }
   val := req.URL.Query()
   val.Set("api_version", "1.0")
   val.Set("input", "3")
   val.Set("method", "song.getListData")
   val.Set("api_token", apiToken)
   req.URL.RawQuery = val.Encode()
   req.AddCookie(&http.Cookie{Name: "sid", Value: sid})
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return songList{}, err
   }
   var list songList
   json.NewDecoder(res.Body).Decode(&list)
   return list, nil
}
