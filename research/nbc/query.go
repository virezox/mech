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
      ...titleV2HomepagePageData
    }
    analytics {
      ...pageAnalyticsAttributes
      ... on VideoPageAnalyticsAttributes {
        convivaAssetName
      }
    }
  }
}

fragment pageAnalyticsAttributes on PageAnalyticsAttributes {
  series
  favoritedSeries
  brand {
    title
  }
}

fragment videoPageData on VideoPageData {
  mpxAccountId
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
`
