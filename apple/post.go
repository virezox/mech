package apple

import (
   "encoding/json"
   "net/http"
)

type Poster struct {
   auth Auth // Header
   env *Environment // Header
   episode *Episode // URL, Body
   pssh string
}

func (p Poster) Request_URL() string {
   return p.episode.Asset().FpsKeyServerUrl
}

func (p Poster) Request_Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + p.env.Media_API.Token)
   head.Set("Content-Type", "application/json")
   head.Set("X-Apple-Music-User-Token", p.auth.media_user_token().Value)
   return head
}

func (p Poster) Request_Body(buf []byte) ([]byte, error) {
   var s struct {
      Challenge []byte `json:"challenge"`
      Key_System string `json:"key-system"`
      Server_Parameters Server_Parameters `json:"extra-server-parameters"`
      URI string `json:"uri"`
   }
   s.Challenge = buf
   s.Key_System = "com.widevine.alpha"
   s.Server_Parameters = p.episode.Asset().FpsKeyServerQueryParameters
   s.URI = p.pssh
   return json.MarshalIndent(s, "", " ")
}

func (Poster) Response_Body(buf []byte) ([]byte, error) {
   var s struct {
      License []byte
   }
   err := json.Unmarshal(buf, &s)
   if err != nil {
      return nil, err
   }
   return s.License, nil
}
