package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "regexp"
)

type youTube struct {
   url.URL
   url.Values
}

func newYouTube() youTube {
   return youTube{
      url.URL{Scheme: "https", Host: "www.youtube.com"}, make(url.Values),
   }
}

func getPlayer() ([]byte, error) {
   yt := newYouTube()
   yt.Path = "iframe_api"
   println("Get", yt.String())
   res, err := http.Get(yt.String())
   if err != nil { return nil, err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return nil, err }
   match := regexp.MustCompile(`/player\\/(\w+)`).FindSubmatch(body)
   id := string(match[1])
   // cache
   cache, err := os.UserCacheDir()
   if err != nil { return nil, err }
   cache = filepath.Join(cache, "youtube")
   play := filepath.Join(cache, id + ".js")
   _, err = os.Stat(play)
   if os.IsNotExist(err) {
      os.Mkdir(cache, os.ModeDir)
      yt.Path = fmt.Sprintf("s/player/%v/player_ias.vflset/en_US/base.js", id)
      res, err := http.Get(yt.String())
      if err != nil { return nil, err }
      defer res.Body.Close()
      file, err := os.Create(play)
      if err != nil { return nil, err }
      defer file.Close()
      file.ReadFrom(res.Body)
   } else if err != nil {
      return nil, err
   } else {
      println("Exist", play)
   }
   return os.ReadFile(play)
}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   yt := newYouTube()
   yt.Set("eurl", yt.String())
   yt.Set("video_id", id)
   yt.RawQuery = yt.Encode()
   yt.Path = "get_video_info"
   println("Get", yt.String())
   res, err := http.Get(yt.String())
   if err != nil {
      return Video{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Video{}, err
   }
   val, err := url.ParseQuery(string(body))
   if err != nil {
      return Video{}, err
   }
   var (
      play = val.Get("player_response")
      vid Video
   )
   err = json.Unmarshal([]byte(play), &vid)
   if err != nil {
      return Video{}, err
   }
   return vid, nil
}
