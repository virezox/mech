package roku

import (
   "bytes"
   "errors"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/json"
   "github.com/89z/mech"
   "io"
   "net/http"
   "net/url"
   "path"
   "strings"
   "time"
)

type CrossSite struct {
   cookie *http.Cookie // has own String method
   token string
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

func (c Content) Base() string {
   var buf strings.Builder
   if c.Meta.MediaType == "episode" {
      buf.WriteString(c.Series.Title)
      buf.WriteByte('-')
      buf.WriteString(c.SeasonNumber)
      buf.WriteByte('-')
      buf.WriteString(c.EpisodeNumber)
      buf.WriteByte('-')
   }
   buf.WriteString(mech.Clean(c.Title))
   return buf.String()
}

type Content struct {
   Meta struct {
      Id string
      MediaType string
   }
   Title string
   Series struct {
      Title string
   }
   SeasonNumber string
   EpisodeNumber string
   ReleaseDate string
   RunTimeSeconds int64
   ViewOptions []struct {
      License string
      Media struct {
         Videos []Video
      }
   }
}

func ContentId(addr string) string {
   return path.Base(addr)
}

func NewContent(id string) (*Content, error) {
   var addr url.URL
   addr.Scheme = "https"
   addr.Host = "content.sr.roku.com"
   addr.Path = "/content/v1/roku-trc/" + id
   addr.RawQuery = url.Values{
      "expand": {"series"},
      "include": {strings.Join([]string{
         "episodeNumber",
         "releaseDate",
         "runTimeSeconds",
         "seasonNumber",
         // this needs to be exactly as is, otherwise size blows up
         "series.seasons.episodes.viewOptions\u2008",
         "series.title",
         "title",
         "viewOptions",
      }, ",")},
   }.Encode()
   var buf strings.Builder
   buf.WriteString("https://therokuchannel.roku.com/api/v2/homescreen/content/")
   buf.WriteString(url.PathEscape(addr.String()))
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

func (c Content) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "ID:", c.Meta.Id)
   fmt.Fprintln(f, "Type:", c.Meta.MediaType)
   fmt.Fprintln(f, "Title:", c.Title)
   if c.Meta.MediaType == "episode" {
      fmt.Fprintln(f, "Series:", c.Series.Title)
      fmt.Fprintln(f, "Season:", c.SeasonNumber)
      fmt.Fprintln(f, "Episode:", c.EpisodeNumber)
   }
   fmt.Fprintln(f, "Date:", c.ReleaseDate)
   fmt.Fprint(f, "Duration: ", c.Duration())
   if verb == 'a' {
      for _, opt := range c.ViewOptions {
         fmt.Fprint(f, "\nLicense: ", opt.License)
         for _, vid := range opt.Media.Videos {
            fmt.Fprint(f, "\nURL: ", vid.Url)
         }
      }
   }
}

var LogLevel format.LogLevel

func (c Content) Duration() time.Duration {
   return time.Duration(c.RunTimeSeconds) * time.Second
}

type Video struct {
   DrmAuthentication *struct{}
   VideoType string
   Url string
}

func (c Content) DASH() *Video {
   for _, opt := range c.ViewOptions {
      for _, vid := range opt.Media.Videos {
         if vid.VideoType == "DASH" {
            return &vid
         }
      }
   }
   return nil
}

func (c Content) Hls() (*Video, error) {
   for _, opt := range c.ViewOptions {
      for _, vid := range opt.Media.Videos {
         if vid.DrmAuthentication == nil {
            if vid.VideoType == "HLS" {
               return &vid, nil
            }
         }
      }
   }
   return nil, errors.New("drmAuthentication")
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
   var scan json.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte("\tcsrf:")
   scan.Scan()
   scan.Sep = nil
   if err := scan.Decode(&site.token); err != nil {
      return nil, err
   }
   return &site, nil
}
