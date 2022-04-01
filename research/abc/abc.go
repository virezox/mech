package abc

import (
   "encoding/json"
   "fmt"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

var LogLevel format.LogLevel

type Route struct {
   Modules []struct {
      Resource string
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

func (r Route) Video() (*Video, error) {
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
      var play struct {
         Video Video
      }
      if err := json.NewDecoder(res.Body).Decode(&play); err != nil {
         return nil, err
      }
      return &play.Video, nil
   }
   return nil, notFound{"resource"}
}

type Video struct {
   ID string
   Show struct {
      Title string
   }
   Title string
   Assets []struct {
      Format string
      Value string
   }
}

func (v *Video) Authorize() error {
   var addr strings.Builder
   addr.WriteString("http://api.entitlement.watchabc.go.com")
   addr.WriteString("/vp2/ws-secure/entitlement/2020/authorize.json")
   body := url.Values{
      "brand": {"001"},
      "device": {"031_04"},
      "video_id": {v.ID},
      "video_type": {"lf"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", addr.String(), strings.NewReader(body),
   )
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var auth struct {
      UplynkData struct {
         SessionKey string
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&auth); err != nil {
      return err
   }
   for i, asset := range v.Assets {
      addr, err := url.Parse(asset.Value)
      if err != nil {
         return err
      }
      addr.RawQuery = auth.UplynkData.SessionKey
      v.Assets[i].Value = addr.String()
   }
   return nil
}

func (v Video) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "ID:", v.ID)
   fmt.Fprintln(f, "Show:", v.Show.Title)
   fmt.Fprint(f, "Title: ", v.Title)
   if verb == 'a' {
      for _, asset := range v.Assets {
         fmt.Fprint(f, "\nFormat:", asset.Format)
         fmt.Fprint(f, " Value:", asset.Value)
      }
   }
}

func (v Video) ULNK() string {
   for _, asset := range v.Assets {
      if asset.Format == "ULNK" {
         return asset.Value
      }
   }
   return ""
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return fmt.Sprintf("%q is not found", n.value)
}
