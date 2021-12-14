package tiktok

import (
   "net/http"
)

func get() (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "https://api-h2.tiktokv.com/aweme/v1/aweme/detail/", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "User-Agent":[]string{"okhttp/3.10.0.1"},
      "X-Gorgon":[]string{"036141108000b99bdc06b51fbe74c6b04a86b996c439e9b5a3ab"},
      "X-Khronos":[]string{"1639442178"},
   }
   req.URL.RawQuery = "aweme_id=7038818332270808325&os_api=25&device_type=ONEPLUS+A3010&app_name=trill&version_name=17.8.4&channel=apkpure&device_platform=android&iid=7032045377013942018&version_code=170804&device_id=7031670777339250182&os_version=7.1.1&aid=1180&_rticket=1639442178986"
   return new(http.Transport).RoundTrip(req)
}
