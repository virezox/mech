package apple

import (
   "encoding/json"
   "net/http"
)

func (p Poster) Body(buf []byte) ([]byte, error) {
   var s struct {
      Challenge []byte `json:"challenge"`
      Server_Parameters Server_Parameters `json:"extra-server-parameters"`
      Key_System string `json:"key-system"`
      URI string `json:"uri"`
   }
   s.Challenge = buf
   s.Key_System = "com.widevine.alpha"
   s.Server_Parameters = p.episode.Asset().FpsKeyServerQueryParameters
   s.URI = p.pssh
   return json.Marshal(s)
}

func (p Poster) Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + p.env.Media_API.Token)
   head.Set("Content-Type", "application/json")
   head.Set("X-Apple-Music-User-Token", p.auth.media_user_token().Value)
   return head
}

type Server_Parameters struct {
   Adam_ID string `json:"adamId"`
   Svc_ID string `json:"svcId"`
}

func (p Poster) License_URL() string {
   return p.episode.Asset().FpsKeyServerUrl
}

type Poster struct {
   auth Auth // Header
   env *Environment // Header
   episode *Episode // URL, Body
   pssh string
}
