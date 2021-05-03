package youtube

import (
   "errors"
   "fmt"
   "net/url"
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
   body, err := readAll("https://www.youtube.com/iframe_api")
   if err != nil { return "", err }
   id := regexp.MustCompile(`/player\\/(\w+)`).FindSubmatch(body)
   base := url.URL{Scheme: "https", Host: "www.youtube.com"}
   base.Path = fmt.Sprintf("/s/player/%s/player_ias.vflset/en_US/base.js", id[1])
   body, err = readAll(base.String())
   if err != nil { return "", err }
   err = decrypt(sig, body)
   if err != nil { return "", err }
   return query.Get("url") + "&sig=" + string(sig), nil
}
