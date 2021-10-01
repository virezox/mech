package goinsta

import (
	"encoding/json"
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
