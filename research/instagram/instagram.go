package instagram

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

func post() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["User-Agent"] = []string{"Instagram 214.1.0.29.120 Android"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "i.instagram.com"
   req.URL.Path = "/graphql/query/"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{auth}
   req.Body = io.NopCloser(strings.NewReader(`
   {
      "query_hash": "7d4d42b121a214d23bd43206e5142c8c",
      "variables": {
         "shortcode": "CY7zg-ulZEZ",
         "fetch_comment_count": 9
      }
   }
   `))
   return new(http.Transport).RoundTrip(&req)
}

type info struct {
   Data struct {
      Shortcode_Media struct {
         Display_URL string
      }
   }
}

type old struct {
   Display_URL string
   Edge_Media_Preview_Like struct { // Likes
      Count int64
   }
   Edge_Media_To_Parent_Comment struct { // Comments
      Edges []struct {
         Node struct {
            Text string
         }
      }
   }
   Edge_Sidecar_To_Children *struct { // Sidecar
      Edges []struct {
         Node struct {
            Display_URL string
            Video_URL string
         }
      }
   }
   Video_URL string
}
