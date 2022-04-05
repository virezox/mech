package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "os"
   "testing"
   gp "github.com/89z/googleplay"
)

func TestAndroid(t *testing.T) {
   var (
      err error
      play player
   )
   play.VideoID = "eZHsmb4ezEk"
   play.Context.Client.ClientName = "ANDROID"
   play.Context.Client.ClientVersion, err = appVersion(
      "com.google.android.youtube",
   )
   if err != nil {
      t.Fatal(err)
   }
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(play)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      t.Fatal(err)
   }
   req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, play.Context.Client)
}
