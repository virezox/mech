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
      id
      pageType
      name
      data {
         sections {
            ... on Error {
              error
            }
            ...on VideoPlayer {
               data {
               player {
               v4ID
               mpxGuid
               locked
               programmingType
               title
               secondaryTitle
               tertiaryTitle
               seasonNumber
               episodeNumber
               genre
               amazonGenre
               duration
               percentViewed
               airDate
               copyright
               dayPart
               description
               shortDescription
               sunset
               keywords
               permalink
               rating
               seriesShortTitle
               seriesUrlAlias
               lightPrimaryColor
               gradientStart
               gradientEnd
               whiteBrandLogo
               colorBrandLogo
               brandDisplayTitle
               brandMachineName
               mpxAccountId
               mpxAdPolicy
               resourceId
               allowMobileWebPlayback
               ariaLabel
               startRecapTiming
               endRecapTiming
               startTeaserTiming
               endTeaserTiming
               startIntroTiming
               endIntroTiming
               cuePoint
               externalAdvertiserId
               ratingAdvisories
               tmsId
               movieShortTitle
               allowSkipButtons
               skipButtonsDuration
               }
               }
               analytics {
               tmsId
               isLongFormContent
               mpxGuid
               durationInMilliseconds
               airDate
               dayPart
               webBrandDomain
               episodeNumber
               permalink
               series
               movie
               seasonNumber
               genre
               secondaryGenre
               title
               locked
               titleTmsId
               clipCategory
               adobeVideoPlatform
               adobeContentType
               adobeBrand
               convivaAssetName
               videoBroadcast
               isOlympics
               listOfGenres
               programmingType
               rating
               ratingAdvisories
               duration
               nielsenProgen
               nielsenBrand
               nielsenSfCode
               nielsenClientId
               nielsenChannel
               ottPlatform
               league
               sport
               event
               language
               brand {
               title
               }
               }
               }
         }
      }
   }
}
`
