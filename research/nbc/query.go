package main

const query = `
query bonanzaPage(
  $id: ID
  $name: String!
  $type: EntityPageType!
  $nationalBroadcastType: String
  $userId: String!
  $platform: SupportedPlatforms!
  $device: String
  $profile: JSON
  $ld: Boolean
  $oneApp: Boolean
  $timeZone: String
  $deepLinkHandle: String
  $app: NBCUBrands!
  $nbcAffiliateName: String
  $telemundoAffiliateName: String
  $language: Languages
  $playlistMachineName: String
  $mpxGuid: String
  $authorized: Boolean
  $minimumTiles: Int
  $endCardMpxGuid: String
  $endCardTagLine: String
  $isDayZero: Boolean
) {
  bonanzaPage(
    id: $id
    name: $name
    type: $type
    nationalBroadcastType: $nationalBroadcastType
    userId: $userId
    platform: $platform
    device: $device
    profile: $profile
    ld: $ld
    oneApp: $oneApp
    timeZone: $timeZone
    deepLinkHandle: $deepLinkHandle
    app: $app
    nbcAffiliateName: $nbcAffiliateName
    telemundoAffiliateName: $telemundoAffiliateName
    language: $language
    playlistMachineName: $playlistMachineName
    mpxGuid: $mpxGuid
    authorized: $authorized
    minimumTiles: $minimumTiles
    endCardMpxGuid: $endCardMpxGuid
    endCardTagLine: $endCardTagLine
    isDayZero: $isDayZero
  ) {
    id
    pageType
    name
    data {
      sections {
         ...error
         ...videoPlayer
      }
    }
  }
}

fragment error on Error {
  error
}

fragment videoPlayer on VideoPlayer {
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
`
