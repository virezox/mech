package insta

import (
   "crypto/md5"
   "encoding/hex"
   "encoding/json"
   "fmt"
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

// Instagram represent the main API handler. We recommend to use Export and
// Import functions after first Login. Also you can use SetProxy and UnsetProxy
// to set and unset proxy. Golang also provides the option to set a proxy using
// HTTP_PROXY env var.
type Instagram struct {
   user string
   pass string
   // id: android-1923fjnma8123
   dID string
   c *http.Client
}

// New creates Instagram structure
func New(username, password string) *Instagram {
   seed := generateMD5Hash(username + password)
   hash := generateMD5Hash(seed + "12345")
   insta := &Instagram{
      c: &http.Client{
         Transport: &http.Transport{Proxy: http.ProxyFromEnvironment},
      },
      dID: "android-" + hash[:16],
      pass: password,
      user: username,
   }
   return insta
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
         "country_code": "[{\"country_code\":\"44\",\"source\":[\"default\"]}]",
         "enc_password":        encrypted,
         "google_tokens":       "[]",
         "login_attempt_count": 0,
         "username":            insta.user,
      },
   )
   if err != nil {
      return err
   }
   val := url.Values{
      "signed_body": {
         "SIGNATURE." + string(result),
      },
   }
   req, err := http.NewRequest(
      "POST", "https://i.instagram.com/api/v1/accounts/login/",
      strings.NewReader(val.Encode()),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded; charset=UTF-8"},
      "User-Agent": {"Instagram 195.0.0.31.123 Android (30/11; 560dpi; 1440x2898; samsung; SM-G975F; beyond2; exynos9820; en_US; 302733750)"},
      "X-Bloks-Is-Layout-Rtl": {"false"},
      "X-Bloks-Is-Panorama-Enabled": {"true"},
      "X-Fb-Client-Ip": {"True"},
      "X-Fb-Http-Engine": {"Liger"},
      "X-Fb-Server-Cluster": {"True"},
      "X-Ig-App-Startup-Country": {"unkown"},
      "X-Ig-Www-Claim": {"0"},
   }
   dReq, err := httputil.DumpRequest(req, true)
   if err != nil {
      return err
   }
   os.Stdout.Write(append(dReq, '\n'))
   resp, err := insta.c.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   dRes, err := httputil.DumpResponse(resp, true)
   if err != nil {
      return err
   }
   os.Stdout.Write(dRes)
   return nil
}
