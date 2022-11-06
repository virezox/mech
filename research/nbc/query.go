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
    experiments {
      ...ldExperiment
    }
    data {
      featured {
        ...slideshow
        ...hero
        ...ctaHero
      }
      notification {
        ...notification
      }
      sections {
         ...error
         
         ...guide
         ...marketingModule
         ...message
         ...navigationMenu
         ...onAirNowShelf
         ...premiumShelf
         ...stack
         ...tabsSelectableGroup
         ...videoPlayer
      }
    }
  }
}

fragment error on Error {
  error
}

fragment image on Image {
  path
  width
  height
}

fragment component on Component {
  component
  meta
  treatments
}

fragment potentialComponent on PotentialComponent {
  targetComponent
}

fragment section on Section {
  logicName
  deepLinkHandle
}

fragment marketingModule on MarketingModule {
  ...component
  ...section
  marketingModuleData: data {
    v4ID
    machineName
    description
    descriptionColor
    logo {
      ...image
    }
    logoAltText
    isSponsored
    sponsorName
    sponsorLogo {
      ...image
    }
    sponsorLogoAltText
    mainPreview
    mainImage {
      ...image
    }
    backgroundPreview
    backgroundFallbackImage {
      ...image
    }
    locked
    externalAdvertiserId
    badge
    gradientStart
    gradientEnd
    primaryCTA {
      ...ctaLink
    }
    secondaryCTA {
      ...ctaLink
    }
    ariaLabel
  }
  analytics {
    itemsList
    listTitle
    sponsorName
    machineName
    isSponsoredMarketingModule
    entityTitle
  }
}
fragment overlayData on OverlayData {
  content {
    ...descriptionSection
  }
}
fragment overlay on Overlay {
  ...component
  data {
    ...overlayData
  }
}
fragment descriptionSection on DescriptionSection {
  ...component
  data {
    ...descriptionData
  }
}
fragment shelf on Shelf {
  ...component
  ...section
  data {
    ...tileList
  }
  analytics {
    isPlaylist
    playlistMachineName
    listTitle
    isSponsoredContent
    sponsorName
    isMixedTiles
    machineName
    itemsList
  }
}
fragment premiumShelf on PremiumShelf {
  ...component
  ...section
  data {
    ...premiumTileList
  }
  analytics {
    itemsList
    machineName
    listTitle
    isSponsoredContent
    sponsorName
    shelfType
  }
}
fragment stack on Stack {
  ...component
  ...section
  data {
    ...tileList
  }
  analytics {
    playlistMachineName
    listTitle
    isSponsoredContent
    sponsorName
  }
}
fragment onAirNowShelf on OnAirNowShelf {
  ...component
  ...section
  data {
    ...onAirNowList
  }
  analytics {
    itemsList
    machineName
    listTitle
  }
}
fragment potentialShelf on PotentialShelf {
  ...component
  ...potentialComponent
  ...section
  data {
    ...potentialComponentData
  }
}
fragment lazyShelf on LazyShelf {
  ...component
  ...section
  ...lazyComponent
}
fragment lazyGrid on LazyGrid {
  ...component
  ...section
  ...lazyComponent
}
fragment message on Message {
  ...component
  ...section
  data {
    ...messageData
  }
}
fragment grid on Grid {
  ...component
  ...section
  data {
    ...tileList
  }
  analytics {
    listTitle
    playlistMachineName
    isSponsoredContent
    sponsorName
  }
}
fragment shelfGroup on ShelfGroup {
  ...component
  ...section
  data {
    ...shelfList
  }
}
fragment potentialShelfGroup on PotentialShelfGroup {
  ...component
  ...potentialComponent
  ...section
  data {
    ...potentialComponentData
  }
}
fragment lazyShelfGroup on LazyShelfGroup {
  ...component
  ...lazyComponent
  ...section
  data {
    ...lazyComponentData
  }
}
fragment tabsSelectableGroup on TabsSelectableGroup {
  ...component
  ...section
  data {
    ...stringSelectableComponentList
  }
}
fragment slideshow on Slideshow {
  ...component
  ...section
  data {
    ...slideList
  }
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
  data {
    ...ctaHeroData
  }
}
fragment navigationMenu on NavigationMenu {
  ...component
  ...section
  data {
    ...componentData
    favoriteInteraction {
      default {
        ...hypermediaLink
      }
      undo {
        ...hypermediaLink
      }
    }
    shortTitle
    tuneIn
    links {
      ...component
      data {
        ...componentData
        items {
          title
          href
          isCoppaCompliant
        }
      }
    }
  }
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
      image
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
  ...lazyComponent
  data {
    ...lazyComponentData
  }
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
fragment ctaLink on CTALink {
  ...component
  data {
    ...ctaData
  }
  analytics {
    smartDynamicCTA
    smartTileLabel
    smartTileLogic
    editorialLogic
    smartTileScenario
    ctaTitle
    destinationType
    destination
    brand {
      title
    }
    series
    movie
    isMovie
    videoTitle
    locked
    programmingType
    seasonNumber
    episodeNumber
    mpxGuid
    duration
    isPlaylist
    playlistMachineName
    playlistTitle
    isLive
    sponsorName
    isSponsoredTitle
    isTrailer
    liveEntitlement
    isVote
    isSportVideo
    language
    league
    event
    sport
  }
}
fragment seriesTile on SeriesTile {
  ...component
  data {
    ...seriesItem
  }
  analytics {
    series
    brand {
      title
    }
  }
}
fragment movieTile on MovieTile {
  ...component
  data {
    ...movieItem
  }
  analytics {
    movie
    brand {
      title
    }
  }
}
fragment onAirNowTile on OnAirNowTile {
  ...component
  onAirNowTileData: data {
    ...onAirNowItem
  }
  analytics {
    isLive
    episodeNumber
    seasonNumber
    programTitle
    episodeTitle
    tmsId
    liveEntitlement
    adobeVideoResearchTitle
    league
    isOlympics
    sport
    nielsenSfCode
    nielsenChannel
    nielsenClientId
    videoBroadcast
    brand {
      title
    }
  }
}
fragment upcomingLiveTile on UpcomingLiveTile {
  ...component
  data {
    ...upcomingLiveItem
  }
  analytics {
    programType
    genre
    secondaryGenre
    listOfGenres
    title
    secondaryTitle
    liveEntitlement
    league
    sport
    videoBroadcast
    nielsenClientId
    nielsenSfCode
    nielsenChannel
    adobeVideoResearchTitle
    isOlympics
    brand {
      title
    }
  }
}
fragment upcomingModal on UpcomingModal {
  ...component
  data {
    ...upcomingModalData
  }
  analytics {
    modalName
    modalType
    dismissText
    programTitle
    brand {
      title
    }
  }
}
fragment brandTile on BrandTile {
  ...component
  data {
    ...brandItem
  }
  analytics {
    brand {
      title
    }
  }
}
fragment slide on Slide {
  ...component
  data {
    ...slideItem
  }
  analytics {
    entityTitle
    isSponsoredSlide
    sponsorName
    dynamicallyGenerated
    dynamicGenerationLogic
  }
}
fragment slideTile on SlideTile {
  ...component
  data {
    ...slideItem
  }
  analytics {
    entityTitle
    isSponsoredSlide
    sponsorName
    dynamicallyGenerated
    dynamicGenerationLogic
  }
}
fragment upcomingLiveSlide on UpcomingLiveSlide {
  ...component
  data {
    ...upcomingLiveSlideData
  }
  analytics {
    analyticsType
    ctaLiveTitle
    ctaUpcomingTitle
    ctaNotInPackageTitle
    isLiveCallout
    isSponsoredSlide
    sponsorName
    programType
    genre
    secondaryGenre
    listOfGenres
    title
    secondaryTitle
    liveEntitlement
    league
    sport
    videoBroadcast
    nielsenClientId
    nielsenChannel
    nielsenSfCode
    isOlympics
    adobeVideoResearchTitle
    brand {
      title
    }
  }
}
fragment premiumTile on PremiumTile {
  ...component
  data {
    ...premiumItem
  }
  analytics {
    entityTitle
  }
}
fragment videoTile on VideoTile {
  ...component
  data {
    ...videoItem
  }
  analytics {
    brand {
      title
    }
    series
    title
    programmingType
    episodeNumber
    seasonNumber
    mpxGuid
    locked
    duration
    movie
    genre
    sport
    league
    language
    event
  }
}
fragment featureTile on FeatureTile {
  ...component
  data {
    ...featureItem
  }
  analytics {
    series
    brand {
      title
    }
    playlistMachineName
    listTitle
  }
}
fragment playlistTile on PlaylistTile {
  ...component
  data {
    ...playlistItem
  }
  analytics {
    brand {
      title
    }
    playlistMachineName
    listTitle
  }
}
fragment marketingBand on MarketingBand {
  ...component
  data {
    ...marketingBandData
  }
  analytics {
    series
    brand {
      title
    }
  }
}
fragment smartTile on SmartTile {
  ...component
  data {
    ...smartTileData
  }
  analytics {
    smartTileLabel
    smartTileLogic
    editorialLogic
    smartTileScenario
  }
}
fragment ctaSmartTile on CTASmartTile {
  ...component
  data {
    ...ctaSmartTileData
  }
  analytics {
    brand {
      title
    }
    title
    programmingType
    episodeNumber
    seasonNumber
    mpxGuid
    locked
    duration
    movie
    series
    genre
    smartTileLabel
    smartTileLogic
    editorialLogic
    smartTileScenario
    sponsorName
    isSponsoredTitle
  }
}
fragment guide on Guide {
  ...component
  data {
    ...guideData
  }
}
fragment guideSchedule on GuideSchedule {
  ...component
  data {
    ...guideScheduleData
  }
}
fragment guideStream on GuideStream {
  ...component
  data {
    ...guideStreamData
  }
  analytics {
    brand {
      title
    }
    callSign
    xyFallback
  }
}
fragment guideProgram on GuideProgram {
  ...component
  guideProgramData: data {
    ...guideProgramData
  }
  analytics {
    brand {
      title
    }
    currentVideo {
      tmsId
      mpxGuid
      rating
      relativePath
      programmingType
      seasonNumber
      episodeNumber
      duration
      videoTitle
      series
      movie
      locked
      callSign
      ottPlatform
      airDate
      webBrandDomain
      programTitle
      isFullEpisode
      durationInMilliseconds
      genre
      secondaryGenre
      externalAdvertiserId
      titleTmsId
      dayPart
      videoType
      clipCategory
      adobeVideoPlatform
      adobeContentType
      adobeBrand
      playerUrl
      liveEntitlement
      listOfGenres
      ratingWithAdvisories
      nielsenClientId
      nielsenSfCode
      nielsenProgen
      nielsenBrand
      nielsenChannel
      isOlympics
      adobeVideoResearchTitle
      sport
      league
      comscoreCallSign
      videoBroadCast
      brand {
        title
      }
    }
    nextVideo {
      tmsId
      mpxGuid
      relativePath
      programmingType
      seasonNumber
      episodeNumber
      duration
      videoTitle
      series
      movie
      locked
      brand {
        title
      }
    }
    titlePage {
      series
      movie
      isMovie
      brand {
        title
      }
    }
  }
}
fragment componentData on ComponentData {
  instanceID
}
fragment potentialComponentData on PotentialComponentData {
  ...componentData
  link
  path
}
fragment stringSelectableComponentList on StringSelectableComponentList {
  ...componentData
  initiallySelected
  itemLabels
  itemLabelsTitle
  optionalTitle: title
  items {
    ...shelf
    ...potentialShelf
    ...lazyShelf
    ...shelfGroup
    ...potentialShelfGroup
    ...lazyShelfGroup
    ...grid
  }
}
fragment ctaData on CTAData {
  ...componentData
  color
  gradientStart
  gradientEnd
  text
  destinationType
  destination
  endCardMpxGuid
  endCardTagLine
  playlistMachineName
  playlistCount
  urlAlias
  isLive
  isPlaylist
  title
  secondaryTitle
  secondaryTitleTag
  isTrailer
}
fragment messageData on MessageData {
  ...componentData
  textRow1
  textRow2
}
fragment descriptionData on DescriptionData {
  ...componentData
  optionalTitle
  description
  image
  shortTitle
  gradientStart
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
  image
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
  titleArt {
    ...image
  }
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
  image
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
fragment ctaHeroData on CTAHeroData {
  ...componentData
  title
  titleLogo {
    ...image
  }
  gradientStart
  gradientEnd
  description
  secondaryDescription
  heroImage: image {
    ...image
  }
  heroCompactImage: compactImage {
    ...image
  }
  favoriteInteraction {
    default {
      ...hypermediaLink
    }
    undo {
      ...hypermediaLink
    }
  }
  sponsorLogo {
    ...image
  }
  sponsorLogoAltText
  sponsorName
  brandDisplayTitle
  colorBrandLogo {
    ...image
  }
  whiteBrandLogo {
    ...image
  }
  smartTileCTA {
    ...ctaSmartTile
  }
  aboutOverlay {
    ...overlay
  }
  primaryCTA {
    ...ctaLink
  }
  secondaryCTA {
    ...ctaLink
  }
}
fragment smartTileData on SmartTileData {
  ...componentData
  label
  tile {
    ...videoTile
  }
}
fragment ctaSmartTileData on CTASmartTileData {
  ...componentData
  label
  title
  secondaryTitle
  secondaryTitleTag
  description
  locked
  labelBadge
  mpxGuid
  permalink
  gradientStart
  gradientEnd
  programmingType
  percentViewed
  lastWatched
  mpxAccountId
  mpxAdPolicy
  resourceId
  channelId
  externalAdvertiserId
}
fragment item on Item {
  v4ID
  title
  secondaryTitle
  tertiaryTitle
  description
  image
  gradientStart
  gradientEnd
  labelBadge
  lastModified
}
fragment videoItem on VideoItem {
  ...componentData
  ...item
  secondaryTitleTag
  locked
  mpxGuid
  programmingType
  episodeNumber
  seasonNumber
  airDate
  percentViewed
  permalink
  lastWatched
  duration
  genre
  rating
  lightPrimaryColor
  darkPrimaryColor
  seriesShortTitle
  movieShortTitle
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  mpxAccountId
  mpxAdPolicy
  resourceId
  channelId
  rating
  externalAdvertiserId
  ariaLabel
  longDescription
  ctaText
  ctaTextColor
  brandMachineName
  durationBadge
}
fragment featureItem on FeatureItem {
  ...componentData
  ...item
  title
  secondaryTitle
  seriesShortTitle
  image
  link
  brandDisplayTitle
  whiteBrandLogo
  colorBrandLogo
  destinationType
  destination
  playlistMachineName
  ariaLabel
}
fragment playlistItem on PlaylistItem {
  ...componentData
  ...item
  brandDisplayTitle
  whiteBrandLogo
  colorBrandLogo
  destination
  destType: destinationType
  playlistMachineName
  externalAdvertiserId
  ariaLabel
}
fragment marketingBandData on MarketingBandData {
  ...componentData
  v4ID
  primaryImage
  compactImage
  link
  seriesShortTitle
  lastModified
  brandDisplayTitle
  whiteBrandLogo
  colorBrandLogo
  ariaLabel
}
fragment seriesItem on SeriesItem {
  ...componentData
  ...item
  seriesName
  shortTitle
  urlAlias
  favoritedOn
  favoriteID
  posterImage
  lightPrimaryColor
  darkPrimaryColor
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  landscapePreview
  portraitPreview
  ariaLabel
}
fragment movieItem on MovieItem {
  ...componentData
  ...item
  urlAlias
  favoritedOn
  favoriteID
  posterImage
  lightPrimaryColor
  darkPrimaryColor
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  landscapePreview
  portraitPreview
  rating
  ariaLabel
}
fragment onAirNowItem on OnAirNowItem {
  ...componentData
  v4ID
  image
  title
  secondaryTitle
  startTime
  endTime
  machineName
  whiteBrandLogo
  brandDisplayTitle
  brandLightPrimary
  brandDarkPrimary
  isNew
  audioDescription
  ratingWithAdvisories
  resourceId
  channelId
  badge
  resourceId
  channelId
  watchTagline
  ariaLabel
  notification {
    ...notification
  }
}
fragment upcomingLiveItem on UpcomingLiveItem {
  instanceID
  v4ID
  machineName
  title
  secondaryTitle
  shortDescription
  liveBadge
  upcomingBadge
  image
  startTime
  endTime
  brandV4ID
  whiteBrandLogo
  brandDisplayTitle
  brandLightPrimary
  brandDarkPrimary
  liveAriaLabel
  upcomingAriaLabel
  upcomingModal {
    ...upcomingModal
  }
  resourceId
  channelId
  streamAccessName
  directToLiveThreshold
  notification {
    ...notification
  }
}
fragment upcomingModalData on UpcomingModalData {
  machineName
  title
  description
  ctaText
  dismissText
  lastMinuteModalLifespan
  countdownDayLabel
  countdownHourLabel
  countdownMinLabel
  startTime
  backgroundImage
  backgroundVideo
  resourceId
  channelId
  streamAccessName
}
fragment brandItem on BrandItem {
  ...componentData
  v4ID
  lastModified
  displayTitle
  machineName
  lightPrimaryColor
  darkPrimaryColor
  colorBrandLogo
  whiteBrandLogo
  horizontalPreview
  staticPreviewImage
  ariaLabel
}
fragment slideItem on SlideItem {
  ...componentData
  ...item
  titleColor
  secondaryTitleColor
  description
  descriptionColor
  compactImage
  videoTitle
  percentViewed
  episodeNumber
  seasonNumber
  seriesShortTitle
  programmingType
  portraitPreview
  landscapePreview
  titleLogo
  brandDisplayTitle
  whiteBrandLogo
  colorBrandLogo
  tuneIn
  rating
  locked
  airDate
  sponsorLogo
  sponsorLogoAltText
  sponsorName
  externalAdvertiserId
  ariaLabel
  playlistBadge
  cta {
    ...ctaLink
    ...smartTile
  }
}
fragment upcomingLiveSlideData on UpcomingLiveSlideData {
  ...componentData
  v4ID
  title
  secondaryTitle
  description
  image
  gradientStart
  gradientEnd
  lastModified
  liveTuneIn
  upcomingTuneIn
  liveBadge
  titleColor
  secondaryTitleColor
  descriptionColor
  compactImage
  landscapePreview
  titleLogo
  brandDisplayTitle
  whiteBrandLogo
  colorBrandLogo
  sponsorLogo
  sponsorLogoAltText
  sponsorName
  startTime
  endTime
  liveAriaLabel
  upcomingAriaLabel
  liveCtaColor
  upcomingCtaColor
  liveCtaText
  upcomingCtaText
  notInPackageCtaText
  resourceId
  channelId
  machineName
  streamAccessName
  directToLiveThreshold
  upcomingModal {
    ...upcomingModal
  }
  notification {
    ...notification
  }
}
fragment guideData on GuideData {
  ...componentData
  slotDuration
  slots
  start
  schedules {
    ...guideSchedule
  }
  streams {
    ...guideStream
  }
}
fragment guideScheduleData on GuideScheduleData {
  ...componentData
  programs {
    ...guideProgram
  }
}
fragment guideStreamData on GuideStreamData {
  ...componentData
  lightPrimaryColor
  darkPrimaryColor
  brandDisplayTitle
  machineName
  secondaryTitle
  colorBrandLogo
  whiteBrandLogo
  resourceId
  channelId
  v4ID
  streamAccessName
  xyFallback
}
fragment guideProgramData on GuideProgramData {
  ...componentData
  v4ID
  startSlot
  endSlot
  slotSpan
  image
  episodeNumber
  seasonNumber
  episodeTitle
  programTitle
  programDescription
  ratingWithAdvisories
  startTime
  endTime
  audioDescription
  isNew
  seriesUrlAlias
  backgroundGradientStart
  backgroundGradientEnd
  resourceId
  channelId
  streamAccessName
  tmsId
  whiteBrandLogo
  machineName
  brandV4ID
  notification {
    ...notification
  }
}
fragment slideList on SlideList {
  ...componentData
  lastModified
  items {
    ...slide
    ...upcomingLiveSlide
  }
}
fragment tileList on TileList {
  ...componentData
  playlistMachineName
  listTitle
  ariaLabel
  listTitleImage
  sponsorLogo
  sponsorName
  sponsorLogoAltText
  lastModified
  items {
    ...videoTile
    ...videoStoryTile
    ...seriesTile
    ...movieTile
    ...brandTile
    ...featureTile
    ...playlistTile
    ...marketingBand
    ...slideTile
    ...smartTile
    ...onAirNowTile
    ...upcomingLiveTile
  }
  moreItems {
    ...lazyShelf
    ...lazyGrid
  }
}
fragment premiumTileList on PremiumTileList {
  ...componentData
  listTitle
  listTitleImage
  sponsorLogo
  sponsorLogoAltText
  sponsorName
  lastModified
  ariaLabel
  items {
    ...premiumTile
  }
}
fragment onAirNowList on OnAirNowList {
  ...componentData
  listTitle
  lastModified
  items {
    ...onAirNowTile
  }
}
fragment shelfList on ShelfList {
  ...componentData
  listTitle
  items {
    ...shelf
  }
}
fragment ldExperiment on LdExperiment {
  name
  bucket
}
fragment lazyComponent on LazyComponent {
  targetComponent
  data {
    ...componentData
    queryName
    queryVariables
    entryField
    path
  }
}
fragment lazyComponentData on LazyComponentData {
  ...componentData
  queryName
  variables {
    ...brandSeriesGroupedByCategoryQueryVariables
    ...endCardQueryVariables
    ...videoSectionQueryVariables
  }
  path
}
fragment brandSeriesGroupedByCategoryQueryVariables on BrandSeriesGroupedByCategoryQueryVariables {
  brand
  userId
  platform
  timeZone
  ld
  profile
  oneApp
  app
}
fragment endCardQueryVariables on EndCardQueryVariables {
  type
  mpxGuid
  userId
  platform
  timeZone
  ld
  profile
  oneApp
  app
}
fragment videoSectionQueryVariables on VideosSectionQueryVariables {
  userId
  platform
  profile
  seriesName
  seasonNumber
  programmingType
  currentMpxGuid
  oneApp
  app
}
fragment videoStoryTile on VideoStoryTile {
  ...component
  data {
    ...videoStoryItem
  }
  analytics {
    series
    brand {
      title
    }
    programmingType
    episodeNumber
    seasonNumber
    mpxGuid
    locked
    duration
  }
}
fragment premiumItem on PremiumItem {
  ...componentData
  v4ID
  description
  descriptionColor
  logo {
    ...image
  }
  logoAltText
  popOutImage {
    ...image
  }
  previewStaticImage {
    ...image
  }
  fallbackImage {
    ...image
  }
  backgroundImage {
    ...image
  }
  mainPreview
  locked
  externalAdvertiserId
  gradientStart
  gradientEnd
  ariaLabel
  cta {
    ...ctaLink
    ...smartTile
  }
}
fragment videoStoryItem on VideoStoryItem {
  ...componentData
  v4ID
  playablePublicUrl
  videoID
  titleLogo
  gradientStart
  gradientEnd
  lightPrimaryColor
  darkPrimaryColor
  colorBrandLogo
  whiteBrandLogo
  brandDisplayTitle
  titleKeyArt
  title
  secondaryTitle
  lastModified
  watched
  ariaLabel
  onClickAriaLabel
  cta {
    ...ctaLink
  }
}
fragment notification on Notification {
  ...component
  data {
    ...componentData
    v4ID
    machineName
    headline
    headlineColor
    message
    messageColor
    logo
    logoAltText
    portraitImage
    landscapeImage
    cta {
      ...ctaLink
    }
    dismissText
  }
  analytics {
    entityTitle
    dismissText
    placement
  }
}
`
