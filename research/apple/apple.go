package apple

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
)

var LogLevel format.LogLevel

type Config struct {
   WebBag struct {
      AppIdKey string
   }
}

func NewConfig() (*Config, error) {
   req, err := http.NewRequest(
      "GET", "https://amp-account.tv.apple.com/account/web/config", nil,
   )
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(Config)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

const (
   sf_max = 143499
   sf_min = 143441
   v_max = 58
   v_min = 50
)

func episodes() (*http.Response, error) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "tv.apple.com"
   req.URL.Path = "/api/uts/v3/episodes/umc.cmc.45cu44369hb2qfuwr3fihnr8e"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["caller"] = []string{"web"}
   val["locale"] = []string{"en-US"}
   val["pfm"] = []string{"web"}
   val["sf"] = []string{strconv.Itoa(sf_max)}
   val["v"] = []string{strconv.Itoa(v_max)}
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   // "adamId"
   return new(http.Transport).RoundTrip(req)
}
