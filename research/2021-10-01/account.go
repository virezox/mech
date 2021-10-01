package goinsta

import (
   "crypto/aes"
   "crypto/cipher"
   "crypto/rand"
   "crypto/rsa"
   "crypto/x509"
   "encoding/base64"
   "encoding/binary"
   "encoding/pem"
   "errors"
   "fmt"
   "net/http"
   "strconv"
   "time"
)

// Endpoints (with format vars)
const (
	// Login
	urlMsisdnHeader               = "accounts/read_msisdn_header/"
	urlGetPrefill                 = "accounts/get_prefill_candidates/"
	urlContactPrefill             = "accounts/contact_point_prefill/"
	urlGetAccFamily               = "multiple_accounts/get_account_family/"
	urlZrToken                    = "zr/token/result/"
	urlLogin                      = "accounts/login/"
	urlLogout                     = "accounts/logout/"
	urlAutoComplete               = "friendships/autocomplete_user_list/"
	urlQeSync                     = "qe/sync/"
	urlSync                       = "launcher/sync/"
	urlLogAttribution             = "attribution/log_attribution/"
	urlMegaphoneLog               = "megaphone/log/"
	urlExpose                     = "qe/expose/"
	urlGetNdxSteps                = "devices/ndx/api/async_get_ndx_ig_steps/"
	urlBanyan                     = "banyan/banyan/"
	urlCooldowns                  = "qp/get_cooldowns/"
	urlFetchConfig                = "loom/fetch_config/"
	urlBootstrapUserScores        = "scores/bootstrap/users/"
	urlStoreClientPushPermissions = "notifications/store_client_push_permissions/"
	urlProcessContactPointSignals = "accounts/process_contact_point_signals/"
	// urls
	baseUrl        = "https://i.instagram.com/"
	instaAPIUrl    = "https://i.instagram.com/api/v1/"
	instaAPIUrlb   = "https://b.i.instagram.com/api/v1/"
	instaAPIUrlv2  = "https://i.instagram.com/api/v2/"
	instaAPIUrlv2b = "https://b.i.instagram.com/api/v2/"
	// header values
	bloksVerID         = "927f06374b80864ae6a0b04757048065714dc50ff15d2b8b3de8d0b6de961649"
	fbAnalytics        = "567067343352427"
	igCapabilities     = "3brTvx0="
	connType           = "WIFI"
	instaSigKeyVersion = "4"
	locale             = "en_US"
	appVersion         = "195.0.0.31.123"
	appVersionCode     = "302733750"
)

// Account & Login Errors
var ErrBadPassword     = errors.New("Password is incorrect")

var ErrInstaNotDefined = errors.New(
   "Insta has not been defined, this is most likely a bug in the code. Please " +
   "backtrack which call this error came from, and open an issue detailing " +
   "exactly how you got to this error.",
)

var ErrAllSaved = errors.New("Unable to call function for collection all posts")

var (
	defaultHeaderOptions = map[string]string{
		"X-Ig-Www-Claim": "0",
	}
	omitAPIHeadersExclude = []string{
		"X-Ig-Bandwidth-Speed-Kbps",
		"Ig-U-Shbts",
		"X-Ig-Mapped-Locale",
		"X-Ig-Family-Device-Id",
		"X-Ig-Android-Id",
		"X-Ig-Timezone-Offset",
		"X-Ig-Device-Locale",
		"X-Ig-Device-Id",
		"Ig-Intended-User-Id",
		"X-Ig-App-Locale",
		"X-Bloks-Is-Layout-Rtl",
		"X-Pigeon-Rawclienttime",
		"X-Bloks-Version-Id",
		"X-Ig-Bandwidth-Totalbytes-B",
		"X-Ig-Bandwidth-Totaltime-Ms",
		"X-Ig-App-Startup-Country",
		"X-Ig-Www-Claim",
		"X-Bloks-Is-Panorama-Enabled",
	}
	// Default Device
	GalaxyS10 = Device{
		Manufacturer:     "samsung",
		Model:            "SM-G975F",
		CodeName:         "beyond2",
		AndroidVersion:   30,
		AndroidRelease:   11,
		ScreenDpi:        "560dpi",
		ScreenResolution: "1440x2898",
		Chipset:          "exynos9820",
	}
	timeOffset = getTimeOffset()
)

func AESGCMEncrypt(key, data, additionalData []byte) (iv, encrypted, tag []byte, err error) {
	iv = make([]byte, 12)
	rand.Read(iv)

	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when creating cipher: %s", err))
		fmt.Println(err)
		return
	}

	var aesgcm cipher.AEAD
	aesgcm, err = cipher.NewGCM(block)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when creating gcm: %s", err))
		fmt.Println(err)
		return
	}

	encrypted = aesgcm.Seal(nil, iv, data, additionalData)
	tag = encrypted[len(encrypted)-16:]       // Extracting last 16 bytes authentication tag
	encrypted = encrypted[:len(encrypted)-16] // Extracting raw Encrypted data without IV & Tag for use in NodeJS

	return
}

func EncryptPassword(password, pubKeyEncoded string, pubKeyVersion int, t string) (string, error) {
	if t == "" {
		t = strconv.Itoa(int(time.Now().Unix()))
	}
	// Get the public key
	publicKey, err := RSADecodePublicKeyFromBase64(pubKeyEncoded)
	if err != nil {
		return "", err
	}
	// Data to be encrypted by RSA PKCS1
	randKey := make([]byte, 32)
	rand.Read(randKey)
	// Encrypt the random key that will be used to encrypt the password
	randKeyEncrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, randKey)
	if err != nil {
		return "", err
	}
	// Get the size of the encrypted random key
	randKeyEncryptedSize := make([]byte, 2)
	binary.LittleEndian.PutUint16(randKeyEncryptedSize[:], uint16(len(randKeyEncrypted)))
	// Encrypt the password using AES GCM with the random key
	iv, encrypted, tag, err := AESGCMEncrypt(randKey, []byte(password), []byte(t))
	if err != nil {
		return "", err
	}
	// Combine the parts
	s := []byte{}
	prefix := []byte{1, byte(pubKeyVersion)}
	parts := [][]byte{prefix, iv, randKeyEncryptedSize, randKeyEncrypted, tag, encrypted}
	for _, b := range parts {
		s = append(s, b...)
	}
	encoded := base64.StdEncoding.EncodeToString(s)
	return fmt.Sprintf("#PWD_INSTAGRAM:4:%s:%s", t, encoded), nil
}

func RSADecodePublicKeyFromBase64(pubKeyBase64 string) (*rsa.PublicKey, error) {
	pubKey, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	block, _ := pem.Decode(pubKey)
	pKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return pKey.(*rsa.PublicKey), nil
}

type Account struct {
   ID                         int64        `json:"pk"`
   insta *Instagram
}

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
   AndroidRelease   int    `json:"android_release"`
   AndroidVersion   int    `json:"android_version"`
   Chipset          string `json:"chipset"`
   CodeName         string `json:"code_name"`
   Manufacturer     string `json:"manufacturer"`
   Model            string `json:"model"`
   ScreenDpi        string `json:"screen_dpi"`
   ScreenResolution string `json:"screen_resolution"`
}

type accountResp struct {
   Account Account `json:"logged_in_user"`
   ErrorType         string         `json:"error_type"`
   Message           string         `json:"message"`
   Status  string  `json:"status"`
   TwoFactorRequired bool           `json:"two_factor_required"`
}
