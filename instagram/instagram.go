package instagram

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "io"
   "net/http"
)

const (
   OriginI = "https://i.instagram.com"
   OriginWWW = "https://www.instagram.com"
   queryHash = "2efa04f61586458cef44441f474eee7c"
   userAgent = "Instagram 207.0.0.39.120 Android"
)

var Verbose = mech.Verbose

// instagram.com/p/CT-cnxGhvvO
func Valid(shortcode string) error {
   if len(shortcode) == 11 {
      return nil
   }
   return fmt.Errorf("%q invalid as shortcode", shortcode)
}

type Edge struct {
   Node struct {
      Display_URL string
   }
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
   res, err := mech.RoundTrip(req)
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
   val := req.URL.Query()
   val.Set("clips_media_shortcode", shortcode)
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Authorization", l.Get("Ig-Set-Authorization"))
   req.Header.Set("User-Agent", userAgent)
   res, err := mech.RoundTrip(req)
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
      Edge_Sidecar_To_Children *struct {
         Edges []Edge
      }
      Video_URL string
   }
}

// If `auth` is `nil`, then anonymous request will be used.
func GraphQL(shortcode string, auth *Login) (*Media, error) {
   req, err := http.NewRequest("GET", OriginWWW + "/p/" + shortcode + "/", nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("__a", "1")
   req.URL.RawQuery = val.Encode()
   req.Header = http.Header{
      "User-Agent": {userAgent},
   }
   if auth != nil && auth.Header != nil {
      req.Header.Set("Authorization", auth.Get("Ig-Set-Authorization"))
   }
   res, err := mech.RoundTrip(req)
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
   var q Query
   q.Query_Hash = queryHash
   q.Variables.Shortcode = shortcode
   return q
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
   res, err := mech.RoundTrip(req)
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
