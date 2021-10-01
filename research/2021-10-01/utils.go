package goinsta

import (
   "crypto/md5"
   "crypto/rand"
   "encoding/hex"
   "fmt"
   "io"
   "strconv"
   "time"
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


const (
	volatileSeed = "12345"
)

func generateMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateDeviceID(seed string) string {
	hash := generateMD5Hash(seed + volatileSeed)
	return "android-" + hash[:16]
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

type ChallengeStepData struct {
	Choice           string      `json:"choice"`
	FbAccessToken    string      `json:"fb_access_token"`
	BigBlueToken     string      `json:"big_blue_token"`
	GoogleOauthToken string      `json:"google_oauth_token"`
	Email            string      `json:"email"`
	SecurityCode     string      `json:"security_code"`
	ResendDelay      interface{} `json:"resend_delay"`
	ContactPoint     string      `json:"contact_point"`
	FormType         string      `json:"form_type"`
}

type Challenge struct {
	insta *Instagram

	LoggedInUser *Account `json:"logged_in_user,omitempty"`
	UserID       int64    `json:"user_id"`
	Status       string   `json:"status"`

	ApiPath           string            `json:"api_path"`
	Context           *ChallengeContext `json:"challenge_context"`
	FlowRenderType    int               `json:"flow_render_type"`
	HideWebviewHeader bool              `json:"hide_webview_header"`
	Lock              bool              `json:"lock"`
	Logout            bool              `json:"logout"`
	NativeFlow        bool              `json:"native_flow"`
	URL               string            `json:"url"`

	TwoFactorRequired bool
	TwoFactorInfo     TwoFactorInfo
}

type ChallengeContext struct {
	TypeEnum    string            `json:"challenge_type_enum"`
	IsStateless bool              `json:"is_stateless"`
	Action      string            `json:"action"`
	NonceCode   string            `json:"nonce_code"`
	StepName    string            `json:"step_name"`
	StepData    ChallengeStepData `json:"step_data"`
	UserID      int64             `json:"user_id"`
}

type TwoFactorInfo struct {
	insta *Instagram

	ElegibleForMultipleTotp    bool   `json:"elegible_for_multiple_totp"`
	ObfuscatedPhoneNr          string `json:"obfuscated_phone_number"`
	PendingTrustedNotification bool   `json:"pending_trusted_notification"`
	ShouldOptInTrustedDevice   bool   `json:"should_opt_in_trusted_device_option"`
	ShowMessengerCodeOption    bool   `json:"show_messenger_code_option"`
	ShowTrustedDeviceOption    bool   `json:"show_trusted_device_option"`
	SMSNotAllowedReason        string `json:"sms_not_allowed_reason"`
	SMSTwoFactorOn             bool   `json:"sms_two_factor_on"`
	TotpTwoFactorOn            bool   `json:"totp_two_factor_on"`
	WhatsappTwoFactorOn        bool   `json:"whatsapp_two_factor_on"`
	TwoFactorIdentifier        string `json:"two_factor_identifier"`
	Username                   string `json:"username"`

	PhoneVerificationSettings phoneVerificationSettings `json:"phone_verification_settings"`
}

type phoneVerificationSettings struct {
	MaxSMSCount          int  `json:"max_sms_count"`
	ResendSMSDelaySec    int  `json:"resend_sms_delay_sec"`
	RobocallAfterMaxSms  bool `json:"robocall_after_max_sms"`
	RobocallCountDownSec int  `json:"robocall_count_down_time_sec"`
}
