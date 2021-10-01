package goinsta

import (
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
)

type accountResp struct {
	Status  string  `json:"status"`
	Account Account `json:"logged_in_user"`

	ErrorType         string         `json:"error_type"`
	Message           string         `json:"message"`
	TwoFactorRequired bool           `json:"two_factor_required"`
	TwoFactorInfo     *TwoFactorInfo `json:"two_factor_info"`

	PhoneVerificationSettings phoneVerificationSettings `json:"phone_verification_settings"`
	Challenge                 *Challenge                `json:"challenge"`
}

// Account is personal account object
//
// See examples: examples/account/*
type Account struct {
	insta *Instagram
	ID                         int64        `json:"pk"`
	Username                   string       `json:"username"`
	FullName                   string       `json:"full_name"`
	Biography                  string       `json:"biography"`
	ProfilePicURL              string       `json:"profile_pic_url"`
	Email                      string       `json:"email"`
	PhoneNumber                string       `json:"phone_number"`
	IsBusiness                 bool         `json:"is_business"`
	Gender                     int          `json:"gender"`
	ProfilePicID               string       `json:"profile_pic_id"`
	CanSeeOrganicInsights      bool         `json:"can_see_organic_insights"`
	ShowInsightsTerms          bool         `json:"show_insights_terms"`
	HasAnonymousProfilePicture bool         `json:"has_anonymous_profile_picture"`
	IsPrivate                  bool         `json:"is_private"`
	IsUnpublished              bool         `json:"is_unpublished"`
	AllowedCommenterType       string       `json:"allowed_commenter_type"`
	IsVerified                 bool         `json:"is_verified"`
	MediaCount                 int          `json:"media_count"`
	FollowerCount              int          `json:"follower_count"`
	FollowingCount             int          `json:"following_count"`
	GeoMediaCount              int          `json:"geo_media_count"`
	ExternalURL                string       `json:"external_url"`
	HasBiographyTranslation    bool         `json:"has_biography_translation"`
	ExternalLynxURL            string       `json:"external_lynx_url"`
	HdProfilePicURLInfo        PicURLInfo   `json:"hd_profile_pic_url_info"`
	HdProfilePicVersions       []PicURLInfo `json:"hd_profile_pic_versions"`
	UsertagsCount              int          `json:"usertags_count"`
	HasChaining                bool         `json:"has_chaining"`
	ReelAutoArchive            string       `json:"reel_auto_archive"`
	PublicEmail                string       `json:"public_email"`
	PublicPhoneNumber          string       `json:"public_phone_number"`
	PublicPhoneCountryCode     string       `json:"public_phone_country_code"`
	ContactPhoneNumber         string       `json:"contact_phone_number"`
	Byline                     string       `json:"byline"`
	SocialContext              string       `json:"social_context,omitempty"`
	SearchSocialContext        string       `json:"search_social_context,omitempty"`
	MutualFollowersCount       float64      `json:"mutual_followers_count"`
	LatestReelMedia            int64        `json:"latest_reel_media,omitempty"`
	CityID                     int64        `json:"city_id"`
	CityName                   string       `json:"city_name"`
	AddressStreet              string       `json:"address_street"`
	DirectMessaging            string       `json:"direct_messaging"`
	Latitude                   float64      `json:"latitude"`
	Longitude                  float64      `json:"longitude"`
	Category                   string       `json:"category"`
	BusinessContactMethod      string       `json:"business_contact_method"`
	IsCallToActionEnabled      bool         `json:"is_call_to_action_enabled"`
	FbPageCallToActionID       string       `json:"fb_page_call_to_action_id"`
	Zip                        string       `json:"zip"`
	AllowContactsSync          bool         `json:"allow_contacts_sync"`
	CanBoostPost               bool         `json:"can_boost_post"`
}

// Sync updates account information
func (account *Account) Sync() error {
	insta := account.insta
	body, _, err := insta.sendRequest(&reqOptions{
		Endpoint: urlCurrentUser,
		Query: map[string]string{
			"edit": "true",
		},
	})
	if err != nil {
		return err
	}

	resp := profResp{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	*account = resp.Account
	account.insta = insta
	return nil
}

type profResp struct {
	Status  string  `json:"status"`
	Account Account `json:"user"`
}

type SyncAnswer struct {
	Users []struct {
		Pk                         int64  `json:"pk"`
		Username                   string `json:"username"`
		FullName                   string `json:"full_name"`
		IsPrivate                  bool   `json:"is_private"`
		ProfilePicURL              string `json:"profile_pic_url"`
		ProfilePicID               string `json:"profile_pic_id"`
		IsVerified                 bool   `json:"is_verified"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		ReelAutoArchive            string `json:"reel_auto_archive"`
		AddressbookName            string `json:"addressbook_name"`
	} `json:"users"`
	Warning string `json:"warning"`
	Status  string `json:"status"`
}

var ErrAllSaved = errors.New("Unable to call function for collection all posts")

type (
	fetchReason string
)

var (
	PULLTOREFRESH fetchReason = "pull_to_refresh"
	COLDSTART     fetchReason = "cold_start_fetch"
	WARMSTART     fetchReason = "warm_start_fetch"
	PAGINATION    fetchReason = "pagination"
	AUTOREFRESH   fetchReason = "auto_refresh" // so far unused
)


// ConfigFile is a structure to store the session information so that can be exported or imported.
type ConfigFile struct {
	ID            int64             `json:"id"`
	User          string            `json:"username"`
	DeviceID      string            `json:"device_id"`
	FamilyID      string            `json:"family_id"`
	UUID          string            `json:"uuid"`
	RankToken     string            `json:"rank_token"`
	Token         string            `json:"token"`
	PhoneID       string            `json:"phone_id"`
	XmidExpiry    int64             `json:"xmid_expiry"`
	HeaderOptions map[string]string `json:"header_options"`
	Cookies       []*http.Cookie    `json:"cookies"`
	Account       *Account          `json:"account"`
	Device        Device            `json:"device"`
}

type Device struct {
	Manufacturer     string `json:"manufacturer"`
	Model            string `json:"model"`
	CodeName         string `json:"code_name"`
	AndroidVersion   int    `json:"android_version"`
	AndroidRelease   int    `json:"android_release"`
	ScreenDpi        string `json:"screen_dpi"`
	ScreenResolution string `json:"screen_resolution"`
	Chipset          string `json:"chipset"`
}

// School is void structure (yet). Whats this even for lol
type School struct{}

// PicURLInfo repre
type PicURLInfo struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

// ErrorN is general instagram error
type ErrorN struct {
	Message   string `json:"message"`
	Endpoint  string `json:"endpoint"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

// Error503 is instagram API error
type Error503 struct {
	Message string
}

func (e Error503) Error() string {
	return e.Message
}

func (e ErrorN) Error() string {
	return fmt.Sprintf(
		"Error while calling %s, status code %s: %s (%s)",
		e.Endpoint, e.Status, e.Message, e.ErrorType,
	)
}

type Cooldowns struct {
	Default int    `json:"default"`
	Global  int    `json:"global"`
	Status  string `json:"status"`
	TTL     int    `json:"ttl"`
	Slots   []struct {
		Cooldown int    `json:"cooldown"`
		Slot     string `json:"slot"`
	} `json:"slots"`
	Surfaces []struct {
		Cooldown int    `json:"cooldown"`
		Slot     string `json:"slot"`
	} `json:"surfaces"`
}
