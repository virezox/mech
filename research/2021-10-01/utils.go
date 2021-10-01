package goinsta

import (
   "bytes"
   "crypto/hmac"
   "crypto/md5"
   "crypto/rand"
   "crypto/sha256"
   "encoding/base64"
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
