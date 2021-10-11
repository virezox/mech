package spotify

import (
   "encoding/json"
   "github.com/89z/mech"
   "github.com/89z/parse/html"
   "io"
   "net/http"
)

const Origin = "https://api.spotify.com"

var Verbose = mech.Verbose

type Album struct {
   Tracks struct {
      Items []Track
   }
}

type Config struct {
   AccessToken string
}

func NewConfig() (*Config, error) {
   req, err := http.NewRequest("GET", "https://open.spotify.com", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "Firefox/60")
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   lex := html.NewLexer(res.Body)
   lex.NextAttr("id", "config")
   cfg := new(Config)
   if err := json.Unmarshal(lex.Bytes(), cfg); err != nil {
      return nil, err
   }
   return cfg, nil
}

func (c Config) Album(id string) (*Album, error) {
   req, err := http.NewRequest("GET", Origin + "/v1/albums/" + id, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + c.AccessToken)
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   alb := new(Album)
   if err := json.NewDecoder(res.Body).Decode(alb); err != nil {
      return nil, err
   }
   return alb, nil
}

// This can be used to decode a previously saved token.
func (c *Config) Decode(r io.Reader) error {
   return json.NewDecoder(r).Decode(c)
}

func (c Config) Encode(w io.Writer) error {
   enc := json.NewEncoder(w)
   enc.SetIndent("", " ")
   return enc.Encode(c)
}

func (c Config) Playlist(id string) (*Playlist, error) {
   req, err := http.NewRequest("GET", Origin + "/v1/playlists/" + id, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + c.AccessToken)
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   list := new(Playlist)
   if err := json.NewDecoder(res.Body).Decode(list); err != nil {
      return nil, err
   }
   return list, nil
}

type Playlist struct {
   Tracks struct {
      Items []struct {
         Track Track
      }
   }
}

type Track struct {
   Name string
}
