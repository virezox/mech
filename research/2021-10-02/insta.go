package insta

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
   "strconv"
   "time"
)

const (
   origin = "https://i.instagram.com"
   userAgent = "Instagram 207.0.0.39.120 Android"
)

type login http.Header

// Ig-Set-Authorization
func newLogin(user, pass string) (login, error) {
   buf := bytes.NewBufferString("signed_body=SIGNATURE.")
   now := strconv.FormatInt(time.Now().Unix(), 10)
   sig := map[string]string{
      "device_id": userAgent,
      "enc_password": "#PWD_INSTAGRAM:0:" + now + ":" + pass,
      "username": user,
   }
   if err := json.NewEncoder(buf).Encode(sig); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/api/v1/accounts/login/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {userAgent},
   }
   dum, err := httputil.DumpRequest(req, true)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   return login(res.Header), nil
}
