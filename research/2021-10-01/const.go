package goinsta

import "errors"

const (
	// urls
	baseUrl        = "https://i.instagram.com/"
	instaAPIUrl    = "https://i.instagram.com/api/v1/"
	instaAPIUrlb   = "https://b.i.instagram.com/api/v1/"
	instaAPIUrlv2  = "https://i.instagram.com/api/v2/"
	instaAPIUrlv2b = "https://b.i.instagram.com/api/v2/"
	// header values
	bloksVerID         = "927f06374b80864ae6a0b04757048065714dc50ff15d2b8b3de8d0b6de961649"
	fbAnalytics        = "567067343352427"
	igCapabilities     = "3brTvx0="
	connType           = "WIFI"
	instaSigKeyVersion = "4"
	locale             = "en_US"
	appVersion         = "195.0.0.31.123"
	appVersionCode     = "302733750"
	// Other
	software = "Android RP1A.200720.012.G975FXXSBFUF3"
	hmacKey  = "iN4$aGr0m"
)

var (
	defaultHeaderOptions = map[string]string{
		"X-Ig-Www-Claim": "0",
	}
	omitAPIHeadersExclude = []string{
		"X-Ig-Bandwidth-Speed-Kbps",
		"Ig-U-Shbts",
		"X-Ig-Mapped-Locale",
		"X-Ig-Family-Device-Id",
		"X-Ig-Android-Id",
		"X-Ig-Timezone-Offset",
		"X-Ig-Device-Locale",
		"X-Ig-Device-Id",
		"Ig-Intended-User-Id",
		"X-Ig-App-Locale",
		"X-Bloks-Is-Layout-Rtl",
		"X-Pigeon-Rawclienttime",
		"X-Bloks-Version-Id",
		"X-Ig-Bandwidth-Totalbytes-B",
		"X-Ig-Bandwidth-Totaltime-Ms",
		"X-Ig-App-Startup-Country",
		"X-Ig-Www-Claim",
		"X-Bloks-Is-Panorama-Enabled",
	}
	// Default Device
	GalaxyS10 = Device{
		Manufacturer:     "samsung",
		Model:            "SM-G975F",
		CodeName:         "beyond2",
		AndroidVersion:   30,
		AndroidRelease:   11,
		ScreenDpi:        "560dpi",
		ScreenResolution: "1440x2898",
		Chipset:          "exynos9820",
	}
	G6 = Device{
		Manufacturer:     "LGE/lge",
		Model:            "LG-H870DS",
		CodeName:         "lucye",
		AndroidVersion:   28,
		AndroidRelease:   9,
		ScreenDpi:        "560dpi",
		ScreenResolution: "1440x2698",
		Chipset:          "lucye",
	}
	timeOffset = getTimeOffset()
)

type muteOption string

const (
	MuteAll   muteOption = "all"
	MuteStory muteOption = "reel"
	MutePosts muteOption = "post"
)

// Endpoints (with format vars)
const (
	// Login
	urlMsisdnHeader               = "accounts/read_msisdn_header/"
	urlGetPrefill                 = "accounts/get_prefill_candidates/"
	urlContactPrefill             = "accounts/contact_point_prefill/"
	urlGetAccFamily               = "multiple_accounts/get_account_family/"
	urlZrToken                    = "zr/token/result/"
	urlLogin                      = "accounts/login/"
	urlLogout                     = "accounts/logout/"
	urlAutoComplete               = "friendships/autocomplete_user_list/"
	urlQeSync                     = "qe/sync/"
	urlSync                       = "launcher/sync/"
	urlLogAttribution             = "attribution/log_attribution/"
	urlMegaphoneLog               = "megaphone/log/"
	urlExpose                     = "qe/expose/"
	urlGetNdxSteps                = "devices/ndx/api/async_get_ndx_ig_steps/"
	urlBanyan                     = "banyan/banyan/"
	urlCooldowns                  = "qp/get_cooldowns/"
	urlFetchConfig                = "loom/fetch_config/"
	urlBootstrapUserScores        = "scores/bootstrap/users/"
	urlStoreClientPushPermissions = "notifications/store_client_push_permissions/"
	urlProcessContactPointSignals = "accounts/process_contact_point_signals/"
	// Account
	urlCurrentUser      = "accounts/current_user/"
	urlChangePass       = "accounts/change_password/"
	urlSetPrivate       = "accounts/set_private/"
	urlSetPublic        = "accounts/set_public/"
	urlRemoveProfPic    = "accounts/remove_profile_picture/"
	urlChangeProfPic    = "accounts/change_profile_picture/"
	urlFeedSaved        = "feed/saved/all/"
	urlFeedSavedPosts   = "feed/saved/posts/"
	urlFeedSavedIGTV    = "feed/saved/igtv/"
	urlEditProfile      = "accounts/edit_profile/"
	urlFeedLiked        = "feed/liked/"
	urlConsent          = "consent/existing_user_flow/"
	urlNotifBadge       = "notifications/badge/"
	urlFeaturedAccounts = "multiple_accounts/get_featured_accounts/"
)

// Errors
var (
	RespErr2FA = "two_factor_required"

	// Account & Login Errors
	ErrBadPassword     = errors.New("Password is incorrect")
	ErrTooManyRequests = errors.New("Too many requests, please wait a few minutes before you try again")

	// Upload Errors
	ErrInvalidFormat      = errors.New("Invalid file type, please use one of jpeg, jpg, mp4")
	ErrCarouselType       = errors.New("Invalid file type, please use a jpeg or jpg image")
	ErrCarouselMediaLimit = errors.New("Carousel media limit of 10 exceeded")
	ErrStoryBadMediaType  = errors.New("When uploading multiple items to your story at once, all have to be mp4")
	ErrStoryMediaTooLong  = errors.New("Story media must not exceed 15 seconds per item")

	// Search Errors
	ErrSearchUserNotFound = errors.New("User not found in search result")

	// IGTV
	ErrIGTVNoSeries = errors.New(
		"User has no IGTV series, unable to fetch. If you think this was a mistake please update the user",
	)

	// Feed Errors
	ErrInvalidTab   = errors.New("Invalid tab, please select top or recent")
	ErrNoMore       = errors.New("No more posts availible, page end has been reached")
	ErrNotHighlight = errors.New("Unable to sync, Reel is not of type highlight")
	ErrMediaDeleted = errors.New("Sorry, this media has been deleted")

	// Inbox
	ErrConvNotPending = errors.New("Unable to perform action, conversation is not pending")

	// Misc
	ErrByteIndexNotFound = errors.New("Failed to index byte slice, delim not found")
	ErrNoMedia           = errors.New("Failed to download, no media found")
	ErrInstaNotDefined   = errors.New(
		"Insta has not been defined, this is most likely a bug in the code. Please backtrack which call this error came from, and open an issue detailing exactly how you got to this error.",
	)
	ErrNoValidLogin = errors.New("No valid login found")
)
