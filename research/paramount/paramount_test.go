package paramount

import (
   "fmt"
   "testing"
   "time"
)

var issues = map[string][]string{
   "github.com/ytdl-org/youtube-dl/issues/17870": {
      "paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
   },
   "github.com/ytdl-org/youtube-dl/issues/29038": {
      "paramountplus.com/movies/spongebob-movie-sponge-on-the-run/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD",
   },
   "github.com/ytdl-org/youtube-dl/issues/29089": {
      "paramountplus.com/shows/yo-soy-franky/video/GvrmB2nmQgk4posE5Se7GUOrMqxa08aO/yo-soy-franky-nace-franky",
   },
   "github.com/ytdl-org/youtube-dl/issues/29351": {
      "paramountplus.com/shows/rupauls-drag-race/video/XeLhvObaelBXGoTnchPDWHH2_W76Jwgc/rupaul-s-drag-race-drama-queens",
   },
   "github.com/ytdl-org/youtube-dl/issues/29454": {
      "paramountplus.com/shows/star-trek-discovery/video/UF5kG511ESgjBRjYvmmmJFGso_wKdLPg/star-trek-discovery-such-sweet-sorrow",
      "paramountplus.com/shows/star-trek-discovery/video/trdn_QshF1DlbsdedG6xdsnd5WwF6032/star-trek-discovery-si-vis-pacem-para-bellum",
   },
   "github.com/ytdl-org/youtube-dl/issues/29564": {
      "paramountplus.com/shows/catdog/video/Oe44g5_NrlgiZE3aQVONleD6vXc8kP0k/catdog-climb-every-catdog-the-canine-mutiny",
   },
   "github.com/ytdl-org/youtube-dl/issues/30066": {
      "cbs.com/shows/survivor/video/EktRdacCXDsJo8Qqyw18Rg26tLwfzk4N/survivor-juggling-chainsaws",
      "cbs.com/shows/survivor/video/Rcmd9CGvysrl13uSff93up_SnD5RzVfH/survivor-my-million-dollar-mistake",
      "paramountplus.com/shows/bull/video/TUT_4UVB87huHEOfPCjMkxOW_Xe1hNWw/bull-gone",
      "paramountplus.com/shows/csi-vegas/video/ct82smMrQDbogrBqISHZEJKj3IUOI1yo/csi-vegas-legac",
      "paramountplus.com/shows/fbi-most-wanted/video/xKRZ7nMYi0i9WjKk5yUBwIPmmEVRrXvM/fbi-most-wanted-tough-love",
      "paramountplus.com/shows/ncis/video/LzcJLBA2OGq_fE1bDD3NF7lx_zGEyt3L/ncis-road-to-nowhere",
      "paramountplus.com/shows/survivor/video/Rcmd9CGvysrl13uSff93up_SnD5RzVfH/survivor-my-million-dollar-mistake",
   },
   "github.com/ytdl-org/youtube-dl/issues/30247": {
      "paramountplus.com/shows/star-trek-prodigy/video/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy/star-trek-prodigy-dreamcatcher",
   },
}

func TestParamount(t *testing.T) {
   LogLevel = -1
   for issue, videos := range issues {
      fmt.Print("## ", issue, "\n\n")
      for _, video := range videos {
         id := VideoID(video)
         med, err := Media(id)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(video)
         fmt.Print(med.Host, "\n\n")
         time.Sleep(99 * time.Millisecond)
      }
   }
}
