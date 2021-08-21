package vimeo

import (
   "encoding/json"
   "github.com/89z/mech/html"
   "github.com/89z/mech/js"
   "net/http"
   "net/http/httputil"
   "os"
)

type Config struct {
   Clip struct {
      ID json.Number
      Uploaded_On string
   }
   JWT string
   Owner struct {
      Display_Name string
   }
   Title string
}

func NewConfig(clipID string) (*Config, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/" + clipID, nil)
   if err != nil {
      return nil, err
   }
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   lex := html.NewLexer(res.Body)
   lex.NextAttr("class", "app_banner_container")
   lex.NextTag("script")
   val := js.NewLexer(lex.Bytes()).Values()
   src := val["window.vimeo.clip_page_config"]
   dst := new(Config)
   if err := json.Unmarshal(src, dst); err != nil {
      return nil, err
   }
   return dst, nil
}

func (c Config) Video() (*Video, error) {
   req, err := http.NewRequest(
      "GET", "https://api.vimeo.com/videos/" + c.Clip.ID.String(), nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "jwt " + c.JWT)
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

type Video struct {
   Files []struct {
      Link string
   }
   Name string
   Release_Time string
   User struct {
      Name string
   }
}
