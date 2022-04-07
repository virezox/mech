package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "testing"
   "time"
)

// 35
var alfas = []string{
   /* alfa */ "1t24XAntNCY",
   /* alfa */ "2NUZ8W2llS4",
   /* alfa */ "3b8nCWDgZ6Q",
   /* alfa */ "5KLPxDtMqe8",
   /* alfa */ "9lWxNJF-ufM",
   /* alfa */ "BGQWPY4IigY",
   /* alfa */ "BaW_jenozKc",
   /* alfa */ "CHqg6qOn4no",
   /* alfa */ "FIl7x6_3R5Y",
   /* alfa */ "FRhJzUSJbGI",
   /* alfa */ "FlRa-iH7PGw",
   /* alfa */ "IB3lcPjvWLA",
   /* alfa */ "M4gD1WSo5mA",
   /* alfa */ "MeJVWBSsPAY",
   /* alfa */ "MgNrAu2pzNs",
   /* alfa */ "OtqTfy26tG0",
   /* alfa */ "XclachpHxis",
   /* alfa */ "YOelRv7fMxY",
   /* alfa */ "Yh0AhrY9GjA",
   /* alfa */ "Z4Vy8R84T1U",
   /* alfa */ "__2ABJjxzNo",
   /* alfa */ "_b-2C3KPAM0",
   /* alfa */ "a9LDPn-MO4I",
   /* alfa */ "cBvYw8_A0vQ",
   /* alfa */ "eQcmzGIKrzg",
   /* alfa */ "gVfLd0zydlo",
   /* alfa */ "gVfgbahppCY",
   /* alfa */ "iqKdEhx-dD4",
   /* alfa */ "jvGDaLqkpTg",
   /* alfa */ "kgx4WGK0oNU",
   /* alfa */ "lqQg6PlCWgI",
   /* alfa */ "lsguqyKfVQg",
   /* alfa */ "mzZzzBU6lrM",
   /* alfa */ "wsQiKKfKxug",
   /* alfa */ "x41yOUIvK2k",
}

// 2
var bravos = []string{
   /* bravo */ "HtVdAasjOgU",
   /* bravo */ "WaOKSUlf4TM",
}

// 5
var charlies = []string{
   "SZJvDhaSDnc", // pass
   "Tq92D6wQ1mg", // pass
   "i1Ko8UG-Tdo", // pass
   "nGC3D_FkCmg", // pass
   "yYr8q0y5Jfg", // pass
}

// 1
var deltas = []string{
   "HsUATh_Nc2U", // TVHTML5_SIMPLY_EMBEDDED_PLAYER
}

// 12
var echos = []string{
   "63RmMXCd_bQ", // This live stream recording is not available
   "6SJNVb0GnPI", // This video has been removed for violating
   "Cr381pDsSsA", // Sign in to confirm your age
   "CsmdDsKjzN8", // This live stream recording is not available
   "DJztXj2GPfl", // Video unavailable
   "Ms7iBXnlUO8", // Video unavailable
   "Q39EVAstoRM", // Video unavailable
   "V36LpHqtcDY", // Private video
   "qEJwOuvDf7I", // This live stream recording is not available
   "s7_qI6_mIXc", // This video is DRM protected
   "sJL6WA-aGkQ", // Video unavailable
   "yZIXLfi8CZQ", // Private video
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
