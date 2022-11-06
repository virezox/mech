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
    metadata {
      __typename
      ...videoPageData
      ...seriesHomepagePageData
      ...movieHomepagePageData
      ...brandLandingPageMetadata
      ...titleV2HomepagePageData
    }
    analytics {
      ...brandLandingPageAnalyticsAttributes
      ...pageAnalyticsAttributes
      ...videoPageAnalyticsAttributes
      ...titlePageAnalyticsAttributes
      ...titleV2PageAnalyticsAttributes
    }
    data {
      featured {
        ...slideshow
        ...hero
        ...ctaHero
      }
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


fragment component on Component {
  component
  meta
  treatments
}
fragment section on Section {
  logicName
  deepLinkHandle
}

fragment slideshow on Slideshow {
  ...component
  ...section
  analytics {
    itemsList
    machineName
    listTitle
  }
}
fragment hero on Hero {
  ...component
  ...section
  data {
    ...heroData
  }
}
fragment ctaHero on CTAHero {
  ...component
  ...section
}
fragment videoPlayer on VideoPlayer {
  ...component
  ...section
  data {
    ...componentData
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
    endCard {
      ...lazyEndCard
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
fragment lazyEndCard on LazyEndCard {
  ...component
}
fragment hypermediaLink on HypermediaLink {
  linkTitle
  rel
  request {
    location
    method
    body
    headers {
      name
      value
    }
  }
}

fragment smartTile on SmartTile {
  ...component
  analytics {
    smartTileLabel
    smartTileLogic
    editorialLogic
    smartTileScenario
  }
}

fragment componentData on ComponentData {
  instanceID
}

fragment brandLandingPageAnalyticsAttributes on BrandLandingPageAnalyticsAttributes {
  brand {
    title
  }
}
fragment pageAnalyticsAttributes on PageAnalyticsAttributes {
  series
  favoritedSeries
  brand {
    title
  }
}
fragment titlePageAnalyticsAttributes on TitlePageAnalyticsAttributes {
  series
  movie
  favoritedSeries
  brand {
    title
  }
}
fragment titleV2PageAnalyticsAttributes on TitleV2PageAnalyticsAttributes {
  series
  movie
  favoritedSeries
  genre
  secondaryGenres
  isTitleHub
  titleHub
  isSponsoredTitle
  sponsorName
  hasTrailerCTA
  pageType
  brand {
    title
  }
  emptyStateErrorDescription
  isEmpty
  category
  seriesType
  dartTag
  v4ID
  listOfGenres
}
fragment videoPageAnalyticsAttributes on VideoPageAnalyticsAttributes {
  series
  favoritedSeries
  brand {
    title
  }
  title
  movie
  programmingType
  episodeNumber
  seasonNumber
  mpxGuid
  locked
  entitlement
  duration
  playlistTitle
  playlistPosition
  tmsId
  isLongFormContent
  durationInMilliseconds
  airDate
  dayPart
  webBrandDomain
  permalink
  genre
  secondaryGenre
  titleTmsId
  clipCategory
  adobeVideoPlatform
  adobeContentType
  adobeBrand
  convivaAssetName
  videoBroadcast
  isOlympics
  listOfGenres
  rating
  ratingAdvisories
  nielsenProgen
  nielsenBrand
  nielsenSfCode
  nielsenClientId
  nielsenChannel
  sport
  event
  league
  language
}
fragment videoPageData on VideoPageData {
  title
  secondaryTitle
  tertiaryTitle
  playlistTitle
  playlistImage {
    path
    width
    height
  }
  description
  shortDescription
  gradientStart
  gradientEnd
  lightPrimaryColor
  brandLightPrimaryColor
  brandDarkPrimaryColor
  seasonNumber
  episodeNumber
  airDate
  rating
  copyright
  locked
  programmingType
  genre
  duration
  permalink
  percentViewed
  labelBadge
  mpxGuid
  authEnds
  externalAdvertiserId
  mpxEntitlementWindows {
    availStartDateTime
    availEndDateTime
    entitlement
    device
  }
  tveEntitlementWindows {
    availStartDateTime
    availEndDateTime
    entitlement
    device
  }
  cast {
    characterFirstName
    characterLastName
    talentFirstName
    talentLastName
  }
  v4ID
  seriesShortTitle
  seriesShortDescription
  multiPlatformLargeImage
  multiPlatformSmallImage
  urlAlias
  dartTag
  seriesType
  dayPart
  sunrise
  sunset
  ratingAdvisories
  width
  height
  selectedCountries
  keywords
  watchId
  referenceUrl
  numberOfEpisodes
  numberOfSeasons
  channelId
  resourceId
  mpxAccountId
  mpxAdPolicy
  brandDisplayTitle
  colorBrandLogo
  whiteBrandLogo
}
fragment brandLandingPageMetadata on BrandLandingPageMetadata {
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  lightPrimaryColor
  darkPrimaryColor
  brandLandingHeadline
  brandLandingDescription
  brandLandingLogo
  brandLandingBackgroundImage
  brandLandingBackgroundPreview
}
fragment titleV2HomepagePageData on TitleV2HomepagePageData {
  pageType
  isEmpty
  socialMedia {
    name
    url
    handle
  }
  referenceUrl
  shortDescription
  shortTitle
  isCoppaCompliant
  schemaType
  numberOfEpisodes
  numberOfSeasons
  brandDisplayTitle
}
fragment movieHomepagePageData on MovieHomepagePageData {
  gradientStart
  gradientEnd
  lightPrimaryColor
  darkPrimaryColor
  brandLightPrimaryColor
  brandDarkPrimaryColor
  genres
  category
  socialMedia {
    name
    url
    handle
  }
  dartTag
  referenceUrl
  description
  shortDescription
  shortTitle
  isCoppaCompliant
  schemaType
  v4ID
  titleArt {
    path
    width
    height
  }
  multiPlatformLargeImage
  multiPlatformSmallImage
  credits {
    personFirstName
    personLastName
    characterFirstName
    characterLastName
  }
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  canonicalUrl
  titleLogo
}
fragment seriesHomepagePageData on SeriesHomepagePageData {
  gradientStart
  gradientEnd
  lightPrimaryColor
  darkPrimaryColor
  brandLightPrimaryColor
  brandDarkPrimaryColor
  genres
  category
  seriesType
  socialMedia {
    name
    url
    handle
  }
  dartTag
  referenceUrl
  availableSeasons
  description
  shortDescription
  shortTitle
  isCoppaCompliant
  schemaType
  v4ID
  titleArt {
    path
    width
    height
  }
  multiPlatformLargeImage
  multiPlatformSmallImage
  credits {
    personFirstName
    personLastName
    characterFirstName
    characterLastName
  }
  numberOfEpisodes
  numberOfSeasons
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  canonicalUrl
  titleLogo
}
fragment heroData on HeroData {
  ...componentData
  title
  secondaryTitle
  compactImage
  favoriteID
  favoriteInteraction {
    default {
      ...hypermediaLink
    }
    undo {
      ...hypermediaLink
    }
  }
  smartTile {
    ...smartTile
  }
}
`
