package goinsta

import (
	"encoding/json"
	"fmt"
	"net/http"
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

// Error400 is error returned by HTTP 400 status code.
type Error400 struct {
	ChallengeError
	Endpoint   string `json:"endpoint"`
	Action     string `json:"action"`
	StatusCode string `json:"status_code"`
	Payload    struct {
		ClientContext string `json:"client_context"`
		Message       string `json:"message"`
	} `json:"payload"`
	DebugInfo struct {
		Message   string `json:"string"`
		Retriable bool   `json:"retriable"`
		Type      string `json:"type"`
	} `json:"debug_info"`
	Code   int
	Status string `json:"status"`
}

func (e Error400) Error() string {
	var msg string
	if e.Payload.Message != "" {
		msg = e.Payload.Message
	}
	if e.DebugInfo.Message != "" {
		msg = e.DebugInfo.Message
	}
	if e.ChallengeError.Message != "" {
		if msg != "" {
			msg += "; " + e.ChallengeError.Message
		} else {
			msg = e.ChallengeError.Message
		}
	}

	if e.Code == 0 {
		e.Code = 400
	}
	return fmt.Sprintf("Request Status Code %d: %s, %s", e.Code, e.Status, msg)
}

// ChallengeError is error returned by HTTP 400 status code.
type ChallengeError struct {
	insta *Instagram

	Message   string `json:"message"`
	Challenge struct {
		URL               string `json:"url"`
		APIPath           string `json:"api_path"`
		HideWebviewHeader bool   `json:"hide_webview_header"`
		Lock              bool   `json:"lock"`
		Logout            bool   `json:"logout"`
		NativeFlow        bool   `json:"native_flow"`
	} `json:"challenge"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

func (e ChallengeError) Error() string {
	return fmt.Sprintf("Challenge Required: %s, %s", e.Status, e.Message)
}

// Nametag is part of the account information.
type Nametag struct {
	Mode          int64       `json:"mode"`
	Gradient      json.Number `json:"gradient,Number"`
	Emoji         string      `json:"emoji"`
	SelfieSticker json.Number `json:"selfie_sticker,Number"`
}

type ErrChallengeProcess struct {
	StepName string
}

func (ec ErrChallengeProcess) Error() string {
	return ec.StepName
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

type ScoresBootstrapUsers struct {
	Status   string `json:"status"`
	Surfaces []struct {
		Name      string         `json:"name"`
		RankToken string         `json:"rank_token"`
		Scores    map[string]int `json:"scores"`
		TTLSecs   int            `json:"ttl_secs"`
	} `json:"surfaces"`
	Users []*User `json:"users"`
}

type CommentOffensive struct {
	BullyClassifier  float64 `json:"bully_classifier"`
	SexualClassifier float64 `json:"sexual_classifier"`
	HateClassifier   float64 `json:"hate_classifier"`
	IsOffensive      bool    `json:"is_offensive"`
	Status           string  `json:"status"`
	TextLanguage     string  `json:"text_language"`
}
