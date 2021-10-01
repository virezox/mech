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
