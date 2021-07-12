package main

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "time"
)

const origin = "https://www.youtube.com"

type video struct {
   id string
   el string
}

var videos = []video{
   {id: "-xNN-bJQ4vI"}, //Omnidirectional
   {id: "54e6lBE3BoQ"}, // ANDROID
   {id: "5VGm0dczmHc"}, //RatingDisabled
   {id: "5qap5aO4i9A"}, //LiveStream
   {id: "9bZkp7q19f0"}, //Normal
   {id: "AI7ULzgf8RU"}, //ContainsDashManifest
   {id: "HtVdAasjOgU"}, // WEB_EMBEDDED_PLAYER
   {id: "MeJVWBSsPAY", el: "detailpage"}, //EmbedRestrictedByAuthor
   {id: "SkRSXFQerZs", el: ""}, //AgeRestricted
   {id: "V5Fsj_sCKdg"}, //ContainsHighQualityStreams
   {id: "XeojXq6ySs4"}, // Topic
   {id: "YltHGKX80Y8"}, //ContainsClosedCaptions
   {id: "ZGdLIwrGHG8"}, //Unlisted
   {id: "_kmeFXjjGfk"}, //EmbedRestrictedByYouTube
   {id: "hySoCSoH-g8"}, //AgeRestrictedEmbedRestricted
   {id: "rsAAeyAr-9Y"}, //LiveStreamRecording
   {id: "vX2vsvdq8nw", el: "detailpage"}, //HighDynamicRange
}

func main() {
   req, err := http.NewRequest("GET", origin + "/get_video_info", nil)
   if err != nil {
      panic(err)
   }
   q := req.URL.Query()
   q.Set("c", "ANDROID")
   q.Set("cver", "16.05")
   q.Set("eurl", origin)
   q.Set("html5", "1")
   for _, vid := range videos {
      q.Set("video_id", vid.id)
      if vid.el != "" {
         q.Set("el", "detailpage")
      } else {
         q.Del("el")
      }
      req.URL.RawQuery = q.Encode()
      fmt.Println(req.Method, req.URL)
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
      var p player
      json.Unmarshal([]byte(play), &p)
      if pass(p.StreamingData.AdaptiveFormats) {
         fmt.Println("pass", vid.id)
      } else {
         fmt.Println("fail", vid.id)
      }
      time.Sleep(100 * time.Millisecond)
   }
}

func pass(fs []format) bool {
   for _, f := range fs {
      if f.Itag == 134 && f.URL != "" {
         return true
      }
   }
   return false
}

type format struct {
   Itag int
   URL string
}

type player struct {
   StreamingData struct {
      AdaptiveFormats []format
   }
}
