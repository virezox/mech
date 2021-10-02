package goinsta

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "strconv"
   "time"
)

// Endpoints (with format vars)
const (
   appVersion = "195.0.0.31.123"
   appVersionCode = "302733750"
   baseUrl        = "https://i.instagram.com/"
   bloksVerID = "927f06374b80864ae6a0b04757048065714dc50ff15d2b8b3de8d0b6de961649"
   connType = "WIFI"
   fbAnalytics = "567067343352427"
   igCapabilities = "3brTvx0="
   instaAPIUrl    = "https://i.instagram.com/api/v1/"
   instaAPIUrlb   = "https://b.i.instagram.com/api/v1/"
   instaAPIUrlv2  = "https://i.instagram.com/api/v2/"
   instaAPIUrlv2b = "https://b.i.instagram.com/api/v2/"
   locale = "en_US"
   urlLogin                      = "accounts/login/"
)

var defaultHeaderOptions = map[string]string{"X-Ig-Www-Claim": "0"}

var omitAPIHeadersExclude = []string{
   "Ig-Intended-User-Id",
   "Ig-U-Shbts",
   "X-Bloks-Is-Layout-Rtl",
   "X-Bloks-Is-Panorama-Enabled",
   "X-Bloks-Version-Id",
   "X-Ig-Android-Id",
   "X-Ig-App-Locale",
   "X-Ig-App-Startup-Country",
   "X-Ig-Bandwidth-Speed-Kbps",
   "X-Ig-Bandwidth-Totalbytes-B",
   "X-Ig-Bandwidth-Totaltime-Ms",
   "X-Ig-Device-Id",
   "X-Ig-Device-Locale",
   "X-Ig-Family-Device-Id",
   "X-Ig-Mapped-Locale",
   "X-Ig-Timezone-Offset",
   "X-Ig-Www-Claim",
   "X-Pigeon-Rawclienttime",
}

func defaultHandler(args ...interface{}) {
   fmt.Println(args...)
}

// ConfigFile is a structure to store the session information so that can be
// exported or imported.
type ConfigFile struct {
   DeviceID      string            `json:"device_id"`
   FamilyID      string            `json:"family_id"`
   HeaderOptions map[string]string `json:"header_options"`
   ID            int64             `json:"id"`
   PhoneID       string            `json:"phone_id"`
   Token         string            `json:"token"`
   UUID          string            `json:"uuid"`
   User          string            `json:"username"`
   XmidExpiry    int64             `json:"xmid_expiry"`
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

// Instagram represent the main API handler
//
// We recommend to use Export and Import functions after first Login.
//
// Also you can use SetProxy and UnsetProxy to set and unset proxy.
// Golang also provides the option to set a proxy using HTTP_PROXY env var.
type Instagram struct {
   user string
   pass string
   // id: android-1923fjnma8123
   dID string
   // family id, v4 uuid: 8b13e7b3-28f7-4e05-9474-358c6602e3f8
   fID string
   // uuid: 8493-1233-4312312-5123
   uuid string
   // token -- I think this is depricated, as I don't see any csrf tokens being
   // used anymore, but not 100% sure
   token string
   // phone id v4 uuid: fbf767a4-260a-490d-bcbb-ee7c9ed7c576
   pid string
   // ads id: 5b23a92b-3228-4cff-b6ab-3199f531f05b
   adid string
   // pigeonSessionId
   psID string
   // contains header options set by Instagram
   headerOptions map[string]string
   // expiry of X-Mid cookie
   xmidExpiry int64
   // User-Agent
   userAgent string
   c *http.Client
   // Set to true to debug reponses
   Debug bool
   // Non-error message handlers. By default they will be printed out,
   // alternatively you can e.g. pass them to a logger
   infoHandler  func(...interface{})
   warnHandler  func(...interface{})
   debugHandler func(...interface{})
}

// New creates Instagram structure
func New(username, password string) *Instagram {
   dev := Device{
      AndroidRelease:   11,
      AndroidVersion:   30,
      Chipset:          "exynos9820",
      CodeName:         "beyond2",
      Manufacturer:     "samsung",
      Model:            "SM-G975F",
      ScreenDpi:        "560dpi",
      ScreenResolution: "1440x2898",
   }
   insta := &Instagram{
      c: &http.Client{
         Transport: &http.Transport{Proxy: http.ProxyFromEnvironment},
      },
      dID: generateDeviceID(
         generateMD5Hash(username + password),
      ),
      debugHandler: defaultHandler,
      fID:           generateUUID(),
      headerOptions: map[string]string{},
      infoHandler:  defaultHandler,
      pass: password,
      pid:           generateUUID(),
      psID:          "UFS-" + generateUUID() + "-0",
      user: username,
      userAgent: fmt.Sprintf(
         "Instagram %s Android (%d/%d; %s; %s; %s; %s; %s; %s; %s; %s)",
         appVersion,
         dev.AndroidVersion,
         dev.AndroidRelease,
         dev.ScreenDpi,
         dev.ScreenResolution,
         dev.Manufacturer,
         dev.Model,
         dev.CodeName,
         dev.Chipset,
         locale,
         appVersionCode,
      ),
      uuid:          generateUUID(),
      warnHandler:  defaultHandler,
      xmidExpiry:    -1,
   }
   for k, v := range defaultHeaderOptions {
      insta.headerOptions[k] = v
   }
   return insta
}

// Export exports selected *Instagram object options to an io.Writer
func (insta *Instagram) ExportIO(writer io.Writer) error {
   config := ConfigFile{
      FamilyID:      insta.fID,
      HeaderOptions: map[string]string{},
      PhoneID:       insta.pid,
      Token:         insta.token,
      UUID:          insta.uuid,
      User:          insta.user,
      XmidExpiry:    insta.xmidExpiry,
   }
   for key, value := range insta.headerOptions {
      config.HeaderOptions[key] = value
   }
   bytes, err := json.Marshal(config)
   if err != nil {
      return err
   }
   if _, err := writer.Write(bytes); err != nil {
      return err
   }
   return nil
}

// Login performs instagram login sequence in close resemblance to the android
// apk. Password will be deleted after login.
func (insta *Instagram) Login() error {
   timestamp := strconv.Itoa(int(time.Now().Unix()))
   encrypted := fmt.Sprintf("#PWD_INSTAGRAM:0:%s:%s", timestamp, insta.pass)
   result, err := json.Marshal(
      map[string]interface{}{
         "adid":                insta.adid,
         "country_code":        "[{\"country_code\":\"44\",\"source\":[\"default\"]}]",
         "device_id":           insta.dID,
         "enc_password":        encrypted,
         "google_tokens":       "[]",
         "guid":                insta.uuid,
         "login_attempt_count": 0,
         "phone_id":            insta.fID,
         "username":            insta.user,
      },
   )
   if err != nil {
      return err
   }
   body, _, err := insta.sendRequest(
      &reqOptions{
         Endpoint: urlLogin,
         IsPost:   true,
         Query: map[string]string{
            "signed_body": "SIGNATURE." + string(result),
         },
      },
   )
   if err != nil {
      return err
   }
   var res struct {
      Error_Type string
      Message string
      Status  string
   }
   if err := json.Unmarshal(body, &res); err != nil {
      return err
   }
   if res.Status != "ok" {
      return fmt.Errorf("%+v", res)
   }
   return nil
}
