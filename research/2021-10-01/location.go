package goinsta

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LocationInstance struct {
	insta *Instagram
}

func newLocation(insta *Instagram) *LocationInstance {
	return &LocationInstance{insta: insta}
}

type LayoutSection struct {
	LayoutType    string `json:"layout_type"`
	LayoutContent struct {
		Medias []struct {
			Media Item `json:"media"`
		} `json:"medias"`
	} `json:"layout_content"`
	FeedType        string `json:"feed_type"`
	ExploreItemInfo struct {
		NumColumns      int     `json:"num_columns"`
		TotalNumColumns int     `json:"total_num_columns"`
		AspectRatio     float64 `json:"aspect_ratio"`
		Autoplay        bool    `json:"autoplay"`
	} `json:"explore_item_info"`
}

type Section struct {
	Sections      []LayoutSection `json:"sections"`
	MoreAvailable bool            `json:"more_available"`
	NextPage      int             `json:"next_page"`
	NextMediaIds  []int64         `json:"next_media_ids"`
	NextID        string          `json:"next_max_id"`
	Status        string          `json:"status"`
}

func (l *LocationInstance) Feeds(locationID int64) (*Section, error) {
	// TODO: use pagination for location feeds.
	insta := l.insta
	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlFeedLocations, locationID),
			IsPost:   true,
			Query: map[string]string{
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
				"_uid":           toString(insta.Account.ID),
				"_uuid":          insta.uuid,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	section := &Section{}
	err = json.Unmarshal(body, section)
	return section, err
}

func (l *Location) Feed() (*Section, error) {
	// TODO: use pagination for location feeds.
	insta := l.insta
	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlFeedLocations, l.ID),
			IsPost:   true,
			Query: map[string]string{
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
				"_uid":           toString(insta.Account.ID),
				"_uuid":          insta.uuid,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	section := &Section{}
	err = json.Unmarshal(body, section)
	return section, err
}


type Contacts struct {
	insta *Instagram
}

type Contact struct {
	Numbers []string `json:"phone_numbers"`
	Emails  []string `json:"email_addresses"`
	Name    string   `json:"first_name"`
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

func newContacts(insta *Instagram) *Contacts {
	return &Contacts{insta: insta}
}

func (c *Contacts) SyncContacts(contacts *[]Contact) (*SyncAnswer, error) {
	byteContacts, err := json.Marshal(contacts)
	if err != nil {
		return nil, err
	}

	syncContacts := &reqOptions{
		Endpoint: `address_book/link/`,
		IsPost:   true,
		Gzip:     true,
		Query: map[string]string{
			"phone_id":  c.insta.pid,
			"module":    "find_friends_contacts",
			"source":    "user_setting",
			"device_id": c.insta.uuid,
			"_uuid":     c.insta.uuid,
			"contacts":  string(byteContacts),
		},
	}

	body, _, err := c.insta.sendRequest(syncContacts)
	if err != nil {
		return nil, err
	}

	answ := &SyncAnswer{}
	json.Unmarshal(body, answ)
	return answ, nil
}

func (c *Contacts) UnlinkContacts() error {
	unlinkBody := &reqOptions{
		Endpoint: "address_book/unlink/",
		IsPost:   true,
		Query: map[string]string{
			"phone_id":       c.insta.pid,
			"device_id":      c.insta.uuid,
			"_uuid":          c.insta.uuid,
			"user_initiated": "true",
		},
	}

	_, _, err := c.insta.sendRequest(unlinkBody)
	if err != nil {
		return err
	}
	return nil
}


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
