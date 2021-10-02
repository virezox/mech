package insta

import (
   "bytes"
   "crypto/md5"
   "encoding/hex"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
)

func generateMD5Hash(text string) string {
   hasher := md5.New()
   hasher.Write([]byte(text))
   return hex.EncodeToString(hasher.Sum(nil))
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

// Instagram represent the main API handler. We recommend to use Export and
// Import functions after first Login. Also you can use SetProxy and UnsetProxy
// to set and unset proxy. Golang also provides the option to set a proxy using
// HTTP_PROXY env var.
type Instagram struct {
   user string
   pass string
   // id: android-1923fjnma8123
   dID string
   // contains header options set by Instagram
   headerOptions map[string]string
   // User-Agent
   userAgent string
   c *http.Client
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
   seed := generateMD5Hash(username + password)
   hash := generateMD5Hash(seed + "12345")
   insta := &Instagram{
      c: &http.Client{
         Transport: &http.Transport{Proxy: http.ProxyFromEnvironment},
      },
      dID: "android-" + hash[:16],
      headerOptions: map[string]string{},
      pass: password,
      user: username,
      userAgent: fmt.Sprintf(
         "Instagram %s Android (%d/%d; %s; %s; %s; %s; %s; %s; %s; %s)",
         "195.0.0.31.123",
         dev.AndroidVersion,
         dev.AndroidRelease,
         dev.ScreenDpi,
         dev.ScreenResolution,
         dev.Manufacturer,
         dev.Model,
         dev.CodeName,
         dev.Chipset,
         "en_US",
         "302733750",
      ),
   }
   insta.headerOptions["X-Ig-Www-Claim"] = "0"
   return insta
}

// Export exports selected *Instagram object options to an io.Writer
func (insta *Instagram) ExportIO(writer io.Writer) error {
   return json.NewEncoder(writer).Encode(insta.headerOptions)
}

// Login performs instagram login sequence in close resemblance to the android
// apk. Password will be deleted after login.
func (insta *Instagram) Login() error {
   timestamp := strconv.Itoa(int(time.Now().Unix()))
   encrypted := fmt.Sprintf("#PWD_INSTAGRAM:0:%s:%s", timestamp, insta.pass)
   result, err := json.Marshal(
      map[string]interface{}{
         // need this
         "device_id":           insta.dID,
         // maybe
         "enc_password":        encrypted,
         "country_code":        "[{\"country_code\":\"44\",\"source\":[\"default\"]}]",
         "google_tokens":       "[]",
         "login_attempt_count": 0,
         "username":            insta.user,
      },
   )
   if err != nil {
      return err
   }
   body, err := insta.sendRequest(
      map[string]string{
         "signed_body": "SIGNATURE." + string(result),
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

func (insta *Instagram) sendRequest(query map[string]string) ([]byte, error) {
   if insta == nil {
      return nil, errors.New(
         "insta has not been defined, this is most likely a bug in the code. " +
         "Please backtrack which call this error came from, and open an issue " +
         "detailing exactly how you got to this error",
      )
   }
   vs := make(url.Values)
   for k, v := range query {
      vs.Add(k, v)
   }
   reqData := bytes.NewBuffer([]byte{})
   reqData.WriteString(vs.Encode())
   req, err := http.NewRequest(
      "POST", "https://i.instagram.com/api/v1/accounts/login/", reqData,
   )
   if err != nil {
      return nil, err
   }
   headers := map[string]string{
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
      "User-Agent":                  insta.userAgent,
      "X-Bloks-Is-Layout-Rtl":       "false",
      "X-Bloks-Is-Panorama-Enabled": "true",
      "X-Fb-Client-Ip":              "True",
      "X-Fb-Http-Engine":            "Liger",
      "X-Fb-Server-Cluster":         "True",
      "X-Ig-App-Startup-Country":    "unkown",
   }
   for key, val := range headers {
      req.Header.Set(key, val)
   }
   for key, value := range insta.headerOptions {
      if value != "" {
         req.Header.Set(key, value)
      }
   }
   dum, err := httputil.DumpRequest(req, true)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(append(dum, '\n'))
   resp, err := insta.c.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   insta.extractHeaders(resp.Header)
   return body, nil
}
