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

// 35
var alfas = []string{
   "1t24XAntNCY",
   "2NUZ8W2llS4",
   "3b8nCWDgZ6Q",
   "5KLPxDtMqe8",
   "9lWxNJF-ufM",
   "BGQWPY4IigY",
   "BaW_jenozKc",
   "CHqg6qOn4no",
   "FIl7x6_3R5Y",
   "FRhJzUSJbGI",
   "FlRa-iH7PGw",
   "IB3lcPjvWLA",
   "M4gD1WSo5mA",
   "MeJVWBSsPAY",
   "MgNrAu2pzNs",
   "OtqTfy26tG0",
   "XclachpHxis",
   "YOelRv7fMxY",
   "Yh0AhrY9GjA",
   "Z4Vy8R84T1U",
   "__2ABJjxzNo",
   "_b-2C3KPAM0",
   "a9LDPn-MO4I",
   "cBvYw8_A0vQ",
   "eQcmzGIKrzg",
   "gVfLd0zydlo",
   "gVfgbahppCY",
   "iqKdEhx-dD4",
   "jvGDaLqkpTg",
   "kgx4WGK0oNU",
   "lqQg6PlCWgI",
   "lsguqyKfVQg",
   "mzZzzBU6lrM",
   "wsQiKKfKxug",
   "x41yOUIvK2k",
}

// 2
var bravos = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

// 4
var charlies = []string{
   "Cr381pDsSsA", // Sign in to confirm your age
   "HsUATh_Nc2U", // TVHTML5_SIMPLY_EMBEDDED_PLAYER
   "SZJvDhaSDnc", // TVHTML5_SIMPLY_EMBEDDED_PLAYER
   "Tq92D6wQ1mg", // TVHTML5_SIMPLY_EMBEDDED_PLAYER
}

// 1
var deltas = []string{
   "nGC3D_FkCmg", // TVHTML5_SIMPLY_EMBEDDED_PLAYER
}

// 13
var echos = []string{
   "63RmMXCd_bQ", // This live stream recording is not available
   "6SJNVb0GnPI", // This video has been removed for violating
   "CsmdDsKjzN8", // This live stream recording is not available
   "DJztXj2GPfl", // Video unavailable
   "Ms7iBXnlUO8", // Video unavailable
   "Q39EVAstoRM", // Video unavailable
   "V36LpHqtcDY", // Private video
   "i1Ko8UG-Tdo", // YouTube Premium
   "qEJwOuvDf7I", // This live stream recording is not available
   "s7_qI6_mIXc", // This video is DRM protected
   "sJL6WA-aGkQ", // Video unavailable
   "yYr8q0y5Jfg", // YouTube Movies
   "yZIXLfi8CZQ", // Private video
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
      var play player
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
      var play player
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

type player struct {
   PlayabilityStatus struct {
      Status string
      Reason string
   }
}

func TestAlfa(t *testing.T) {
   for _, alfa := range alfas {
      play, err := Android.Player(alfa)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status == "OK" {
         fmt.Printf("/* alfa */ %q,\n", alfa)
      } else {
         fmt.Printf("/* bravo */ %q,\n", alfa)
      }
      time.Sleep(time.Second)
   }
}

func TestSearch(t *testing.T) {
   search, err := Mweb.Search("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range search.Items() {
      fmt.Println(item.CompactVideoRenderer)
   }
}
