package youtube

import (
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "os"
   "regexp"
)

var API = url.URL{Scheme: "https", Host: "www.youtube.com"}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   val := make(url.Values)
   val.Set("eurl", API.String())
   val.Set("video_id", id)
   API.Path = "get_video_info"
   API.RawQuery = val.Encode()
   body, err := httpGet(API)
   if err != nil {
      return Video{}, err
   }
   val, err = url.ParseQuery(string(body))
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

// GetStream returns the url for a specific format
func (v Video) GetStream(itag int) (string, error) {
   if len(v.StreamingData.AdaptiveFormats) == 0 {
      return "", errors.New("AdaptiveFormats empty")
   }
   // get cipher text
   cipher, err := v.signatureCipher(itag)
   if err != nil { return "", err }
   query, err := url.ParseQuery(cipher)
   if err != nil { return "", err }
   sig := []byte(query.Get("s"))
   // get player
   API.Path = "iframe_api"
   body, err := httpGet(API)
   if err != nil { return "", err }
   id := regexp.MustCompile(`/player\\/(\w+)`).FindSubmatch(body)[1]
   cache, err := os.UserCacheDir()
   if err != nil { return "", err }
   cache += "/youtube"
   play := fmt.Sprintf("%v/%s.js", cache, id)
   _, err = os.Stat(play)
   if os.IsNotExist(err) {
      os.Mkdir(cache, os.ModeDir)
      API.Path = fmt.Sprintf("s/player/%s/player_ias.vflset/en_US/base.js", id)
      res, err := http.Get(API.String())
      if err != nil { return "", err }
      defer res.Body.Close()
      file, err := os.Create(play)
      if err != nil { return "", err }
      defer file.Close()
      file.ReadFrom(res.Body)
   } else if err != nil {
      return "", err
   } else {
      println("Exist", play)
   }
   body, err = os.ReadFile(play)
   if err != nil { return "", err }
   err = decrypt(sig, body)
   if err != nil { return "", err }
   return query.Get("url") + "&sig=" + string(sig), nil
}
