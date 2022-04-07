package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
   "testing"
   "time"
)

func TestSearch(t *testing.T) {
   search, err := Mweb.Search("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range search.Items() {
      fmt.Println(item.CompactVideoRenderer)
   }
}

// ANDROID
const alfa = "zv9NimPx3Es"

func TestAlfa(t *testing.T) {
   play, err := Android.Player(alfa)
   if err != nil {
      t.Fatal(err)
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal(play)
   }
}

// ANDROID_EMBEDDED_PLAYER
var bravos = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

func TestBravo(t *testing.T) {
   const (
      name = "ANDROID_EMBEDDED_PLAYER"
      version = "17.11.34"
   )
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   val := make(url.Values)
   val["key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   for _, bravo := range bravos {
      req.Body = io.NopCloser(strings.NewReader(fmt.Sprintf(`
      {
        "context": {
          "client": {
            "clientName": %q,
            "clientVersion": %q,
          }
        },
        "videoId": %q
      }
      `, name, version, bravo)))
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         t.Fatal(err)
      }
      var play Player
      if err := json.NewDecoder(res.Body).Decode(&play); err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status == "OK" {
         fmt.Printf("/* bravo */ %q,\n", bravo)
      } else {
         fmt.Printf("/* charlie */ %q,\n", bravo)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

// racyCheckOk
var charlies = []string{
   "Cr381pDsSsA",
   "HsUATh_Nc2U",
   "SZJvDhaSDnc",
   "Tq92D6wQ1mg",
}

func TestCharlie(t *testing.T) {
   const (
      name = "ANDROID"
      version = "17.11.34"
   )
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   change, err := OpenExchange(cache, "mech/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, charlie := range charlies {
      body := strings.NewReader(fmt.Sprintf(`
      {
        "context": {
          "client": {
            "clientName": %q,
            "clientVersion": %q,
          }
        },
        "videoId": %q,
      "racyCheckOk": true
      }
      `, name, version, charlie))
      req, err := http.NewRequest(
         "POST", "https://www.youtube.com/youtubei/v1/player", body,
      )
      if err != nil {
         t.Fatal(err)
      }
      req.Header.Set("Authorization", "Bearer " + change.Access_Token)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         t.Fatal(err)
      }
      var play Player
      if err := json.NewDecoder(res.Body).Decode(&play); err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status == "OK" {
         fmt.Printf("/* charlie */ %q,\n", charlie)
      } else {
         fmt.Printf("/* delta */ %q,\n", charlie)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

// contentCheckOk
const delta = "nGC3D_FkCmg"

func TestDelta(t *testing.T) {
   const (
      name = "ANDROID"
      version = "17.11.34"
   )
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   change, err := OpenExchange(cache, "mech/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   body := strings.NewReader(fmt.Sprintf(`
   {
     "context": {
       "client": {
         "clientName": %q,
         "clientVersion": %q,
       }
     },
     "videoId": %q,
   "racyCheckOk": true,
   "contentCheckOk": true
   }
   `, name, version, delta))
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", body,
   )
   if err != nil {
      t.Fatal(err)
   }
   req.Header.Set("Authorization", "Bearer " + change.Access_Token)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   var play Player
   if err := json.NewDecoder(res.Body).Decode(&play); err != nil {
      t.Fatal(err)
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal(play)
   }
}
