package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   
   "fmt"
)

func main() {
   
   url.Values{
      //"filter":[]string{"categoryObjects:genreAppropriate%20eq%20true"},
      "filter":[]string{"categoryObjects:genreAppropriate eq true"},
      "featureInclude":[]string{"bookmark,watchlist,linearSchedule"},
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
         "series.seasons.episodes.viewOptions",
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
   }
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Scheme = "https"
   req.URL.Host = "therokuchannel.roku.com"
   req.URL.Path = "/api/v2/homescreen/content/https://content.sr.roku.com/content/v1/roku-trc/105c41ea75775968b670fbb26978ed76?expand=categoryObjects%2Ccredits%2Cseries%2Cseries.seasons%2Cseries.seasons.episodes%2Cseasons%2Cseasons.episodes%2Cnext%2CviewOptions%2CviewOptions.providerDetails&include=title%2Ctype%2Cimages%2CimageMap.detailPoster%2CimageMap.detailBackground%2CcategoryObjects%2CrunTimeSeconds%2CcastAndCrew%2Csavable%2CstationDma%2Cseasons.castAndCrew%2CkidsDirected%2CreleaseDate%2CreleaseYear%2CepisodeNumber%2CseasonNumber%2Cdescription%2Cdescriptions%2Cindicators%2Cgenres%2Ccredits.birthDate%2Ccredits.meta%2Ccredits.order%2Ccredits.name%2Ccredits.role%2Ccredits.personId%2Ccredits.images%2CparentalRatings%2CreverseChronological%2CcontentRatingClass%2CviewOptions%2Cseasons.title%2Cseasons.seasonNumber%2Cseasons.description%2Cseasons.descriptions%2Cseasons.releaseYear%2Cseasons.credits.birthDate%2Cseasons.credits.meta%2Cseasons.credits.order%2Cseasons.credits.name%2Cseasons.credits.role%2Cseasons.credits.personId%2Cseasons.credits.images%2Cseasons.images%2Cseasons.imageMap.detailBackground%2Cseasons.episodes%2Cseasons.episodes.title%2Cseasons.episodes.description%2Cseasons.episodes.descriptions.40%2Cseasons.episodes.descriptions.60%2Cseasons.episodes.episodeNumber%2Cseasons.episodes.seasonNumber%2Cseasons.episodes.images%2Cseasons.episodes.imageMap.grid%2Cseasons.episodes.indicators%2Cseasons.episodes.releaseDate%2Cseasons.episodes.viewOptions%2Cseries.title%2Cseries.seasons%2Cseries.seasons.title%2Cseries.seasons.seasonNumber%2Cseries.seasons.description%2Cseries.seasons.descriptions%2Cseries.seasons.releaseYear%2Cseries.seasons.credits.birthDate%2Cseries.seasons.credits.meta%2Cseries.seasons.credits.order%2Cseries.seasons.credits.name%2Cseries.seasons.credits.role%2Cseries.seasons.credits.personId%2Cseries.seasons.credits.images%2Cseries.seasons.images%2Cseries.seasons.imageMap.detailBackground%2Cseries.seasons.episodes%2Cseries.seasons.episodes.title%2Cseries.seasons.episodes.description%2Cseries.seasons.episodes.descriptions.40%2Cseries.seasons.episodes.descriptions.60%2Cseries.seasons.episodes.episodeNumber%2Cseries.seasons.episodes.seasonNumber%2Cseries.seasons.episodes.images%2Cseries.seasons.episodes.imageMap.grid%2Cseries.seasons.episodes.imageMap.detailBackground%2Cseries.seasons.episodes.indicators%2Cseries.seasons.episodes.releaseDate%2Cseries.seasons.episodes.viewOptions%E2%80%88&filter=categoryObjects%3AgenreAppropriate%2520eq%2520true&featureInclude=bookmark%2Cwatchlist%2ClinearSchedule"
   req.URL.RawPath = "/api/v2/homescreen/content/https%3A%2F%2Fcontent.sr.roku.com%2Fcontent%2Fv1%2Froku-trc%2F105c41ea75775968b670fbb26978ed76%3Fexpand%3DcategoryObjects%252Ccredits%252Cseries%252Cseries.seasons%252Cseries.seasons.episodes%252Cseasons%252Cseasons.episodes%252Cnext%252CviewOptions%252CviewOptions.providerDetails%26include%3Dtitle%252Ctype%252Cimages%252CimageMap.detailPoster%252CimageMap.detailBackground%252CcategoryObjects%252CrunTimeSeconds%252CcastAndCrew%252Csavable%252CstationDma%252Cseasons.castAndCrew%252CkidsDirected%252CreleaseDate%252CreleaseYear%252CepisodeNumber%252CseasonNumber%252Cdescription%252Cdescriptions%252Cindicators%252Cgenres%252Ccredits.birthDate%252Ccredits.meta%252Ccredits.order%252Ccredits.name%252Ccredits.role%252Ccredits.personId%252Ccredits.images%252CparentalRatings%252CreverseChronological%252CcontentRatingClass%252CviewOptions%252Cseasons.title%252Cseasons.seasonNumber%252Cseasons.description%252Cseasons.descriptions%252Cseasons.releaseYear%252Cseasons.credits.birthDate%252Cseasons.credits.meta%252Cseasons.credits.order%252Cseasons.credits.name%252Cseasons.credits.role%252Cseasons.credits.personId%252Cseasons.credits.images%252Cseasons.images%252Cseasons.imageMap.detailBackground%252Cseasons.episodes%252Cseasons.episodes.title%252Cseasons.episodes.description%252Cseasons.episodes.descriptions.40%252Cseasons.episodes.descriptions.60%252Cseasons.episodes.episodeNumber%252Cseasons.episodes.seasonNumber%252Cseasons.episodes.images%252Cseasons.episodes.imageMap.grid%252Cseasons.episodes.indicators%252Cseasons.episodes.releaseDate%252Cseasons.episodes.viewOptions%252Cseries.title%252Cseries.seasons%252Cseries.seasons.title%252Cseries.seasons.seasonNumber%252Cseries.seasons.description%252Cseries.seasons.descriptions%252Cseries.seasons.releaseYear%252Cseries.seasons.credits.birthDate%252Cseries.seasons.credits.meta%252Cseries.seasons.credits.order%252Cseries.seasons.credits.name%252Cseries.seasons.credits.role%252Cseries.seasons.credits.personId%252Cseries.seasons.credits.images%252Cseries.seasons.images%252Cseries.seasons.imageMap.detailBackground%252Cseries.seasons.episodes%252Cseries.seasons.episodes.title%252Cseries.seasons.episodes.description%252Cseries.seasons.episodes.descriptions.40%252Cseries.seasons.episodes.descriptions.60%252Cseries.seasons.episodes.episodeNumber%252Cseries.seasons.episodes.seasonNumber%252Cseries.seasons.episodes.images%252Cseries.seasons.episodes.imageMap.grid%252Cseries.seasons.episodes.imageMap.detailBackground%252Cseries.seasons.episodes.indicators%252Cseries.seasons.episodes.releaseDate%252Cseries.seasons.episodes.viewOptions%25E2%2580%2588%26filter%3DcategoryObjects%253AgenreAppropriate%252520eq%252520true%26featureInclude%3Dbookmark%252Cwatchlist%252ClinearSchedule"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
