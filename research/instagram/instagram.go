package instagram

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

func GraphQL() (*http.Response, error) {
   body := strings.NewReader(`
   {
      "query_hash":"7d4d42b121a214d23bd43206e5142c8c",
      "variables":{"shortcode":"CLHoAQpCI2i"}
   }
   `)
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "i.instagram.com"
   req.URL.Path = "/graphql/query/"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{auth}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["User-Agent"] = []string{"Instagram 214.1.0.29.120 Android"}
   return new(http.Transport).RoundTrip(&req)
}

func Media() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{auth}
   req.URL = new(url.URL)
   req.URL.Host = "i.instagram.com"
   req.URL.Path = "/api/v1/media/2506147657383710114/info/"
   req.URL.Scheme = "https"
   req.Header["User-Agent"] = []string{"Instagram 214.1.0.29.120 Android"}
   return new(http.Transport).RoundTrip(&req)
}
