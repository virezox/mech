package roku

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}

func (c crossSite) playback() (*playback, error) {
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
   req.Header.Set("CSRF-Token", c.token)
   req.AddCookie(c.cookie)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(playback)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

func newCrossSite() (*crossSite, error) {
   // this has smaller body than www.roku.com
   req, err := http.NewRequest("GET", "https://therokuchannel.roku.com", nil)
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

var LogLevel format.LogLevel
