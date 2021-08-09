package main

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   req, err := http.NewRequest(
      "HEAD",
      "https://api.spotify.com/v1/playlists/6rZ28nCpmG5Wo1Ik64EoDm",
      nil,
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("Authorization", "Bearer BQDLN9hyIJn0zvspT1OOHdJBQh7pJVrA1CMBMsuu3jOFDT6o37ti_pH8__M84opDUd0yqkzKxLIfjv9XmAw")
   q := req.URL.Query()
   q.Set("additional_types", "track,episode")
   q.Set("fields", "collaborative,description,followers(total),images,name,owner(display_name,id,images,uri),public,tracks(items(track.type,track.duration_ms),total),uri")
   q.Set("market", "US")
   req.URL.RawQuery = q.Encode()
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", res)
}
