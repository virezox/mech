package youtube

import (
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "regexp"
)

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
   yt := newYouTube()
   yt.Path = "iframe_api"
   println("Get", yt.String())
   res, err := http.Get(yt.String())
   if err != nil { return "", err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return "", err }
   match := regexp.MustCompile(`/player\\/(\w+)`).FindSubmatch(body)
   id := string(match[1])
   // cache
   cache, err := os.UserCacheDir()
   if err != nil { return "", err }
   cache += "/youtube"
   play := filepath.Join(cache, id + ".js")
   _, err = os.Stat(play)
   if os.IsNotExist(err) {
      os.Mkdir(cache, os.ModeDir)
      yt.Path = fmt.Sprintf("s/player/%v/player_ias.vflset/en_US/base.js", id)
      res, err := http.Get(yt.String())
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
