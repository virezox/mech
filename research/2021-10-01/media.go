package goinsta

import (
	"encoding/json"
	"fmt"
	"io"
	neturl "net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

// Media interface defines methods for both StoryMedia and FeedMedia.
type Media interface {
	// Next allows pagination
	Next(...interface{}) bool
	// Error returns error (in case it have been occurred)
	Error() error
	// ID returns media id
	GetNextID() string
	// Delete removes media
	Delete() error
	getInsta() *Instagram
}

// Item represents media items
//
// All Item has Images or Videos objects which contains the url(s).
// You can use Download function to get the best quality Image or Video from Item.
type Item struct {
	insta    *Instagram
	media    Media
	module   string
	// Post Info
	TakenAt           int64  `json:"taken_at"`
	Pk                int64  `json:"pk"`
	ID                string `json:"id"`
	Index             int    // position in feed
	CommentsDisabled  bool   `json:"comments_disabled"`
	DeviceTimestamp   int64  `json:"device_timestamp"`
	FacepileTopLikers []struct {
		FollowFrictionType float64 `json:"follow_friction_type"`
		FullNeme           string  `json:"ful_name"`
		IsPrivate          bool    `json:"is_private"`
		IsVerified         bool    `json:"is_verified"`
		Pk                 float64 `json:"pk"`
		ProfilePicID       string  `json:"profile_pic_id"`
		ProfilePicURL      string  `json:"profile_pic_url"`
		Username           string  `json:"username"`
	} `json:"facepile_top_likers"`
	MediaType             int     `json:"media_type"`
	Code                  string  `json:"code"`
	ClientCacheKey        string  `json:"client_cache_key"`
	FilterType            int     `json:"filter_type"`
	User                  User    `json:"user"`
	CanReply              bool    `json:"can_reply"`
	CanReshare            bool    `json:"can_reshare"` // Used for stories
	CanViewerReshare      bool    `json:"can_viewer_reshare"`
	Caption               Caption `json:"caption"`
	CaptionIsEdited       bool    `json:"caption_is_edited"`
	LikeViewCountDisabled bool    `json:"like_and_view_counts_disabled"`
	FundraiserTag         struct {
		HasStandaloneFundraiser bool `json:"has_standalone_fundraiser"`
	} `json:"fundraiser_tag"`
	IsSeen                       bool   `json:"is_seen"`
	InventorySource              string `json:"inventory_source"`
	ProductType                  string `json:"product_type"`
	Likes                        int    `json:"like_count"`
	HasLiked                     bool   `json:"has_liked"`
	NearlyCompleteCopyRightMatch bool   `json:"nearly_complete_copyright_match"`
	// Toplikers can be `string` or `[]string`.
	// Use TopLikers function instead of getting it directly.
	Toplikers                    interface{} `json:"top_likers"`
	Likers                       []*User     `json:"likers"`
	CommentLikesEnabled          bool        `json:"comment_likes_enabled"`
	CommentThreadingEnabled      bool        `json:"comment_threading_enabled"`
	HasMoreComments              bool        `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int         `json:"max_num_visible_preview_comments"`
	// Previewcomments can be `string` or `[]string` or `[]Comment`.
	// Use PreviewComments function instead of getting it directly.
	Previewcomments interface{} `json:"preview_comments,omitempty"`
	CommentCount    int         `json:"comment_count"`
	PhotoOfYou      bool        `json:"photo_of_you"`
	// Tags are tagged people in photo
	Tags struct {
		In []Tag `json:"in"`
	} `json:"usertags,omitempty"`
	FbUserTags           Tag    `json:"fb_user_tags"`
	CanViewerSave        bool   `json:"can_viewer_save"`
	OrganicTrackingToken string `json:"organic_tracking_token"`
	// Images contains URL images in different versions.
	// Version = quality.
	Images          Images   `json:"image_versions2,omitempty"`
	OriginalWidth   int      `json:"original_width,omitempty"`
	OriginalHeight  int      `json:"original_height,omitempty"`
	ImportedTakenAt int64    `json:"imported_taken_at,omitempty"`
	Location        Location `json:"location,omitempty"`
	Lat             float64  `json:"lat,omitempty"`
	Lng             float64  `json:"lng,omitempty"`

	// Carousel
	CarouselParentID string `json:"carousel_parent_id"`
	CarouselMedia    []Item `json:"carousel_media,omitempty"`

	// Live
	IsPostLive bool `json:"is_post_live"`

	// Videos
	Videos            []Video `json:"video_versions,omitempty"`
	VideoCodec        string  `json:"video_codec"`
	HasAudio          bool    `json:"has_audio,omitempty"`
	VideoDuration     float64 `json:"video_duration,omitempty"`
	ViewCount         float64 `json:"view_count,omitempty"`
	IsDashEligible    int     `json:"is_dash_eligible,omitempty"`
	IsUnifiedVideo    bool    `json:"is_unified_video"`
	VideoDashManifest string  `json:"video_dash_manifest,omitempty"`
	NumberOfQualities int     `json:"number_of_qualities,omitempty"`

	// IGTV
	Title                    string `json:"title"`
	IGTVExistsInViewerSeries bool   `json:"igtv_exists_in_viewer_series"`
	IGTVSeriesInfo           struct {
		HasCoverPhoto bool `json:"has_cover_photo"`
		ID            int64
		NumEpisodes   int    `json:"num_episodes"`
		Title         string `json:"title"`
	} `json:"igtv_series_info"`
	IGTVAdsInfo struct {
		AdsToggledOn            bool `json:"ads_toggled_on"`
		ElegibleForInsertingAds bool `json:"is_video_elegible_for_inserting_ads"`
	} `json:"igtv_ads_info"`

	// Ads
	IsCommercial        bool   `json:"is_commercial"`
	IsPaidPartnership   bool   `json:"is_paid_partnership"`
	CommercialityStatus string `json:"commerciality_status"`
	AdLink              string `json:"link"`
	AdLinkText          string `json:"link_text"`
	AdLinkHint          string `json:"link_hint_text"`
	AdTitle             string `json:"overlay_title"`
	AdSubtitle          string `json:"overlay_subtitle"`
	AdText              string `json:"overlay_text"`
	AdAction            string `json:"ad_action"`
	AdHeaderStyle       int    `json:"ad_header_style"`
	AdLinkType          int    `json:"ad_link_type"`
	AdMetadata          []struct {
		Type  int         `json:"type"`
		Value interface{} `json:"value"`
	} `json:"ad_metadata"`
	AndroidLinks []struct {
		AndroidClass      string `json:"androidClass"`
		CallToActionTitle string `json:"callToActionTitle"`
		DeeplinkUri       string `json:"deeplinkUri"`
		LinkType          int    `json:"linkType"`
		Package           string `json:"package"`
		WebUri            string `json:"webUri"`
	} `json:"android_links"`

	// Only for stories
	StoryEvents              []interface{}      `json:"story_events"`
	StoryHashtags            []interface{}      `json:"story_hashtags"`
	StoryPolls               []interface{}      `json:"story_polls"`
	StoryFeedMedia           []interface{}      `json:"story_feed_media"`
	StorySoundOn             []interface{}      `json:"story_sound_on"`
	CreativeConfig           interface{}        `json:"creative_config"`
	StoryLocations           []interface{}      `json:"story_locations"`
	StorySliders             []interface{}      `json:"story_sliders"`
	StoryQuestions           []interface{}      `json:"story_questions"`
	StoryProductItems        []interface{}      `json:"story_product_items"`
	StoryCTA                 []StoryCTA         `json:"story_cta"`
	IntegrityReviewDecision  string             `json:"integrity_review_decision"`
	IsReelMedia              bool               `json:"is_reel_media"`
	ProfileGridControl       bool               `json:"profile_grid_control_enabled"`
	ReelMentions             []StoryReelMention `json:"reel_mentions"`
	ExpiringAt               int64              `json:"expiring_at"`
	CanSendCustomEmojis      bool               `json:"can_send_custom_emojis"`
	SupportsReelReactions    bool               `json:"supports_reel_reactions"`
	ShowOneTapFbShareTooltip bool               `json:"show_one_tap_fb_share_tooltip"`
	HasSharedToFb            int64              `json:"has_shared_to_fb"`
	Mentions                 []Mentions
	Audience                 string `json:"audience,omitempty"`
	StoryMusicStickers       []struct {
		X              float64 `json:"x"`
		Y              float64 `json:"y"`
		Z              int     `json:"z"`
		Width          float64 `json:"width"`
		Height         float64 `json:"height"`
		Rotation       float64 `json:"rotation"`
		IsPinned       int     `json:"is_pinned"`
		IsHidden       int     `json:"is_hidden"`
		IsSticker      int     `json:"is_sticker"`
		MusicAssetInfo struct {
			ID                       string `json:"id"`
			Title                    string `json:"title"`
			Subtitle                 string `json:"subtitle"`
			DisplayArtist            string `json:"display_artist"`
			CoverArtworkURI          string `json:"cover_artwork_uri"`
			CoverArtworkThumbnailURI string `json:"cover_artwork_thumbnail_uri"`
			ProgressiveDownloadURL   string `json:"progressive_download_url"`
			HighlightStartTimesInMs  []int  `json:"highlight_start_times_in_ms"`
			IsExplicit               bool   `json:"is_explicit"`
			DashManifest             string `json:"dash_manifest"`
			HasLyrics                bool   `json:"has_lyrics"`
			AudioAssetID             string `json:"audio_asset_id"`
			IgArtist                 struct {
				Pk            int    `json:"pk"`
				Username      string `json:"username"`
				FullName      string `json:"full_name"`
				IsPrivate     bool   `json:"is_private"`
				ProfilePicURL string `json:"profile_pic_url"`
				ProfilePicID  string `json:"profile_pic_id"`
				IsVerified    bool   `json:"is_verified"`
			} `json:"ig_artist"`
			PlaceholderProfilePicURL string `json:"placeholder_profile_pic_url"`
			ShouldMuteAudio          bool   `json:"should_mute_audio"`
			ShouldMuteAudioReason    string `json:"should_mute_audio_reason"`
			OverlapDurationInMs      int    `json:"overlap_duration_in_ms"`
			AudioAssetStartTimeInMs  int    `json:"audio_asset_start_time_in_ms"`
		} `json:"music_asset_info"`
	} `json:"story_music_stickers,omitempty"`
}

func (item *Item) comment(text string) error {
	insta := item.insta
	query := map[string]string{
		// "feed_position":           "",
		"container_module":        "feed_timeline",
		"user_breadcrumb":         generateUserBreadcrumb(text),
		"nav_chain":               "",
		"_uid":                    toString(insta.Account.ID),
		"_uuid":                   insta.uuid,
		"idempotence_token":       generateUUID(),
		"radio_type":              "wifi-none",
		"is_carousel_bumped_post": "false", // not sure when this would be true
		"comment_text":            text,
	}
	if item.module != "" {
		query["container_module"] = item.module
	}
	if item.IsCommercial {
		query["delivery_class"] = "ad"
	} else {
		query["delivery_class"] = "organic"
	}
	if item.InventorySource != "" {
		query["inventory_source"] = item.InventorySource
	}
	if len(item.CarouselMedia) > 0 || item.CarouselParentID != "" {
		query["carousel_index"] = "0"
	}
	b, err := json.Marshal(query)
	if err != nil {
		return err
	}

	// ignoring response
	_, _, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlCommentAdd, item.Pk),
			Query:    map[string]string{"signed_body": "SIGNATURE." + string(b)},
			IsPost:   true,
		},
	)
	return err
}

func (item *Item) CommentCheckOffensive(comment string) (*CommentOffensive, error) {
	insta := item.insta
	data, err := json.Marshal(map[string]string{
		"media_id":           item.ID,
		"_uid":               toString(insta.Account.ID),
		"comment_session_id": generateUUID(),
		"_uuid":              insta.uuid,
		"comment_text":       comment,
	})
	if err != nil {
		return nil, err
	}
	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlCommentOffensive,
			IsPost:   true,
			Query:    generateSignature(data),
		},
	)
	if err != nil {
		return nil, err
	}
	r := &CommentOffensive{}
	err = json.Unmarshal(body, r)
	return r, err
}

// MediaToString returns Item.MediaType as string.
func (item *Item) MediaToString() string {
	return MediaToString(item.MediaType)
}

func MediaToString(t int) string {
	switch t {
	case 1:
		return "photo"
	case 2:
		return "video"
	case 6:
		return "ad_map"
	case 7:
		return "live"
	case 8:
		return "carousel"
	case 9:
		return "live_replay"
	case 10:
		return "collection"
	case 11:
		return "audio"
	case 12:
		return "showreel_native"
	case 13:
		return "guide_facade"
	}
	return ""
}

func getname(name string) string {
	nname := name
	i := 1
	for {
		ext := path.Ext(name)

		_, err := os.Stat(name)
		if err != nil {
			break
		}
		if ext != "" {
			nname = strings.Replace(nname, ext, "", -1)
		}
		name = fmt.Sprintf("%s.%d%s", nname, i, ext)
		i++
	}
	return name
}

type bestMedia struct {
	w, h int
	url  string
}

// GetBest returns url to best quality image or video.
//
// Arguments can be []Video or []Candidate
func GetBest(obj interface{}) string {
	m := bestMedia{}

	switch t := obj.(type) {
	// getting best video
	case []Video:
		for _, video := range t {
			if m.w < video.Width && video.Height > m.h && video.URL != "" {
				m.w = video.Width
				m.h = video.Height
				m.url = video.URL
			}
		}
		// getting best image
	case []Candidate:
		for _, image := range t {
			if m.w < image.Width && image.Height > m.h && image.URL != "" {
				m.w = image.Width
				m.h = image.Height
				m.url = image.URL
			}
		}
	}
	return m.url
}

var rxpTags = regexp.MustCompile(`#\w+`)

// Download downloads media item (video or image) with the best quality.
//
// Input parameters are folder and filename. If filename is "" will be saved with
// the default value name.
//
// If file exists it will be saved
// This function makes folder automatically
//
// See example: examples/media/itemDownload.go
func (item *Item) Download(folder, name string) (err error) {
	insta := item.insta
	os.MkdirAll(folder, 0o777)

	switch item.MediaType {
	case 1:
		return insta.download(folder, name, item.Images.Versions)
	case 2:
		return insta.download(folder, name, item.Videos)
	case 8:
		return item.downloadCarousel(folder, name)
	}

	insta.warnHandler(
		fmt.Sprintf(
			"Unable to download %s media (media type %d), this has not been implemented",
			item.MediaToString(),
			item.MediaType,
		),
	)
	return ErrNoMedia
}

func (item *Item) downloadCarousel(folder, name string) error {
	if name == "" {
		name = item.ID
	}
	for i, media := range item.CarouselMedia {
		n := fmt.Sprintf("%s_%d", name, i+1)
		if err := media.Download(folder, n); err != nil {
			return err
		}
	}
	return nil
}

func (insta *Instagram) download(folder, name string, media interface{}) error {
	url := GetBest(media)
	name, err := getDownloadName(url, name)
	if err != nil {
		return err
	}
	dst := path.Join(folder, name)

	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := insta.c.Get(url)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	return err
}

func getDownloadName(url, name string) (string, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}
	ext := path.Ext(u.Path)
	if name == "" {
		name = path.Base(u.Path)
	} else if !strings.HasSuffix(name, ext) {
		name += ext
	}
	name = getname(name)
	return name, nil
}

// StoryIsCloseFriends returns a bool
// If the returned value is true the story was published only for close friends
func (item *Item) StoryIsCloseFriends() bool {
	return item.Audience == "besties"
}
