package instagram

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "os"
   "path/filepath"
   "strconv"
)

const (
   originI = "https://i.instagram.com"
   // com.instagram.android
   userAgent = "Instagram 216.1.0.21.137 Android"
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
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := res.Header.Get("Ig-Set-Authorization")
   if auth == "" {
      return nil, notFound{"Ig-Set-Authorization"}
   }
   return &Login{auth}, nil
}

func OpenLogin(name string) (*Login, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   log := new(Login)
   if err := json.NewDecoder(file).Decode(log); err != nil {
      return nil, err
   }
   return log, nil
}

func (l Login) Create(name string) error {
   err := os.MkdirAll(filepath.Dir(name), os.ModeDir)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetIndent("", " ")
   return enc.Encode(l)
}

// Request with Authorization
func (l Login) Media(shortcode string) (*Media, error) {
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
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, response{res}
   }
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

// Anonymous request
func NewMedia(shortcode string) (*Media, error) {
   return Login{}.Media(shortcode)
}

func (m Media) Edges() []Edge {
   if m.Shortcode_Media.Edge_Sidecar_To_Children == nil {
      return nil
   }
   return m.Shortcode_Media.Edge_Sidecar_To_Children.Edges
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}

type response struct {
   *http.Response
}

func (r response) Error() string {
   return r.Status
}
