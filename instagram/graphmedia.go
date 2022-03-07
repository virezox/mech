package instagram

import (
   "encoding/json"
   "net/http"
   "strings"
   "time"
)

func NewGraphMedia(shortcode string) (*GraphMedia, error) {
   var buf strings.Builder
   buf.WriteString("https://www.instagram.com/p/")
   buf.WriteString(shortcode)
   buf.WriteByte('/')
   req, err := http.NewRequest("GET", buf.String(), nil)
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
         Shortcode_Media GraphMedia
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
      return nil, err
   }
   return &post.GraphQL.Shortcode_Media, nil
}

func (g GraphMedia) String() string {
   var buf []byte
   buf = append(buf, "Taken: "...)
   buf = append(buf, g.Time().String()...)
   buf = append(buf, "\nOwner: "...)
   buf = append(buf, g.Owner.Username...)
   for _, edge := range g.Edge_Media_To_Caption.Edges {
      buf = append(buf, "\nCaption: "...)
      buf = append(buf, edge.Node.Text...)
   }
   for _, edge := range g.Edge_Media_To_Parent_Comment.Edges {
      buf = append(buf, "\nComment: "...)
      buf = append(buf, edge.Node.Text...)
   }
   for _, addr := range g.URLs() {
      buf = append(buf, "\nURL: "...)
      buf = append(buf, addr...)
   }
   return string(buf)
}

func (g GraphMedia) Time() time.Time {
   return time.Unix(g.Taken_At_Timestamp, 0)
}

func (g GraphMedia) URLs() []string {
   src := make(map[string]bool)
   src[g.Display_URL] = true
   src[g.Video_URL] = true
   for _, edge := range g.Edge_Sidecar_To_Children.Edges {
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

type EdgeMedia struct {
   Edges []struct {
      Node struct {
         Text string
      }
   }
}

type GraphMedia struct {
   Edge_Media_To_Caption EdgeMedia
   Edge_Media_To_Parent_Comment EdgeMedia
   Owner struct {
      Username string
   }
   Taken_At_Timestamp int64
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
