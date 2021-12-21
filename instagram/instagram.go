package instagram

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "io"
   "net/http"
)

const (
   originI = "https://i.instagram.com"
   queryHash = "2efa04f61586458cef44441f474eee7c"
   // com.instagram.android
   userAgent = "Instagram 214.1.0.29.120 Android"
)

var LogLevel mech.LogLevel

// instagram.com/p/CT-cnxGhvvO
// instagram.com/p/yza2PAPSx2
func Valid(shortcode string) bool {
   switch len(shortcode) {
   case 10, 11:
      return true
   }
   return false
}

type Edge struct {
   Node struct {
      Display_URL string
      Video_URL string
   }
}

func (e Edge) URL() string {
   if e.Node.Video_URL != "" {
      return e.Node.Video_URL
   }
   return e.Node.Display_URL
}

type Item struct {
   Media struct {
      Video_Versions []struct {
         URL string
      }
   }
}

type Login struct {
   Authorization string
}

func NewLogin(username, password string) (*Login, error) {
   buf := bytes.NewBufferString("signed_body=SIGNATURE.")
   sig := map[string]string{
      "device_id": userAgent,
      "enc_password": "#PWD_INSTAGRAM:0:0:" + password,
      "username": username,
   }
   if err := json.NewEncoder(buf).Encode(sig); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", originI + "/api/v1/accounts/login/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {userAgent},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var login Login
   login.Authorization = res.Header.Get("Ig-Set-Authorization")
   return &login, nil
}

// This can be used to decode an existing login response.
func (l *Login) Decode(src io.Reader) error {
   return json.NewDecoder(src).Decode(l)
}

func (l Login) Encode(dst io.Writer) error {
   enc := json.NewEncoder(dst)
   enc.SetIndent("", " ")
   return enc.Encode(l)
}

// If `Authorization` field is empty, then anonymous request will be used.
func (l Login) GraphQL(shortcode string) (*Media, error) {
   req, err := http.NewRequest(
      "GET", "https://www.instagram.com/p/" + shortcode + "/", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", userAgent)
   if l.Authorization != "" {
      req.Header.Set("Authorization", l.Authorization)
   }
   req.URL.RawQuery = "__a=1"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var car struct {
      GraphQL Media
   }
   if err := json.NewDecoder(res.Body).Decode(&car); err != nil {
      return nil, err
   }
   return &car.GraphQL, nil
}

type Media struct {
   Shortcode_Media struct {
      Display_URL string
      Edge_Media_Preview_Like struct {
         Count int
      }
      Edge_Media_To_Parent_Comment struct {
         Edges []struct {
            Node struct {
               Text string
            }
         }
      }
      Edge_Sidecar_To_Children *struct {
         Edges []Edge
      }
      Video_URL string
   }
}

func (m Media) Edges() []Edge {
   if m.Shortcode_Media.Edge_Sidecar_To_Children == nil {
      return nil
   }
   return m.Shortcode_Media.Edge_Sidecar_To_Children.Edges
}
