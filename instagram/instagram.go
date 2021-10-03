package instagram

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "os"
)

const (
   Origin = "https://i.instagram.com"
   userAgent = "Instagram 207.0.0.39.120 Android"
)

var Verbose bool

func roundTrip(req *http.Request) (*http.Response, error) {
   if Verbose {
      dum, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      os.Stdout.Write(dum)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   return res, nil
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

func Decode(r io.Reader) (*Login, error) {
   var log Login
   err := json.NewDecoder(r).Decode(&log.Header)
   if err != nil {
      return nil, err
   }
   return &log, nil
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
   req, err := http.NewRequest("POST", Origin + "/api/v1/accounts/login/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {userAgent},
   }
   res, err := roundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return &Login{res.Header}, nil
}

func (l Login) Encode(w io.Writer) error {
   enc := json.NewEncoder(w)
   enc.SetIndent("", " ")
   return enc.Encode(l.Header)
}

func (l Login) Item(code string) (*Item, error) {
   req, err := http.NewRequest("GET", Origin + "/api/v1/clips/item/", nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("clips_media_shortcode", code)
   req.URL.RawQuery = val.Encode()
   req.Header.Set("User-Agent", userAgent)
   req.Header.Set("Authorization", l.Get("Ig-Set-Authorization"))
   res, err := roundTrip(req)
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

type Query struct {
   Query_Hash string `json:"query_hash"`
   Variables struct {
      Shortcode string `json:"shortcode"`
   } `json:"variables"`
}

func NewQuery(code string) Query {
   var q Query
   q.Query_Hash = "1f950d414a6e11c98c556aa007b3157d"
   q.Variables.Shortcode = code
   return q
}

// If `auth` is `nil`, then anonymous request will be used.
func (q Query) Sidecar(auth *Login) (*Sidecar, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(q)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", Origin + "/graphql/query/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "User-Agent": {userAgent},
   }
   if auth != nil {
      req.Header.Set("Authorization", auth.Get("Ig-Set-Authorization"))
   }
   res, err := roundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   car := new(Sidecar)
   if err := json.NewDecoder(res.Body).Decode(car); err != nil {
      return nil, err
   }
   return car, nil
}

type Sidecar struct {
   Data struct {
      Shortcode_Media struct {
         Edge_Sidecar_To_Children struct {
            Edges []struct {
               Node struct {
                  Display_URL string
               }
            }
         }
      }
   }
}
