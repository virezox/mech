package goinsta

import (
   "bytes"
   "encoding/base64"
   "encoding/binary"
   "encoding/json"
   "fmt"
   "image"
   "math"
   "strconv"
   "time"
   _ "image/jpeg"
   _ "image/png"
)

func toString(i interface{}) string {
	switch s := i.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatInt(int64(s), 10)
	case uint64:
		return strconv.FormatInt(int64(s), 10)
	case uint32:
		return strconv.FormatInt(int64(s), 10)
	case uint16:
		return strconv.FormatInt(int64(s), 10)
	case uint8:
		return strconv.FormatInt(int64(s), 10)
	case []byte:
		return string(s)
	case error:
		return s.Error()
	}
	return ""
}

// getImageSize return image dimension , types is .jpg and .png
func getImageSize(b []byte) (int, int, error) {
	buf := bytes.NewReader(b)
	image, _, err := image.DecodeConfig(buf)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}

func getVideoInfo(b []byte) (height, width, duration int, err error) {
	keys := []string{"moov", "trak", "stbl", "avc1"}
	height, err = read16(b, keys, 24)
	if err != nil {
		return
	}
	width, err = read16(b, keys, 26)
	if err != nil {
		return
	}

	duration, err = getMP4Duration(b)
	if err != nil {
		return
	}

	return
}

func getMP4Duration(b []byte) (int, error) {
	keys := []string{"moov", "mvhd"}
	timescale, err := read32(b, keys, 12)
	if err != nil {
		return -1, err
	}
	length, err := read32(b, keys, 12+4)
	if err != nil {
		return -1, err
	}

	return int(math.Floor(float64(length) / float64(timescale) * 1000)), nil
}

func getTimeOffset() string {
	_, offset := time.Now().Zone()
	return strconv.Itoa(offset)
}

func jazoest(str string) string {
	b := []byte(str)
	var s int
	for v := range b {
		s += v
	}
	return "2" + strconv.Itoa(s)
}

func createUserAgent(device Device) string {
	// Instagram 195.0.0.31.123 Android (28/9; 560dpi; 1440x2698; LGE/lge; LG-H870DS; lucye; lucye; en_GB; 302733750)
	// Instagram 195.0.0.31.123 Android (28/9; 560dpi; 1440x2872; Genymotion/Android; Samsung Galaxy S10; vbox86p; vbox86; en_US; 302733773)  # version_code: 302733773
	// Instagram 195.0.0.31.123 Android (30/11; 560dpi; 1440x2898; samsung; SM-G975F; beyond2; exynos9820; en_US; 302733750)
	return fmt.Sprintf("Instagram %s Android (%d/%d; %s; %s; %s; %s; %s; %s; %s; %s)",
		appVersion,
		device.AndroidVersion,
		device.AndroidRelease,
		device.ScreenDpi,
		device.ScreenResolution,
		device.Manufacturer,
		device.Model,
		device.CodeName,
		device.Chipset,
		locale,
		appVersionCode,
	)
}

