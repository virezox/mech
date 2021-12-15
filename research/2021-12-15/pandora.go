package main

import (
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

const body = `
{"stationId":"126608766085892525","isStationStart":true,
"fragmentRequestReason":"Normal","audioFormat":"aacplus",
"startingAtTrackId":null,"onDemandArtistMessageArtistUidHex":null,
"onDemandArtistMessageIdHex":null}
`

func main() {
   req, err := http.NewRequest(
      "POST", "https://pandora.com/api/v1/playlist/getFragment",
      strings.NewReader(body),
   )
   if err != nil {
      panic(err)
   }
   req.Header = http.Header{
      "Content-Type":[]string{"application/json"},
      "Cookie":[]string{"csrftoken=842b12c83a3c5153"},
      "X-Authtoken":[]string{"BXoTKywEhnoiEqDEcu0U/qGlFBEK5Tjblz3fgnLPgFojficRTR8Xm6Lw=="},
      "X-Csrftoken":[]string{"842b12c83a3c5153"},
   }
   buf, err := httputil.DumpRequestOut(req, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err = httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
