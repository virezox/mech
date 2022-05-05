package roku

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
   "io"
)

func NewContent() (*Content, error) {
   var addr url.URL
   addr.Scheme = "https"
   addr.Host = "content.sr.roku.com"
   addr.Path = "/content/v1/roku-trc/105c41ea75775968b670fbb26978ed76"
   addr.RawQuery = url.Values{
      "expand":[]string{strings.Join([]string{
         "categoryObjects",
         "credits",
         "next",
         "seasons",
         "seasons.episodes",
         "series",
         "series.seasons",
         "series.seasons.episodes",
         "viewOptions",
         "viewOptions.providerDetails",
      }, ",")},
      "include":[]string{strings.Join([]string{
         "castAndCrew",
         "categoryObjects",
         "contentRatingClass",
         "credits.birthDate",
         "credits.images",
         "credits.meta",
         "credits.name",
         "credits.order",
         "credits.personId",
         "credits.role",
         "description",
         "descriptions",
         "episodeNumber",
         "genres",
         "imageMap.detailBackground",
         "imageMap.detailPoster",
         "images",
         "indicators",
         "kidsDirected",
         "parentalRatings",
         "releaseDate",
         "releaseYear",
         "reverseChronological",
         "runTimeSeconds",
         "savable",
         "seasonNumber",
         "seasons.castAndCrew",
         "seasons.credits.birthDate",
         "seasons.credits.images",
         "seasons.credits.meta",
         "seasons.credits.name",
         "seasons.credits.order",
         "seasons.credits.personId",
         "seasons.credits.role",
         "seasons.description",
         "seasons.descriptions",
         "seasons.episodes",
         "seasons.episodes.description",
         "seasons.episodes.descriptions.40",
         "seasons.episodes.descriptions.60",
         "seasons.episodes.episodeNumber",
         "seasons.episodes.imageMap.grid",
         "seasons.episodes.images",
         "seasons.episodes.indicators",
         "seasons.episodes.releaseDate",
         "seasons.episodes.seasonNumber",
         "seasons.episodes.title",
         "seasons.episodes.viewOptions",
         "seasons.imageMap.detailBackground",
         "seasons.images",
         "seasons.releaseYear",
         "seasons.seasonNumber",
         "seasons.title",
         "series.seasons",
         "series.seasons.credits.birthDate",
         "series.seasons.credits.images",
         "series.seasons.credits.meta",
         "series.seasons.credits.name",
         "series.seasons.credits.order",
         "series.seasons.credits.personId",
         "series.seasons.credits.role",
         "series.seasons.description",
         "series.seasons.descriptions",
         "series.seasons.episodes",
         "series.seasons.episodes.description",
         "series.seasons.episodes.descriptions.40",
         "series.seasons.episodes.descriptions.60",
         "series.seasons.episodes.episodeNumber",
         "series.seasons.episodes.imageMap.detailBackground",
         "series.seasons.episodes.imageMap.grid",
         "series.seasons.episodes.images",
         "series.seasons.episodes.indicators",
         "series.seasons.episodes.releaseDate",
         "series.seasons.episodes.seasonNumber",
         "series.seasons.episodes.title",
         "series.seasons.episodes.viewOptions\u2008",
         "series.seasons.imageMap.detailBackground",
         "series.seasons.images",
         "series.seasons.releaseYear",
         "series.seasons.seasonNumber",
         "series.seasons.title",
         "series.title",
         "stationDma",
         "title",
         "type",
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

