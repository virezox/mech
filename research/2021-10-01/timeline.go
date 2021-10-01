package goinsta

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
