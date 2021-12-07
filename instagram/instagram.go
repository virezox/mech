package instagram

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech"
   "io"
   "net/http"
   "net/url"
)

const (
   OriginI = "https://i.instagram.com"
   OriginWWW = "https://www.instagram.com"
   queryHash = "2efa04f61586458cef44441f474eee7c"
   // com.instagram.android
   userAgent = "Instagram 214.1.0.29.120 Android"
)

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
   http.Header
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
   req, err := http.NewRequest("POST", OriginI + "/api/v1/accounts/login/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {userAgent},
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return &Login{res.Header}, nil
}

// This can be used to decode an existing login response.
func (l *Login) Decode(r io.Reader) error {
   return json.NewDecoder(r).Decode(&l.Header)
}

func (l Login) Encode(w io.Writer) error {
   enc := json.NewEncoder(w)
   enc.SetIndent("", " ")
   return enc.Encode(l.Header)
}

func (l Login) Item(shortcode string) (*Item, error) {
   req, err := http.NewRequest("GET", OriginI + "/api/v1/clips/item/", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {l.Get("Ig-Set-Authorization")},
      "User-Agent": {userAgent},
   }
   req.URL.RawQuery = "clips_media_shortcode=" + url.QueryEscape(shortcode)
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   item := new(Item)
   if err := json.NewDecoder(res.Body).Decode(item); err != nil {
      return nil, err
   }
   return item, nil
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

// If `auth` is `nil`, then anonymous request will be used.
func GraphQL(shortcode string, auth *Login) (*Media, error) {
   req, err := http.NewRequest(
      "GET", OriginWWW + "/p/" + shortcode + "/?__a=1", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", userAgent)
   if auth != nil && auth.Header != nil {
      req.Header.Set("Authorization", auth.Get("Ig-Set-Authorization"))
   }
   mech.Dump(req)
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

func (m Media) Edges() []Edge {
   if m.Shortcode_Media.Edge_Sidecar_To_Children == nil {
      return nil
   }
   return m.Shortcode_Media.Edge_Sidecar_To_Children.Edges
}

type Query struct {
   Query_Hash string `json:"query_hash"`
   Variables struct {
      Shortcode string `json:"shortcode"`
   } `json:"variables"`
}

func NewQuery(shortcode string) Query {
   var val Query
   val.Query_Hash = queryHash
   val.Variables.Shortcode = shortcode
   return val
}

// If `auth` is `nil`, then anonymous request will be used.
func (q Query) Data(auth *Login) (*Media, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(q)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", OriginI + "/graphql/query/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "User-Agent": {userAgent},
   }
   if auth != nil && auth.Header != nil {
      req.Header.Set("Authorization", auth.Get("Ig-Set-Authorization"))
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var car struct {
      Data Media
   }
   if err := json.NewDecoder(res.Body).Decode(&car); err != nil {
      return nil, err
   }
   return &car.Data, nil
}
