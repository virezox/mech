package spotify

import (
   "encoding/json"
   "github.com/89z/parse/html"
   "net/http"
)

type config struct {
   AccessToken string
}

func newConfig() (*config, error) {
   req, err := http.NewRequest("GET", "https://open.spotify.com", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "Firefox/60")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   lex := html.NewLexer(res.Body)
   lex.NextAttr("id", "config")
   cfg := new(config)
   if err := json.Unmarshal(lex.Bytes(), cfg); err != nil {
      return nil, err
   }
   return cfg, nil
}
