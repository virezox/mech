package main

import (
   "fmt"
   "io"
   "net/http"
   "strings"
   "time"
)

const origin = "https://www.youtube.com"

var ids = []string{
   "-xNN-bJQ4vI", //Omnidirectional
   "5VGm0dczmHc", //RatingDisabled
   "5qap5aO4i9A", //LiveStream
   "9bZkp7q19f0", //Normal
   "AI7ULzgf8RU", //ContainsDashManifest
   "MeJVWBSsPAY", //EmbedRestrictedByAuthor
   "SkRSXFQerZs", //AgeRestricted
   "V5Fsj_sCKdg", //ContainsHighQualityStreams
   "YltHGKX80Y8", //ContainsClosedCaptions
   "ZGdLIwrGHG8", //Unlisted
   "_kmeFXjjGfk", //EmbedRestrictedByYouTube
   "hySoCSoH-g8", //AgeRestrictedEmbedRestricted
   "p3dDcKOFXQg", //RequiresPurchase
   "rsAAeyAr-9Y", //LiveStreamRecording
   "vX2vsvdq8nw", //HighDynamicRange
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
   for _, id := range ids {
      q.Set("video_id", id)
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
      if strings.Contains(play, `"adaptiveFormats"`) {
         fmt.Println("pass", id)
      } else {
         fmt.Println("fail", id)
      }
      time.Sleep(100 * time.Millisecond)
   }
}
