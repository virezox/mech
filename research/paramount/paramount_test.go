package paramount

import (
   "fmt"
   "testing"
   "time"
)


var issues = map[string][]string{
   "github.com/ytdl-org/youtube-dl/pull/30085": {
      "paramountplus.com/shows/the-harper-house/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/the-harper-house-the-harper-house",
   },
   "github.com/ytdl-org/youtube-dl/issues/30491": {
      "paramountplus.com/shows/star-trek-prodigy/video/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy/star-trek-prodigy-dreamcatcher",
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
   "github.com/ytdl-org/youtube-dl/issues/17870": {
      "paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
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
         fmt.Print(med, "\n\n")
         time.Sleep(99 * time.Millisecond)
      }
   }
}
