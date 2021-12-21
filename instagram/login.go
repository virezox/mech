package instagram

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
)

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
   req, err := http.NewRequest("GET", originI + "/api/v1/clips/item/", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {l.Get("Ig-Set-Authorization")},
      "User-Agent": {userAgent},
   }
   req.URL.RawQuery = "clips_media_shortcode=" + url.QueryEscape(shortcode)
   LogLevel.Dump(req)
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

// If `auth` is `nil`, then anonymous request will be used.
func GraphQL(shortcode string, auth *Login) (*Media, error) {
   req, err := http.NewRequest(
      "GET", "https://www.instagram.com/p/" + shortcode + "/", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", userAgent)
   if auth != nil && auth.Header != nil {
      req.Header.Set("Authorization", auth.Get("Ig-Set-Authorization"))
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

// If `auth` is `nil`, then anonymous request will be used.
func (q Query) Data(auth *Login) (*Media, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(q)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", originI + "/graphql/query/", buf)
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
   LogLevel.Dump(req)
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
