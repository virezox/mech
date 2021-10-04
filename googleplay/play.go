package googleplay

import (
   "bytes"
   "fmt"
   "github.com/89z/mech"
   "io"
   "net/http"
   "net/url"
)

const Origin = "https://android.clients.google.com"

var Verbose = mech.Verbose

// text/plain encoding algorithm
// html.spec.whatwg.org/multipage/form-control-infrastructure.html
func ParseQuery(query []byte) url.Values {
   res := make(url.Values)
   for _, pair := range bytes.Split(query, []byte{'\n'}) {
      nv := bytes.SplitN(pair, []byte{'='}, 2)
      res.Add(string(nv[0]), string(nv[1]))
   }
   return res
}

type Ac2dm struct {
   url.Values
}

// Exchange embedded token (oauth2_4) for refresh token (aas_et).
// accounts.google.com/EmbeddedSetup
func NewAc2dm(token string) (*Ac2dm, error) {
   req, err := http.NewRequest("POST", Origin + "/auth", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("ACCESS_TOKEN", "1")
   q.Set("service", "ac2dm")
   q.Set("Token", token)
   req.URL.RawQuery = q.Encode()
   fmt.Println(req.Method, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   query, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &Ac2dm{
      ParseQuery(query),
   }, nil
}

// Exchange refresh token (aas_et) for access token (Auth).
func (a Ac2dm) OAuth2() (*OAuth2, error) {
   req, err := http.NewRequest("POST", Origin + "/auth", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("Token", a.Get("Token"))
   q.Set("service", "oauth2:https://www.googleapis.com/auth/googleplay")
   req.URL.RawQuery = q.Encode()
   fmt.Println(req.Method, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   query, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &OAuth2{
      ParseQuery(query),
   }, nil
}

type OAuth2 struct {
   url.Values
}

// device is Google Service Framework.
func (o OAuth2) Details(device, app string) ([]byte, error) {
   req, err := http.NewRequest("GET", Origin + "/fdfe/details", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("doc", app)
   req.URL.RawQuery = q.Encode()
   req.Header.Set("Authorization", "Bearer " + o.Get("Auth"))
   req.Header.Set("X-DFE-Device-Id", device)
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}
