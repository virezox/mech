package goinsta

import (
   "errors"
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
   urlBanyan                     = "banyan/banyan/"
   urlCooldowns                  = "qp/get_cooldowns/"
   urlFetchConfig                = "loom/fetch_config/"
   urlBootstrapUserScores        = "scores/bootstrap/users/"
   urlStoreClientPushPermissions = "notifications/store_client_push_permissions/"
   urlProcessContactPointSignals = "accounts/process_contact_point_signals/"
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
)

// Account & Login Errors
var ErrBadPassword = errors.New("password is incorrect")

var ErrInstaNotDefined = errors.New(
   "insta has not been defined, this is most likely a bug in the code. Please " +
   "backtrack which call this error came from, and open an issue detailing " +
   "exactly how you got to this error",
)

var ErrAllSaved = errors.New("Unable to call function for collection all posts")
var defaultHeaderOptions = map[string]string{"X-Ig-Www-Claim": "0"}
var timeOffset = getTimeOffset()


var omitAPIHeadersExclude = []string{
   "Ig-Intended-User-Id",
   "Ig-U-Shbts",
   "X-Bloks-Is-Layout-Rtl",
   "X-Bloks-Is-Panorama-Enabled",
   "X-Bloks-Version-Id",
   "X-Ig-Android-Id",
   "X-Ig-App-Locale",
   "X-Ig-App-Startup-Country",
   "X-Ig-Bandwidth-Speed-Kbps",
   "X-Ig-Bandwidth-Totalbytes-B",
   "X-Ig-Bandwidth-Totaltime-Ms",
   "X-Ig-Device-Id",
   "X-Ig-Device-Locale",
   "X-Ig-Family-Device-Id",
   "X-Ig-Mapped-Locale",
   "X-Ig-Timezone-Offset",
   "X-Ig-Www-Claim",
   "X-Pigeon-Rawclienttime",
}

type Account struct {
   ID                         int64        `json:"pk"`
   insta *Instagram
}

// ConfigFile is a structure to store the session information so that can be
// exported or imported.
type ConfigFile struct {
   Account       *Account          `json:"account"`
   Device        Device            `json:"device"`
   DeviceID      string            `json:"device_id"`
   FamilyID      string            `json:"family_id"`
   HeaderOptions map[string]string `json:"header_options"`
   ID            int64             `json:"id"`
   PhoneID       string            `json:"phone_id"`
   RankToken     string            `json:"rank_token"`
   Token         string            `json:"token"`
   UUID          string            `json:"uuid"`
   User          string            `json:"username"`
   XmidExpiry    int64             `json:"xmid_expiry"`
}

type Device struct {
   AndroidRelease   int    `json:"android_release"`
   AndroidVersion   int    `json:"android_version"`
   Chipset          string `json:"chipset"`
   CodeName         string `json:"code_name"`
   Manufacturer     string `json:"manufacturer"`
   Model            string `json:"model"`
   ScreenDpi        string `json:"screen_dpi"`
   ScreenResolution string `json:"screen_resolution"`
}

type accountResp struct {
   Account Account `json:"logged_in_user"`
   ErrorType         string         `json:"error_type"`
   Message           string         `json:"message"`
   Status  string  `json:"status"`
   TwoFactorRequired bool           `json:"two_factor_required"`
}
