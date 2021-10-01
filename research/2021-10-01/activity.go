package goinsta

import (
   "bytes"
   "crypto/hmac"
   "crypto/md5"
   "crypto/sha256"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "io"
   //"math/rand"
   "time"
   "crypto/rand"
)

// Activity is the recent activity menu.
//
// See example: examples/activity/recent.go
type Activity struct {
	insta *Instagram
	err   error

	// Ad is every column of Activity section
	Ad struct {
		Items []struct {
			// User            User          `json:"user"`
			Algorithm       string        `json:"algorithm"`
			SocialContext   string        `json:"social_context"`
			Icon            string        `json:"icon"`
			Caption         string        `json:"caption"`
			MediaIds        []interface{} `json:"media_ids"`
			ThumbnailUrls   []interface{} `json:"thumbnail_urls"`
			LargeUrls       []interface{} `json:"large_urls"`
			MediaInfos      []interface{} `json:"media_infos"`
			Value           float64       `json:"value"`
			IsNewSuggestion bool          `json:"is_new_suggestion"`
		} `json:"items"`
		MoreAvailable bool `json:"more_available"`
	} `json:"aymf"`
	Counts struct {
		Campaign      int `json:"campaign_notification"`
		CommentLikes  int `json:"comment_likes"`
		Comments      int `json:"comments"`
		Fundraiser    int `json:"fundraiser"`
		Likes         int `json:"likes"`
		NewPosts      int `json:"new_posts"`
		PhotosOfYou   int `json:"photos_of_you"`
		Relationships int `json:"relationships"`
		Requests      int `json:"requests"`
		Shopping      int `json:"shopping_notification"`
		UserTags      int `json:"usertags"`
	} `json:"counts"`
	FriendRequestStories []interface{} `json:"friend_request_stories"`
	NewStories           []RecentItems `json:"new_stories"`
	OldStories           []RecentItems `json:"old_stories"`
	ContinuationToken    int64         `json:"continuation_token"`
	Subscription         interface{}   `json:"subscription"`
	NextID               string        `json:"next_max_id"`
	LastChecked          float64       `json:"last_checked"`
	FirstRecTs           float64       `json:"pagination_first_record_timestamp"`

	Status string `json:"status"`
}

// TODO: Needs to be updated
type RecentItems struct {
	Type      int `json:"type"`
	StoryType int `json:"story_type"`
	Args      struct {
		Text  string `json:"text"`
		Links []struct {
			Start int         `json:"start"`
			End   int         `json:"end"`
			Type  string      `json:"type"`
			ID    interface{} `json:"id"`
		} `json:"links"`
		InlineFollow struct {
			UserInfo        User `json:"user_info"`
			Following       bool `json:"following"`
			OutgoingRequest bool `json:"outgoing_request"`
		} `json:"inline_follow"`
		Actions         []string `json:"actions"`
		ProfileID       int64    `json:"profile_id"`
		ProfileImage    string   `json:"profile_image"`
		Timestamp       float64  `json:"timestamp"`
		Tuuid           string   `json:"tuuid"`
		Clicked         bool     `json:"clicked"`
		ProfileName     string   `json:"profile_name"`
		LatestReelMedia int64    `json:"latest_reel_media"`
	} `json:"args"`
	Counts struct{} `json:"counts"`
	Pk     string   `json:"pk"`
}

func (act *Activity) Error() error {
	return act.err
}

// Next function allows pagination over notifications.
//
// See example: examples/activity/recent.go
func (act *Activity) Next() bool {
	if act.err != nil {
		return false
	}

	query := map[string]string{
		"mark_as_seen":    "false",
		"timezone_offset": timeOffset,
	}
	if act.NextID != "" {
		query["max_id"] = act.NextID
		query["last_checked"] = fmt.Sprintf("%f", act.LastChecked)
		query["pagination_first_record_timestamp"] = fmt.Sprintf("%f", act.FirstRecTs)
	}

	insta := act.insta
	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlActivityRecent,
			Query:    query,
			IsPost:   false,
		},
	)
	if err == nil {
		act2 := Activity{}
		err = json.Unmarshal(body, &act2)
		if err == nil {
			*act = act2
			act.insta = insta
			if len(act.NewStories) == 0 || act.NextID == "" {
				act.err = ErrNoMore
			}
			return true
		}
	}
	act.err = err
	return false
}

func newActivity(insta *Instagram) *Activity {
	act := &Activity{
		insta: insta,
	}
	return act
}


const (
	volatileSeed = "12345"
)

func generateMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateHMAC(text, key string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateDeviceID(seed string) string {
	hash := generateMD5Hash(seed + volatileSeed)
	return "android-" + hash[:16]
}

func generateUserBreadcrumb(text string) string {
	ts := time.Now().Unix()
	d := fmt.Sprintf("%d %d %d %d%d",
		len(text), 0, random(3000, 10000), ts, random(100, 999))
	hmac := base64.StdEncoding.EncodeToString([]byte(generateHMAC(d, hmacKey)))
	enc := base64.StdEncoding.EncodeToString([]byte(d))
	return hmac + "\n" + enc + "\n"
}

func generateSignature(d interface{}, extra ...map[string]string) map[string]string {
	var data string
	switch x := d.(type) {
	case []byte:
		data = string(x)
	case string:
		data = x
	}
	r := map[string]string{
		"signed_body": "SIGNATURE." + data,
	}
	for _, e := range extra {
		for k, v := range e {
			r[k] = v
		}
	}

	return r
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func generateUUID() string {
	uuid, err := newUUID()
	if err != nil {
		return "cb479ee7-a50d-49e7-8b7b-60cc1a105e22" // default value when error occurred
	}
	return uuid
}


func readFile(f io.Reader) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(f)
	return buf, err
}
