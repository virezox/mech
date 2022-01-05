package main

import (
   "net/http"
   "net/http/httputil"
   "os"
   "net/url"
)

// 10 minute pass
func main() {
   req, err := http.NewRequest(
      "GET", "https://api2.musical.ly/aweme/v1/user/", nil,
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("User-Agent", "com.zhiliaoapp.musically/2018073103")
   req.URL.RawQuery = url.Values{
      "device_id":[]string{"6796244118550824449"},
      "iid":[]string{"6799460632921310982"},
      "device_type":[]string{"Redmi Note 8 Pro"},
      "os_version":[]string{"10"},
      "version_code":[]string{"800"},
      "user_id":[]string{"6671812072936292358"},
      "aid":[]string{"1233"},
      "app_name":[]string{"musical_ly"},
      "channel":[]string{"googleplay"},
      "device_platform":[]string{"android"},
   }.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
