package goinsta

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Search is the object for all searches like Facebook, Location or Tag search.
type Search struct {
	insta *Instagram
}

type SearchFunc interface{}

type Place struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Location *Location `json:"location"`
}

type TopSearchItem struct {
	insta *Instagram

	Position int      `json:"position"`
	User     *User    `json:"user"`
	Hashtag  *Hashtag `json:"hashtag"`
	Place    Place    `json:"place"`
}

type SearchHistory struct {
	Time int64 `json:"client_time"`
	User User  `json:"user"`
}

// newSearch creates new Search structure
func newSearch(insta *Instagram) *Search {
	search := &Search{
		insta: insta,
	}
	return search
}

func (sb *Search) History() (*[]SearchHistory, error) {
	sb.insta.Discover.Next()
	h, err := sb.history()
	if err != nil {
		return nil, err
	}
	if err := sb.NullState(); err != nil {
		sb.insta.warnHandler("Non fatal error while setting search null state", err)
	}
	return h, nil
}

func (sr *TopSearchItem) RegisterClick() error {
	insta := sr.insta

	var entityType string
	var id int64
	if id = sr.User.ID; id != 0 {
		entityType = "user"
	} else if id = sr.Hashtag.ID; id != 0 {
		entityType = "hashtag"
	} else if id = sr.Place.Location.ID; id != 0 {
		entityType = "place"
	}

	err := insta.sendSearchRegisterRequest(
		map[string]string{
			"entity_id":   toString(id),
			"_uuid":       insta.uuid,
			"entity_type": entityType,
		},
	)
	return err
}

func (insta *Instagram) sendSearchRegisterRequest(query map[string]string) error {
	_, _, err := insta.sendRequest(&reqOptions{
		Endpoint: urlSearchRegisterClick,
		IsPost:   true,
		Query:    query,
	})
	return err
}

func (search *Search) NullState() error {
	_, _, err := search.insta.sendRequest(&reqOptions{
		Endpoint: urlSearchNullState,
		Query:    map[string]string{"type": "blended"},
	})
	return err
}

func (search *Search) history() (*[]SearchHistory, error) {
	body, err := search.insta.sendSimpleRequest(urlSearchRecent)
	if err != nil {
		return nil, err
	}
	s := struct {
		Recent []SearchHistory `json:"recent"`
		Status string          `json:"status"`
	}{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	for _, i := range s.Recent {
		i.User.insta = search.insta
	}
	return &s.Recent, nil
}


// Profiles allows user function interactions
type Profiles struct {
	insta *Instagram
}

func newProfiles(insta *Instagram) *Profiles {
	profiles := &Profiles{
		insta: insta,
	}
	return profiles
}

// ByName return a *User structure parsed by username.
//   not simply call this endpoint. It is better to use insta.Search(user),
//   or insta.Searchbar.SearchUser(user).
//
func (prof *Profiles) ByName(name string) (*User, error) {
	body, err := prof.insta.sendSimpleRequest(urlUserByName, name)
	if err == nil {
		resp := userResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			user := &resp.User
			user.insta = prof.insta
			return user, err
		}
	}
	return nil, err
}

// ByID returns a *User structure parsed by user id.
func (prof *Profiles) ByID(id_ interface{}) (*User, error) {
	var id string
	switch x := id_.(type) {
	case int64:
		id = fmt.Sprintf("%d", x)
	case int:
		id = fmt.Sprintf("%d", x)
	case string:
		id = x
	default:
		return nil, errors.New("Invalid id, please provide a string or int(64)")
	}

	body, _, err := prof.insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlUserByID, id),
		},
	)
	if err == nil {
		resp := userResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			user := &resp.User
			user.insta = prof.insta
			return user, err
		}
	}
	return nil, err
}

// Blocked returns a list of users you have blocked.
func (prof *Profiles) Blocked() ([]BlockedUser, error) {
	body, err := prof.insta.sendSimpleRequest(urlBlockedList)
	if err == nil {
		resp := blockedListResp{}
		err = json.Unmarshal(body, &resp)
		return resp.BlockedList, err
	}
	return nil, err
}
