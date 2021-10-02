package insta

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
   "time"
)

func login(username, password string) error {
   // BEGIN ////////////////////////////////////////////////////////////////////
   sig := map[string]string{
      "device_id": "android-0123456789abcdef",
      "enc_password": fmt.Sprintf(
         "#PWD_INSTAGRAM:0:%v:%v", time.Now().Unix(), password,
      ),
      "username": username,
   }
   result, err := json.Marshal(sig)
   if err != nil {
      return err
   }
   val := url.Values{
      "signed_body": {
         "SIGNATURE." + string(result),
      },
   }
   // END //////////////////////////////////////////////////////////////////////
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
   }
   dReq, err := httputil.DumpRequest(req, true)
   if err != nil {
      return err
   }
   os.Stdout.Write(dReq)
   resp, err := new(http.Transport).RoundTrip(req)
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
