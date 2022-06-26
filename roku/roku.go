package roku

import (
   "bytes"
   "errors"
   "github.com/89z/format/http"
   "github.com/89z/format/json"
   "io"
   "net/url"
   "path"
   "strings"
   "time"
)

func (c Content) String() string {
   var buf strings.Builder
   buf.WriteString("ID: ")
   buf.WriteString(c.Meta.ID)
   buf.WriteString("\nType: ")
   buf.WriteString(c.Meta.MediaType)
   buf.WriteString("\nTitle: ")
   buf.WriteString(c.Title)
   if c.Meta.MediaType == "episode" {
      buf.WriteString("\nSeries: ")
      buf.WriteString(c.Series.Title)
      buf.WriteString("\nSeason: ")
      buf.WriteString(c.SeasonNumber)
      buf.WriteString("\nEpisode: ")
      buf.WriteString(c.EpisodeNumber)
   }
   buf.WriteString("\nDate: ")
   buf.WriteString(c.ReleaseDate)
   buf.WriteString("\nDuration: ")
   buf.WriteString(c.Duration().String())
   return buf.String()
}

var Client = http.Default_Client

type Cross_Site struct {
   cookie *http.Cookie // has own String method
   token string
}

func (c Cross_Site) Playback(id string) (*Playback, error) {
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
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
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
   buf.WriteString(c.Title)
   return buf.String()
}

func Content_ID(addr string) string {
   return path.Base(addr)
}

func New_Content(id string) (*Content, error) {
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
   res, err := Client.Get(buf.String())
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

func (c Content) Duration() time.Duration {
   return time.Duration(c.RunTimeSeconds) * time.Second
}

type Video struct {
   DrmAuthentication *struct{}
   VideoType string
   URL string
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

func (c Content) HLS() (*Video, error) {
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

func New_Cross_Site() (*Cross_Site, error) {
   // this has smaller body than www.roku.com
   res, err := Client.Get("https://therokuchannel.roku.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var site Cross_Site
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

type Content struct {
   Meta struct {
      ID string
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
