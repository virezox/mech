package roku

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func newCrossSite() (*crossSite, error) {
   req, err := http.NewRequest("GET", "https://www.roku.com", nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var site crossSite
   for _, cook := range res.Cookies() {
      if cook.Name == "_csrf" {
         site.cookie = cook
      }
   }
   scan, err := json.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte("\tcsrf:")
   scan.Scan()
   scan.Split = nil
   if err := scan.Decode(&site.token); err != nil {
      return nil, err
   }
   return &site, nil
}

type crossSite struct {
   cookie *http.Cookie // has own String method
   token string
}

func (c crossSite) playback() (*http.Response, error) {
   body := strings.NewReader(`
   {
      "mediaFormat": "mpeg-dash",
      "rokuId": "597a64a4a25c5bf6af4a8c7053049a6f"
   }
   `)
   req := new(http.Request)
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "therokuchannel.roku.com"
   req.URL.Path = "/api/v3/playback"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   //req.Header["Csrf-Token"] = []string{"f1QK2Tae-aS1mNvaUvL6x9L8ZvbkbFZ3qk08"}
   req.Header.Set("CSRF-Token", c.token)
   //req.Header["Cookie"] = []string{"_csrf=47kBPz2aUAj3SrRgMwfkK2iu"}
   req.AddCookie(c.cookie)
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

var LogLevel format.LogLevel
