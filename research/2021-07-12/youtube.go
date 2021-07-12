package main

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "time"
)


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
