package instagram

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
   "strings"
   "time"
)

type EdgeMedia struct {
   Edges []struct {
      Node struct {
         Text string
      }
   }
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

type GraphMedia struct {
   Display_URL string
   Edge_Media_To_Caption EdgeMedia
   Edge_Media_To_Parent_Comment EdgeMedia
   Edge_Sidecar_To_Children struct {
      Edges []struct {
         Node struct {
            Display_URL string
            Video_URL string
         }
      }
   }
   Owner struct {
      Username string
   }
   Taken_At_Timestamp int64
   Video_URL string
}

func (g GraphMedia) Time() time.Time {
   return time.Unix(g.Taken_At_Timestamp, 0)
}

var LogLevel format.LogLevel

func Shortcode(address string) string {
   var prev string
   for _, split := range strings.Split(address, "/") {
      if prev == "p" {
         return split
      }
      prev = split
   }
   return ""
}

// Anonymous request
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

type User struct {
   Edge_Followed_By struct {
      Count int64
   }
   Edge_Follow struct {
      Count int64
   }
}

// Use Authorization
func (l Login) User(username string) (*User, error) {
   var buf strings.Builder
   buf.WriteString("https://www.instagram.com/")
   buf.WriteString(username)
   buf.WriteByte('/')
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", Android.String())
   if l.Authorization != "" {
      req.Header.Set("Authorization", l.Authorization)
   }
   req.URL.RawQuery = "__a=1"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var profile struct {
      GraphQL struct {
         User User
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
      return nil, err
   }
   return &profile.GraphQL.User, nil
}

// Anonymous request
func NewUser(username string) (*User, error) {
   return Login{}.User(username)
}

func (u User) String() string {
   buf := []byte("Followers: ")
   buf = strconv.AppendInt(buf, u.Edge_Followed_By.Count, 10)
   buf = append(buf, "\nFollowing: "...)
   buf = strconv.AppendInt(buf, u.Edge_Follow.Count, 10)
   return string(buf)
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
