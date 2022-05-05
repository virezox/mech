package roku

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
   "io"
)

type Content struct {
   EpisodeNumber string
   ReleaseDate string
   RunTimeSeconds int64
   SeasonNumber string
   Series struct {
      Title string
   }
   Title string
   ViewOptions []struct {
      Media struct {
         Videos []struct {
            URL string
         }
      }
   }
}

func NewContent() (*Content, error) {
   var addr url.URL
   addr.Scheme = "https"
   addr.Host = "content.sr.roku.com"
   addr.Path = "/content/v1/roku-trc/105c41ea75775968b670fbb26978ed76"
   addr.RawQuery = url.Values{
      "expand":[]string{"series.seasons.episodes"},
      "include":[]string{strings.Join([]string{
         "episodeNumber",
         "releaseDate",
         "runTimeSeconds",
         "seasonNumber",
         // this need to be exactly as is, otherwise size blows up
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
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   println(len(body))
   con := new(Content)
   if err := json.Unmarshal(body, con); err != nil {
      return nil, err
   }
   /*
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   */
   return con, nil
}

var LogLevel format.LogLevel

