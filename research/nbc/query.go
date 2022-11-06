package main

const query = `
query bonanzaPage(
   $app: NBCUBrands!
   $name: String!
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      platform: $platform
      type: $type
      userId: $userId
   ) {
      data {
         sections {
            ... on Error {
              error
            }
            ... on VideoPlayer {
               analytics {
                  adobeBrand
                  adobeContentType
                  adobeVideoPlatform
                  airDate
                  brand {
                     title
                  }
                  clipCategory
                  convivaAssetName
                  dayPart
                  duration
                  durationInMilliseconds
                  episodeNumber
                  event
                  genre
                  isLongFormContent
                  isOlympics
                  language
                  league
                  listOfGenres
                  locked
                  movie
                  mpxGuid
                  nielsenBrand
                  nielsenChannel
                  nielsenClientId
                  nielsenProgen
                  nielsenSfCode
                  ottPlatform
                  permalink
                  programmingType
                  rating
                  ratingAdvisories
                  seasonNumber
                  secondaryGenre
                  series
                  sport
                  title
                  titleTmsId
                  tmsId
                  videoBroadcast
                  webBrandDomain
               }
               data {
                  player {
                     airDate
                     allowMobileWebPlayback
                     allowSkipButtons
                     amazonGenre
                     ariaLabel
                     brandDisplayTitle
                     brandMachineName
                     colorBrandLogo
                     copyright
                     cuePoint
                     dayPart
                     description
                     duration
                     endIntroTiming
                     endRecapTiming
                     endTeaserTiming
                     episodeNumber
                     externalAdvertiserId
                     genre
                     gradientEnd
                     gradientStart
                     keywords
                     lightPrimaryColor
                     locked
                     movieShortTitle
                     mpxAccountId
                     mpxAdPolicy
                     mpxGuid
                     percentViewed
                     permalink
                     programmingType
                     rating
                     ratingAdvisories
                     resourceId
                     seasonNumber
                     secondaryTitle
                     seriesShortTitle
                     seriesUrlAlias
                     shortDescription
                     skipButtonsDuration
                     startIntroTiming
                     startRecapTiming
                     startTeaserTiming
                     sunset
                     tertiaryTitle
                     title
                     tmsId
                     v4ID
                     whiteBrandLogo
                  }
               }
            }
         }
      }
      id
      name
      pageType
   }
}
`
