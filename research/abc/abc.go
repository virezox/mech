package abc

import (
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func authorize() (*http.Response, error) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "api.entitlement.watchabc.go.com"
   req.URL.Path = "/vp2/ws-secure/entitlement/2020/authorize.json"
   req.URL.Scheme = "http"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   body := url.Values{
      "brand":[]string{"001"},
      "device":[]string{"031_04"},
      "video_id":[]string{"VDKA26847512"},
      "video_type":[]string{"lf"},
   }
   req.Body = io.NopCloser(strings.NewReader(body.Encode()))
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

var LogLevel format.LogLevel

type Route struct {
   Modules []struct {
      Resource string
   }
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}

func (r Route) Player() (*Player, error) {
   for _, mod := range r.Modules {
      req, err := http.NewRequest("GET", mod.Resource, nil)
      if err != nil {
         return nil, err
      }
      LogLevel.Dump(req)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         return nil, err
      }
      defer res.Body.Close()
      play := new(Player)
      if err := json.NewDecoder(res.Body).Decode(play); err != nil {
         return nil, err
      }
      return play, nil
   }
   return nil, notFound{"resource"}
}

type Player struct {
   Video struct {
      Assets []struct {
         Value string
      }
   }
}

func NewRoute(addr string) (*Route, error) {
   var buf strings.Builder
   buf.WriteString("http://prod.gatekeeper.us-abc.symphony.edgedatg.com")
   buf.WriteString("/api/ws/pluto/v1/layout/route")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Appversion", "10.23.1")
   req.URL.RawQuery = url.Values{
      "brand": {"001"},
      "device": {"031_04"},
      "url": {addr},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   route := new(Route)
   if err := json.NewDecoder(res.Body).Decode(route); err != nil {
      return nil, err
   }
   return route, nil
}

