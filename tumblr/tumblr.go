package tumblr

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
   "strings"
)

const consumerKey = "BUHsuO5U9DF42uJtc8QTZlOmnUaJmBJGuU1efURxeklbdiLn9L"

var LogLevel format.LogLevel

type Element struct {
   Blog struct {
      Title string
   }
   Summary string
   Thumbnail_Width int64
   Thumbnail_Height int64
   Duration int64
   Video_URL string
}

func (e Element) String() string {
   buf := []byte("Width: ")
   buf = strconv.AppendInt(buf, e.Thumbnail_Width, 10)
   buf = append(buf, "\nHeight: "...)
   buf = strconv.AppendInt(buf, e.Thumbnail_Height, 10)
   buf = append(buf, "\nDuration: "...)
   buf = strconv.AppendInt(buf, e.Duration, 10)
   buf = append(buf, "\nBlog: "...)
   buf = append(buf, e.Blog.Title...)
   buf = append(buf, "\nSummary: "...)
   buf = append(buf, e.Summary...)
   buf = append(buf, "\nURL: "...)
   buf = append(buf, e.Video_URL...)
   return string(buf)
}

type BlogPost struct {
   Response struct {
      Timeline struct {
         Elements []Element
      }
   }
}

type Permalink struct {
   BlogName string
   PostID int64
}

func NewPermalink(address string) (*Permalink, error) {
   fields := strings.FieldsFunc(address, func(r rune) bool {
      return r == '/' || r == '.'
   })
   var (
      link Permalink
      prev string
   )
   for _, field := range fields {
      switch prev {
      case "https:":
         link.BlogName = field
      case "post":
         var err error
         link.PostID, err = strconv.ParseInt(field, 10, 64)
         if err != nil {
            return nil, err
         }
      }
      prev = field
   }
   return &link, nil
}

func (p Permalink) BlogPost() (*BlogPost, error) {
   buf := []byte("https://api-http2.tumblr.com/v2/blog/")
   buf = append(buf, p.BlogName...)
   buf = append(buf, "/posts/"...)
   buf = strconv.AppendInt(buf, p.PostID, 10)
   buf = append(buf, "/permalink"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   buf = []byte(`OAuth oauth_consumer_key="`)
   buf = append(buf, consumerKey...)
   buf = append(buf, '"')
   req.Header.Set("Authorization", string(buf))
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   post := new(BlogPost)
   if err := json.NewDecoder(res.Body).Decode(post); err != nil {
      return nil, err
   }
   return post, nil
}
