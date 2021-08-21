package vimeo

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/html"
   "net/http"
   stdhtml "html"
)

type Config struct {
   Request struct {
      Files struct {
         Progressive []struct {
            URL string
         }
      }
   }
   Video struct {
      Title string
   }
}

func NewConfig(id string) (*Config, error) {
   addr := "https://player.vimeo.com/video/" + id + "/config"
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   cfg := new(Config)
   if err := json.NewDecoder(res.Body).Decode(cfg); err != nil {
      return nil, err
   }
   return cfg, nil
}

type Video struct {
   Name string
   UploadDate string
}

func NewVideo(id string) (*Video, error) {
   addr := "https://vimeo.com/" + id
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   lex := html.NewLexer(res.Body)
   lex.NextAttr("type", "application/ld+json")
   src := lex.Bytes()
   var dst []Video
   if err := json.Unmarshal(src, &dst); err != nil {
      return nil, err
   }
   return &dst[0], nil
}

func (v Video) Title() string {
   return stdhtml.UnescapeString(v.Name)
}
