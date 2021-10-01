package goinsta

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrAllSaved = errors.New("Unable to call function for collection all posts")

// MediaItem defines a item media for the
// SavedMedia struct
type MediaItem struct {
	Media Item `json:"media"`
}

// Save saves media item.
//
// You can get saved media using Account.Saved()
func (item *Item) Save() error {
	return item.changeSave(urlMediaSave)
}

// Unsave unsaves media item.
func (item *Item) Unsave() error {
	return item.changeSave(urlMediaUnsave)
}

func (item *Item) changeSave(endpoint string) error {
	insta := item.insta
	query := map[string]string{
		"module_name":     "feed_timeline",
		"client_position": toString(item.Index),
		"nav_chain":       "",
		"_uid":            toString(insta.Account.ID),
		"_uuid":           insta.uuid,
		"radio_type":      "wifi-none",
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
	data, err := json.Marshal(query)
	if err != nil {
		return err
	}

	_, _, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(endpoint, item.ID),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

type reelShare struct {
	Text        string `json:"text"`
	IsPersisted bool   `json:"is_reel_persisted"`
	OwnderID    int64  `json:"reel_owner_id"`
	Type        string `json:"type"`
	ReelType    string `json:"reel_type"`
	Media       Item   `json:"media"`
}

type actionLog struct {
	Description string `json:"description"`
}

type Place struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Location *Location `json:"location"`
}

type mediaItem struct {
	Item *Item `json:"media"`
}

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

type feedCache struct {
	Items []struct {
		MediaOrAd *Item `json:"media_or_ad"`
		EndOfFeed struct {
			Pause    bool   `json:"pause"`
			Title    string `json:"title"`
			Subtitle string `json:"subtitle"`
		} `json:"end_of_feed_demarcator"`
	} `json:"feed_items"`

	MoreAvailable               bool    `json:"more_available"`
	NextID                      string  `json:"next_max_id"`
	NumResults                  float64 `json:"num_results"`
	PullToRefreshWindowMs       float64 `json:"pull_to_refresh_window_ms"`
	RequestID                   string  `json:"request_id"`
	SessionID                   string  `json:"session_id"`
	ViewStateVersion            string  `json:"view_state_version"`
	AutoLoadMore                bool    `json:"auto_load_more_enabled"`
	IsDirectV2Enabled           bool    `json:"is_direct_v2_enabled"`
	ClientFeedChangelistApplied bool    `json:"client_feed_changelist_applied"`
	PreloadDistance             float64 `json:"preload_distance"`
	Status                      string  `json:"status"`
	FeedPillText                string  `json:"feed_pill_text"`
	StartupPrefetchConfigs      struct {
		Explore struct {
			ContainerModule          string `json:"containermodule"`
			ShouldPrefetch           bool   `json:"should_prefetch"`
			ShouldPrefetchThumbnails bool   `json:"should_prefetch_thumbnails"`
		} `json:"explore"`
	} `json:"startup_prefetch_configs"`
	UseAggressiveFirstTailLoad bool    `json:"use_aggressive_first_tail_load"`
	HideLikeAndViewCounts      float64 `json:"hide_like_and_view_counts"`
}
