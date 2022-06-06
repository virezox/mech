package roku
// github.com/89z

import (
   "bytes"
   "errors"
   "github.com/89z/format/json"
   "net/http"
)

type CrossSite struct {
   cookie *http.Cookie // has own String method
   token string
}

func NewCrossSite() (*CrossSite, error) {
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
   var site CrossSite
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

func (c CrossSite) Playback(id string) (*Playback, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "mediaFormat": "mpeg-dash",
      "rokuId": id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://therokuchannel.roku.com/api/v3/playback", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
   }
   req.AddCookie(c.cookie)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   play := new(Playback)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}
