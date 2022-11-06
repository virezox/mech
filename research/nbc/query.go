package main

const query = `
query bonanzaPage(
   $app: NBCUBrands!
   $authorized: Boolean
   $deepLinkHandle: String
   $device: String
   $endCardMpxGuid: String
   $endCardTagLine: String
   $id: ID
   $isDayZero: Boolean
   $language: Languages
   $ld: Boolean
   $minimumTiles: Int
   $mpxGuid: String
   $name: String!
   $nationalBroadcastType: String
   $nbcAffiliateName: String
   $oneApp: Boolean
   $platform: SupportedPlatforms!
   $playlistMachineName: String
   $profile: JSON
   $telemundoAffiliateName: String
   $timeZone: String
   $type: EntityPageType!
   $userId: String!
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
      ... on VideoPageData {
        mpxAccountId
      }
    }
    analytics {
      ... on VideoPageAnalyticsAttributes {
        convivaAssetName
      }
    }
  }
}
`
