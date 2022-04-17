package abc

import (
   "encoding/json"
   "fmt"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

const (
   MaxAppVersion = "99.99.9"
   MaxDevice = "033_05"
   MinAppVersion = "10.12.0"
   MinDevice = "031_01"
   brand = "001"
)

var LogLevel format.LogLevel

type Route struct {
   Modules []struct {
      Resource string
   }
}

func NewRoute(addr string) (*Route, error) {
   parse, err := url.Parse(addr)
   if err != nil {
      return nil, err
   }
   var buf strings.Builder
   buf.WriteString("http://prod.gatekeeper.us-abc.symphony.edgedatg.com")
   buf.WriteString("/api/ws/pluto/v1/layout/route")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Appversion", MaxAppVersion)
   req.URL.RawQuery = url.Values{
      "brand": {brand},
      "device": {MaxDevice},
      "url": {parse.Path},
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
   if len(r.Modules) == 0 {
      return nil, notFound{"resource"}
   }
   req, err := http.NewRequest("GET", r.Modules[0].Resource, nil)
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

type Video struct {
   ID string
   Show struct {
      Title string
   }
   Title string
   SeasonNumber string
   EpisodeNumber string
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
      "brand": {brand},
      "device": {MaxDevice},
      "video_id": {v.ID},
      // this can be empty, but it must be present
      "video_type": {""},
   }.Encode()
   req, err := http.NewRequest(
      "POST", addr.String(), strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
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
   return res.Body.Close()
}

func (v Video) Base() string {
   var buf strings.Builder
   buf.WriteString(v.Show.Title)
   buf.WriteByte('-')
   buf.WriteString(v.Title)
   buf.WriteByte('-')
   buf.WriteString(v.SeasonNumber)
   buf.WriteByte('-')
   buf.WriteString(v.EpisodeNumber)
   return buf.String()
}

func (v Video) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "ID:", v.ID)
   fmt.Fprintln(f, "Show:", v.Show.Title)
   fmt.Fprintln(f, "Title:", v.Title)
   fmt.Fprintln(f, "Season:", v.SeasonNumber)
   fmt.Fprint(f, "Episode: ", v.EpisodeNumber)
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