// ExportAsBytes exports selected *Instagram object as []byte
func (insta *Instagram) ExportAsBytes() ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := insta.ExportIO(buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// ExportAsBase64String exports selected *Instagram object as base64 encoded string
func (insta *Instagram) ExportAsBase64String() (string, error) {
	bytes, err := insta.ExportAsBytes()
	if err != nil {
		return "", err
	}

	sEnc := base64.StdEncoding.EncodeToString(bytes)
	return sEnc, nil
}

// ImportFromBytes imports instagram configuration from an array of bytes.
//
// This function does not set proxy automatically. Use SetProxy after this call.
func ImportFromBytes(inputBytes []byte, args ...interface{}) (*Instagram, error) {
	return ImportReader(bytes.NewReader(inputBytes), args...)
}

// ImportFromBase64String imports instagram configuration from a base64 encoded string.
//
// This function does not set proxy automatically. Use SetProxy after this call.
func ImportFromBase64String(base64String string, args ...interface{}) (*Instagram, error) {
	sDec, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return ImportFromBytes(sDec, args...)
}

func MergeMapI(one map[string]interface{}, extra ...map[string]interface{}) map[string]interface{} {
	for _, e := range extra {
		for k, v := range e {
			one[k] = v
		}
	}
	return one
}

func MergeMapS(one map[string]string, extra ...map[string]string) map[string]string {
	for _, e := range extra {
		for k, v := range e {
			one[k] = v
		}
	}
	return one
}

func read16(b []byte, keys []string, offset int) (int, error) {
	start, err := getStartByte(b, keys, offset)
	if err != nil {
		return -1, nil
	}
	r := binary.BigEndian.Uint16(b[start+offset:])
	return int(r), nil
}

func read32(b []byte, keys []string, offset int) (int, error) {
	start, err := getStartByte(b, keys, offset)
	if err != nil {
		return -1, nil
	}
	r := binary.BigEndian.Uint32(b[start+offset:])
	return int(r), nil
}

func getStartByte(b []byte, keys []string, offset int) (int, error) {
	start := 0
	for _, key := range keys {
		n := bytes.Index(b[start:], []byte(key))
		if n == -1 {
			return -1, ErrByteIndexNotFound
		}
		start += n + len(key)
	}
	return start, nil
}

func getSupCap() (string, error) {
	query := []trayRequest{
		{"SUPPORTED_SDK_VERSIONS", supportedSdkVersions},
		{"FACE_TRACKER_VERSION", facetrackerVersion},
		{"segmentation", segmentation},
		{"COMPRESSION", compression},
		{"world_tracker", worldTracker},
		{"gyroscope", gyroscope},
	}
	data, err := json.Marshal(query)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


// StoryReelMention represent story reel mention
type StoryReelMention struct {
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Z           int     `json:"z"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Rotation    float64 `json:"rotation"`
	IsPinned    int     `json:"is_pinned"`
	IsHidden    int     `json:"is_hidden"`
	IsSticker   int     `json:"is_sticker"`
	IsFBSticker int     `json:"is_fb_sticker"`
	User        User
	DisplayType string `json:"display_type"`
}

// StoryCTA represent story cta
type StoryCTA struct {
	Links []struct {
		LinkType                                int         `json:"linkType"`
		WebURI                                  string      `json:"webUri"`
		AndroidClass                            string      `json:"androidClass"`
		Package                                 string      `json:"package"`
		DeeplinkURI                             string      `json:"deeplinkUri"`
		CallToActionTitle                       string      `json:"callToActionTitle"`
		RedirectURI                             interface{} `json:"redirectUri"`
		LeadGenFormID                           string      `json:"leadGenFormId"`
		IgUserID                                string      `json:"igUserId"`
		AppInstallObjectiveInvalidationBehavior interface{} `json:"appInstallObjectiveInvalidationBehavior"`
	} `json:"links"`
}

// StoryMedia is the struct that handles the information from the methods to
// get info about Stories.
type StoryMedia struct {
	Reel       Reel         `json:"reel"`
	Status     string       `json:"status"`
}

// Reel represents a single user's story collection.
// Every user has one reel, and one reel can contain many story items
type Reel struct {
	insta *Instagram

	ID                     interface{} `json:"id"`
	Items                  []*Item     `json:"items"`
	MediaCount             int         `json:"media_count"`
	MediaIDs               []int64     `json:"media_ids"`
	Muted                  bool        `json:"muted"`
	LatestReelMedia        int64       `json:"latest_reel_media"`
	LatestBestiesReelMedia float64     `json:"latest_besties_reel_media"`
	ExpiringAt             float64     `json:"expiring_at"`
	Seen                   float64     `json:"seen"`
	SeenRankedPosition     int         `json:"seen_ranked_position"`
	CanReply               bool        `json:"can_reply"`
	CanGifQuickReply       bool        `json:"can_gif_quick_reply"`
	ClientPrefetchScore    float64     `json:"client_prefetch_score"`
	Title                  string      `json:"title"`
	CanReshare             bool        `json:"can_reshare"`
	ReelType               string      `json:"reel_type"`
	ReelMentions           []string    `json:"reel_mentions"`
	PrefetchCount          int         `json:"prefetch_count"`
	// this field can be int or bool
	HasBestiesMedia       interface{} `json:"has_besties_media"`
	HasPrideMedia         bool        `json:"has_pride_media"`
	HasVideo              bool        `json:"has_video"`
	IsCacheable           bool        `json:"is_cacheable"`
	IsSensitiveVerticalAd bool        `json:"is_sensitive_vertical_ad"`
	RankedPosition        int         `json:"ranked_position"`
	RankerScores          struct {
		Fp   float64 `json:"fp"`
		Ptap float64 `json:"ptap"`
		Vm   float64 `json:"vm"`
	} `json:"ranker_scores"`
	StoryRankingToken    string `json:"story_ranking_token"`
	FaceFilterNuxVersion int    `json:"face_filter_nux_version"`
	HasNewNuxStory       bool   `json:"has_new_nux_story"`
	User                 User   `json:"user"`
}

type trayRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
