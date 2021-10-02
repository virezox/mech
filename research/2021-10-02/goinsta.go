package goinsta

import (
   "bytes"
   "compress/gzip"
   "crypto/md5"
   "crypto/rand"
   "encoding/hex"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "io/ioutil"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strconv"
   "strings"
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

func (insta *Instagram) extractHeaders(h http.Header) {
   extract := func(in string, out string) {
      x := h[in]
      if len(x) > 0 && x[0] != "" {
         // prevent from auth being set without token post login
         if in == "Ig-Set-Authorization" {
            old, ok := insta.headerOptions[out]
            if ok && len(old) != 0 {
               current := strings.Split(old, ":")
               newHeader := strings.Split(x[0], ":")
               if len(current[2]) > len(newHeader[2]) {
                  return
               }
            }
         }
         insta.headerOptions[out] = x[0]
      }
   }
   extract("Ig-Set-Authorization", "Authorization")
   extract("Ig-Set-Ig-U-Ds-User-Id", "Ig-U-Ds-User-Id")
   extract("Ig-Set-Ig-U-Ig-Direct-Region-Hint", "Ig-U-Ig-Direct-Region-Hint")
   extract("Ig-Set-Ig-U-Rur", "Ig-U-Rur")
   extract("Ig-Set-Ig-U-Shbid", "Ig-U-Shbid")
   extract("Ig-Set-Ig-U-Shbts", "Ig-U-Shbts")
   extract("Ig-Set-X-Mid", "X-Mid")
   extract("X-Ig-Set-Www-Claim", "X-Ig-Www-Claim")
}

func (insta *Instagram) sendRequest(o *reqOptions) (body []byte, h http.Header, err error) {
   if insta == nil {
      return nil, nil, errors.New(
         "insta has not been defined, this is most likely a bug in the code. " +
         "Please backtrack which call this error came from, and open an issue " +
         "detailing exactly how you got to this error",
      )
   }
   method := "GET"
   if o.IsPost {
      method = "POST"
   }
   if o.Connection == "" {
      o.Connection = "close"
   }
   if o.Timestamp == "" {
      o.Timestamp = strconv.Itoa(int(time.Now().Unix()))
   }
   var nu string
   if o.Useb {
      nu = instaAPIUrlb
   } else {
      nu = instaAPIUrl
   }
   if o.UseV2 && !o.Useb {
      nu = instaAPIUrlv2
   } else if o.UseV2 && o.Useb {
      nu = instaAPIUrlv2b
   }
   if o.OmitAPI {
      nu = baseUrl
      o.IgnoreHeaders = append(o.IgnoreHeaders, omitAPIHeadersExclude...)
   }
   u, err := url.Parse(nu + o.Endpoint)
   if err != nil {
      return nil, nil, err
   }
   vs := url.Values{}
   bf := bytes.NewBuffer([]byte{})
   reqData := bytes.NewBuffer([]byte{})
   for k, v := range o.Query {
      vs.Add(k, v)
   }
   // If DataBytes has been passed, use that as data, else use Query
   if o.DataBytes != nil {
      reqData = o.DataBytes
   } else {
      reqData.WriteString(vs.Encode())
   }
   var contentEncoding string
   if o.IsPost && o.Gzip {
      // If gzip encoding needs to be applied
      zw := gzip.NewWriter(bf)
      defer zw.Close()
      if _, err := zw.Write(reqData.Bytes()); err != nil {
         return nil, nil, err
      }
      if err := zw.Close(); err != nil {
         return nil, nil, err
      }
      contentEncoding = "gzip"
   } else if o.IsPost {
      // use post form if POST request
      bf = reqData
   } else {
      // append query to url if GET request
      for k, v := range u.Query() {
         vs.Add(k, strings.Join(v, " "))
      }
      u.RawQuery = vs.Encode()
   }
   var req *http.Request
   req, err = http.NewRequest(method, u.String(), bf)
   if err != nil {
      return
   }
   ignoreHeader := func(h string) bool {
      for _, k := range o.IgnoreHeaders {
         if k == h {
            return true
         }
      }
      return false
   }
   setHeaders := func(h map[string]string) {
      for k, v := range h {
         if v != "" && !ignoreHeader(k) {
            req.Header.Set(k, v)
         }
      }
   }
   headers := map[string]string{
      "Accept-Encoding": "gzip,deflate",
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
      "User-Agent":                  insta.userAgent,
      "X-Bloks-Is-Layout-Rtl":       "false",
      "X-Bloks-Is-Panorama-Enabled": "true",
      "X-Fb-Client-Ip":              "True",
      "X-Fb-Http-Engine":            "Liger",
      "X-Fb-Server-Cluster":         "True",
      "X-Ig-App-Startup-Country":    "unkown",
   }
   if contentEncoding != "" {
      headers["Content-Encoding"] = contentEncoding
   }
   setHeaders(headers)
   setHeaders(o.ExtraHeaders)
   for key, value := range insta.headerOptions {
      if value != "" && !ignoreHeader(key) {
         req.Header.Set(key, value)
      }
   }
   dum, err := httputil.DumpRequest(req, true)
   if err != nil {
      return nil, nil, err
   }
   os.Stdout.Write(append(dum, '\n'))
   resp, err := insta.c.Do(req)
   if err != nil {
      return nil, nil, err
   }
   defer resp.Body.Close()
   body, err = ioutil.ReadAll(resp.Body)
   if err != nil {
      return nil, nil, err
   }
   insta.extractHeaders(resp.Header)
   // Decode gzip encoded responses
   encoding := resp.Header.Get("Content-Encoding")
   if encoding != "" && encoding == "gzip" {
      buf := bytes.NewBuffer(body)
      zr, err := gzip.NewReader(buf)
      if err != nil {
         return nil, nil, err
      }
      body, err = ioutil.ReadAll(zr)
      if err != nil {
         return nil, nil, err
      }
      if err := zr.Close(); err != nil {
         return nil, nil, err
      }
   }
   return body, resp.Header.Clone(), err
}


const volatileSeed = "12345"

func generateDeviceID(seed string) string {
   hash := generateMD5Hash(seed + volatileSeed)
   return "android-" + hash[:16]
}

func generateMD5Hash(text string) string {
   hasher := md5.New()
   hasher.Write([]byte(text))
   return hex.EncodeToString(hasher.Sum(nil))
}

func generateUUID() string {
   uuid, err := newUUID()
   if err != nil {
      // default value when error occurred
      return "cb479ee7-a50d-49e7-8b7b-60cc1a105e22"
   }
   return uuid
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

type reqOptions struct {
   // Connection is connection header. Default is "close".
   Connection string
   // Endpoint is the request path of instagram api
   Endpoint string
   // Omit API omit the /api/v1/ part of the url
   OmitAPI bool
   // IsPost set to true will send request with POST method. By default this
   // option is false.
   IsPost bool
   // Compress post form data with gzip
   Gzip bool
   // UseV2 is set when API endpoint uses v2 url.
   UseV2 bool
   // Use b.i.instagram.com
   Useb bool
   // Query is the parameters of the request. This parameters are independents
   // of the request method (POST|GET)
   Query map[string]string
   // DataBytes can be used to pass raw data to a request, instead of a form
   // using the Query param. This is used for e.g. photo and vieo uploads.
   DataBytes *bytes.Buffer
   // List of headers to ignore
   IgnoreHeaders []string
   // Extra headers to add
   ExtraHeaders map[string]string
   // Timestamp
   Timestamp string
}
