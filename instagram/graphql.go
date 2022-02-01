package instagram

import (
   "bytes"
   "encoding/json"
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

type GraphQL struct {
   Edge_Media_Preview_Like struct {
      Count int64
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

type graphqlRequest struct {
   Query_Hash string `json:"query_hash"`
   Variables struct {
      Shortcode string `json:"shortcode"`
   } `json:"variables"`
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " not found"
}

// Anonymous request
func NewGraphQL(shortcode string) (*GraphQL, error) {
   var body graphqlRequest
   body.Query_Hash = queryHash
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
         Shortcode_Media GraphQL
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&med); err != nil {
      return nil, err
   }
   return &med.Data.Shortcode_Media, nil
}

func (g GraphQL) String() string {
   buf := []byte("Likes: ")
   buf = strconv.AppendInt(buf, g.Edge_Media_Preview_Like.Count, 10)
   buf = append(buf, "\nURLs: "...)
   for _, addr := range g.URLs() {
      buf = append(buf, "\n- "...)
      buf = append(buf, addr...)
   }
   return string(buf)
}

func (g GraphQL) URLs() []string {
   src := make(map[string]bool)
   src[g.Display_URL] = true
   src[g.Video_URL] = true
   if g.Edge_Sidecar_To_Children != nil {
      for _, edge := range g.Edge_Sidecar_To_Children.Edges {
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
