package instagram

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
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

type Media struct {
   Display_URL string
   Video_URL string
   Edge_Sidecar_To_Children struct {
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
   var addr strings.Builder
   addr.WriteString("https://www.instagram.com/p/")
   addr.WriteString(shortcode)
   addr.WriteByte('/')
   req, err := http.NewRequest("GET", addr.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", Android.String())
   req.URL.RawQuery = "__a=1"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var post struct {
      GraphQL struct {
         Shortcode_Media Media
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
      return nil, err
   }
   return &post.GraphQL.Shortcode_Media, nil
}

func (m Media) String() string {
   var buf []byte
   for i, addr := range m.URLs() {
      if i >= 1 {
         buf = append(buf, "\n---\n"...)
      }
      buf = append(buf, addr...)
   }
   return string(buf)
}

func (m Media) URLs() []string {
   src := make(map[string]bool)
   src[m.Display_URL] = true
   src[m.Video_URL] = true
   for _, edge := range m.Edge_Sidecar_To_Children.Edges {
      src[edge.Node.Display_URL] = true
      src[edge.Node.Video_URL] = true
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
