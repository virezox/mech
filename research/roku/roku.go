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
      "featureInclude":[]string{"bookmark,watchlist,linearSchedule"},
      "filter":[]string{"categoryObjects:genreAppropriate%20eq%20true"},
      "expand":[]string{"categoryObjects,credits,series,series.seasons,series.seasons.episodes,seasons,seasons.episodes,next,viewOptions,viewOptions.providerDetails"},
      "include":[]string{"title,type,images,imageMap.detailPoster,imageMap.detailBackground,categoryObjects,runTimeSeconds,castAndCrew,savable,stationDma,seasons.castAndCrew,kidsDirected,releaseDate,releaseYear,episodeNumber,seasonNumber,description,descriptions,indicators,genres,credits.birthDate,credits.meta,credits.order,credits.name,credits.role,credits.personId,credits.images,parentalRatings,reverseChronological,contentRatingClass,viewOptions,seasons.title,seasons.seasonNumber,seasons.description,seasons.descriptions,seasons.releaseYear,seasons.credits.birthDate,seasons.credits.meta,seasons.credits.order,seasons.credits.name,seasons.credits.role,seasons.credits.personId,seasons.credits.images,seasons.images,seasons.imageMap.detailBackground,seasons.episodes,seasons.episodes.title,seasons.episodes.description,seasons.episodes.descriptions.40,seasons.episodes.descriptions.60,seasons.episodes.episodeNumber,seasons.episodes.seasonNumber,seasons.episodes.images,seasons.episodes.imageMap.grid,seasons.episodes.indicators,seasons.episodes.releaseDate,seasons.episodes.viewOptions,series.title,series.seasons,series.seasons.title,series.seasons.seasonNumber,series.seasons.description,series.seasons.descriptions,series.seasons.releaseYear,series.seasons.credits.birthDate,series.seasons.credits.meta,series.seasons.credits.order,series.seasons.credits.name,series.seasons.credits.role,series.seasons.credits.personId,series.seasons.credits.images,series.seasons.images,series.seasons.imageMap.detailBackground,series.seasons.episodes,series.seasons.episodes.title,series.seasons.episodes.description,series.seasons.episodes.descriptions.40,series.seasons.episodes.descriptions.60,series.seasons.episodes.episodeNumber,series.seasons.episodes.seasonNumber,series.seasons.episodes.images,series.seasons.episodes.imageMap.grid,series.seasons.episodes.imageMap.detailBackground,series.seasons.episodes.indicators,series.seasons.episodes.releaseDate,series.seasons.episodes.viewOptions\u2008"},
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

