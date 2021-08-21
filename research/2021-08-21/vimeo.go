package vimeo

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/html"
   "github.com/89z/mech/js"
   "net/http"
   "os"
)

type config struct {
   Clip struct {
      ID json.Number
   }
   JWT string
}

func (c config) videos() error {
   req, err := http.NewRequest(
      "GET", "https://api.vimeo.com/videos/" + c.Clip.ID.String(), nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("Authorization", "jwt " + c.JWT)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   f, err := os.Create("videos.json")
   if err != nil {
      return err
   }
   defer f.Close()
   f.ReadFrom(res.Body)
   return nil
}

func newConfig(clipID string) (*config, error) {
   addr := "https://vimeo.com/" + clipID
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   lex := html.NewLexer(res.Body)
   lex.NextAttr("class", "app_banner_container")
   lex.NextTag("script")
   val := js.NewLexer(lex.Bytes()).Values()
   src := val["window.vimeo.clip_page_config"]
   dst := new(config)
   if err := json.Unmarshal(src, dst); err != nil {
      return nil, err
   }
   return dst, nil
}
