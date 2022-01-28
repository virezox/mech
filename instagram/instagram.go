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
   origin = "https://i.instagram.com"
   queryHash = "7d4d42b121a214d23bd43206e5142c8c"
   // com.instagram.android
   userAgent = "Instagram 214.1.0.29.120 Android"
)

var LogLevel format.LogLevel

// instagram.com/p/CT-cnxGhvvO
// instagram.com/p/yza2PAPSx2
func Valid(shortcode string) bool {
   switch len(shortcode) {
   case 10, 11:
      return true
   }
   return false
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
   req, err := http.NewRequest("POST", origin + "/api/v1/accounts/login/", buf)
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
   var body mediaRequest
   body.Query_Hash = queryHash
   body.Variables.Fetch_Comment_Count = 9
   body.Variables.Shortcode = shortcode
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/graphql/query/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "User-Agent": {userAgent},
   }
   if l.Authorization != "" {
      req.Header.Set("Authorization", l.Authorization)
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var med struct {
      Data struct {
         Shortcode_Media Media
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&med); err != nil {
      return nil, err
   }
   return &med.Data.Shortcode_Media, nil
}

type Media struct {
   Edge_Media_Preview_Like struct {
      Count int64
   }
   Edge_Media_To_Comment struct {
      Edges []struct {
         Node struct {
            Text string
         }
      }
   }
   Display_URL string
   Video_URL string
   Edge_Sidecar_To_Children *struct {
      Edges []struct {
         Node struct {
            Display_URL string
            Video_URL string
         }
      }
   }
}

// Anonymous request
func NewMedia(shortcode string) (*Media, error) {
   return Login{}.Media(shortcode)
}

func (m Media) String() string {
   buf := []byte("Likes: ")
   buf = strconv.AppendInt(buf, m.Edge_Media_Preview_Like.Count, 10)
   buf = append(buf, "\nURLs: "...)
   for _, addr := range m.URLs() {
      buf = append(buf, "\n- "...)
      buf = append(buf, addr...)
   }
   buf = append(buf, "\nComments: "...)
   for _, edge := range m.Edge_Media_To_Comment.Edges {
      buf = append(buf, "\n- "...)
      buf = append(buf, edge.Node.Text...)
   }
   return string(buf)
}

func (m Media) URLs() []string {
   src := make(map[string]bool)
   src[m.Display_URL] = true
   src[m.Video_URL] = true
   if m.Edge_Sidecar_To_Children != nil {
      for _, edge := range m.Edge_Sidecar_To_Children.Edges {
         src[edge.Node.Display_URL] = true
         src[edge.Node.Video_URL] = true
      }
   }
   var dst []string
   for key := range src {
      if key != "" {
         dst = append(dst, key)
      }
   }
   return dst
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

type mediaRequest struct {
   Query_Hash string `json:"query_hash"`
   Variables struct {
      Shortcode string `json:"shortcode"`
      Fetch_Comment_Count int `json:"fetch_comment_count"`
   } `json:"variables"`
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
