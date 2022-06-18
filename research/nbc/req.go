package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "bytes"
   "encoding/json"
   "os"
)

const query = `
query endCard(  $id: ID  $type: VideoType!  $mpxGuid: String!  $forcedMpxGuid:
String  $userId: String!  $timeZone: String!  $platform: SupportedPlatforms!
$profile: JSON  $ld: Boolean  $oneApp: Boolean  $app: NBCUBrands!  $language:
Languages  $playlistMachineName: String  $tagLine: String) {  endCard(    id:
$id    type: $type    mpxGuid: $mpxGuid    forcedMpxGuid: $forcedMpxGuid
userId: $userId    timeZone: $timeZone    platform: $platform    profile:
$profile    ld: $ld    oneApp: $oneApp    app: $app    language: $language
playlistMachineName: $playlistMachineName    tagLine: $tagLine  ) {
...endCard    ...endTiles  }}fragment endCard on EndCard {  ...component  data
{    ...endCardData  }  analytics {    ...endCardAnalytics  }}fragment endTiles
on EndTiles {  ...component  data {    ...endTilesData  }}fragment
endCardAnalytics on EndCardAnalyticsAttributes {  brand {    title  }
recommendationType  series  movie  title  programmingType  episodeNumber
seasonNumber  mpxGuid  locked  genre  duration  playlistMachineName
endCardVariant  endCardLogic}fragment component on Component {  component  meta
treatments}fragment endCardData on EndCardData {  ...componentData  titleTitle
title  image  description  videoMetaData  tagLine  urlAlias  buttonLabel
mpxGuid  permalink  programmingType  cuePoint  locked  labelBadge
percentViewed  whiteBrandLogo  colorBrandLogo  brandDisplayTitle  mpxAccountId
mpxAdPolicy  resourceId  channelId  titleLogo  playlistMachineName  rating
alternateGroupTagLine  alternateOne {    ...endCardAlternate  }  alternateTwo {
...endCardAlternate  }  entityType  v4ID  gradientEnd  gradientStart
lightPrimaryColor  darkPrimaryColor}fragment endTilesData on EndTilesData {
...componentData  titleKeyArt  groupTagLine  tileOne {    ...endCardAlternate
}  tileTwo {    ...endCardAlternate  }  notification {    ...notification
}}fragment notification on Notification {  ...component  data {
...componentData    v4ID    machineName    headline    headlineColor    message
messageColor    logo    logoAltText    portraitImage    landscapeImage    cta {
...ctaLink    }  }  analytics {    entityTitle  }}fragment ctaLink on CTALink {
...component  data {    ...ctaData  }  analytics {    smartDynamicCTA
smartTileLabel    smartTileScenario    ctaTitle    destinationType
destination    brand {      title    }    series    movie    isMovie
videoTitle    locked    programmingType    seasonNumber    episodeNumber
mpxGuid    duration    isPlaylist    playlistMachineName    playlistTitle
isLive  }}fragment ctaData on CTAData {  ...componentData  color  text
destinationType  destination  playlistMachineName  playlistCount  urlAlias
isLive}fragment endCardAlternate on EndCardAlternate {  ...component  data {
...endCardAlternateData  }  analytics {    ...endCardAnalytics  }}fragment
endCardAlternateData on EndCardAlternateData {  ...componentData  titleKeyArt
tagLine  playlistMachineName  tile {    ...videoTile    ...seriesTile
...movieTile  }}fragment videoTile on VideoTile {  ...component  data {
...videoItem  }  analytics {    brand {      title    }    series    title
programmingType    episodeNumber    seasonNumber    mpxGuid    locked
duration    genre    movie  }}fragment seriesTile on SeriesTile {  ...component
data {    ...seriesItem  }  analytics {    series    genre    brand {
title    }  }}fragment movieTile on MovieTile {  ...component  data {
...movieItem  }  analytics {    movie    genre    brand {      title    }
}}fragment videoItem on VideoItem {  ...componentData  ...item  tertiaryTitle
locked  mpxGuid  programmingType  episodeNumber  seasonNumber  airDate
percentViewed  permalink  lastWatched  duration  genre  rating
lightPrimaryColor  darkPrimaryColor  seriesShortTitle  movieShortTitle
whiteBrandLogo  colorBrandLogo  brandDisplayTitle  mpxAccountId  mpxAdPolicy
resourceId  channelId  rating  externalAdvertiserId  ariaLabel
secondaryTitleTag  longDescription  ctaText  ctaTextColor  brandMachineName
durationBadge}fragment seriesItem on SeriesItem {  ...componentData  ...item
seriesName  shortTitle  urlAlias  favoritedOn  favoriteID  posterImage
lightPrimaryColor  darkPrimaryColor  whiteBrandLogo  colorBrandLogo
brandDisplayTitle  landscapePreview  portraitPreview}fragment movieItem on
MovieItem {  ...componentData  ...item  urlAlias  favoritedOn  favoriteID
posterImage  image  lightPrimaryColor  darkPrimaryColor  whiteBrandLogo
colorBrandLogo  brandDisplayTitle  landscapePreview  portraitPreview
rating}fragment item on Item {  v4ID  title  secondaryTitle  description  image
gradientStart  gradientEnd  labelBadge  lastModified}fragment componentData on
ComponentData {  instanceID}
`

type friendship struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"`
      MpxGuid string `json:"mpxGuid"`
      OneApp bool `json:"oneApp"`
      Platform string `json:"platform"`
      TimeZone string `json:"timeZone"`
      Type string `json:"type"`
      UserId string `json:"userId"`
   } `json:"variables"`
}

func main() {
   var f friendship
   f.Query = query
   f.Variables.App="nbc"
   f.Variables.MpxGuid="9000245869"
   f.Variables.OneApp=true
   f.Variables.Platform="android"
   f.Variables.TimeZone="America/Chicago"
   f.Variables.Type="segment"
   f.Variables.UserId="8292284999374523746"
   body := new(bytes.Buffer)
   err := json.NewEncoder(body).Encode(f)
   if err != nil {
      panic(err)
   }
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "friendship.nbc.co"
   req.URL.Path = "/v2/graphql"
   req.URL.Scheme = "https"
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
