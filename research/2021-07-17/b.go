package main

import (
   "io"
   "net/http"
   "os"
)

func main() {
   req, err := http.NewRequest(
      "GET", "https://www.youtube.com/get_video_info", nil,
   )
   if err != nil {
      panic(err)
   }
   q := req.URL.Query()
   q.Set("c", "TVHTML5")
   q.Set("cver", "7.20210428.10.00")
   q.Set("el", "detailpage")
   q.Set("html5", "1")
   q.Set("video_id", "Cr381pDsSsA")
   q.Set("access_token", "ya29.a0ARrdaM9yZ-mM-s1CCR0uP0mPUY3h4WiUEbc-tW2FkX...")
   req.URL.RawQuery = q.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   req.URL.RawQuery = string(body)
   play := req.URL.Query().Get("player_response")
   f, err := os.Create("yt.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   f.WriteString(play)
}
