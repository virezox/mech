package youtube_test

import (
   "github.com/89z/mech/youtube"
   "io"
   "testing"
   "time"
)

var ids = []string{
   // pass
   "9bZkp7q19f0",//Normal
   // todo
   "5qap5aO4i9A",//LiveStream
   "rsAAeyAr-9Y",//LiveStreamRecording
   "V5Fsj_sCKdg",//ContainsHighQualityStreams
   "AI7ULzgf8RU",//ContainsDashManifest
   "-xNN-bJQ4vI",//Omnidirectional
   "vX2vsvdq8nw",//HighDynamicRange
   "YltHGKX80Y8",//ContainsClosedCaptions
   "_kmeFXjjGfk",//EmbedRestrictedByYouTube
   "MeJVWBSsPAY",//EmbedRestrictedByAuthor
   "SkRSXFQerZs",//AgeRestricted
   "hySoCSoH-g8",//AgeRestrictedEmbedRestricted
   "5VGm0dczmHc",//RatingDisabled
   "p3dDcKOFXQg",//RequiresPurchase
   // fail
   "ZGdLIwrGHG8",//Unlisted
}

func TestAndroid(t *testing.T) {
   for _, id := range ids {
      a, err := youtube.NewAndroid(id)
      if err != nil {
         t.Fatal(err)
      }
      a.AdaptiveFormats.Sort(func(a, b youtube.Format) bool {
         return a.ContentLength < b.ContentLength
      })
      if err := a.AdaptiveFormats[0].Write(io.Discard); err != nil {
         t.Fatal(err)
      }
      time.Sleep(100 * time.Millisecond)
   }
}
