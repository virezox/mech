package bbcamerica

import (
   "github.com/89z/format"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

var LogLevel format.LogLevel

func unauth() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/unauth", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Amcn-Device-Id": {"!"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"bbca"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

func playback() {
   var body = strings.NewReader(`
   {
      "adtags": {
         "lat": 0,
         "url": "https://www.bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529",
         "playerWidth": 1920,
         "playerHeight": 1080,
         "ppid": 1,
         "mode": "on-demand"
      }
   }
   `)
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/playback-id/api/v1/playback/1052529", body,
   )
   if err != nil {
      panic(err)
   }
   req.Header["Authorization"] = []string{"Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJ1bmF1dGgiLCJhdWQiOiJyZXNvdXJjZV9zZXJ2ZXIiLCJhdXRoX3R5cGUiOiJiZWFyZXIiLCJyb2xlcyI6WyJ1bmF1dGgiXSwiaXNzIjoiaXAtMTAtMi0xNzktMjE3LmVjMi5pbnRlcm5hbCIsInRva2VuX3R5cGUiOiJhdXRoIiwiZXhwIjoxNjUzNTExMDEwLCJkZXZpY2UtaWQiOiIwNzZlMDMyOS1kMzRlLTQ4ODAtOWFhYS01NWJkMDc1ZjM4MWIiLCJpYXQiOjE2NTM1MTA0MTAsImp0aSI6IjUwMmU5MGFhLTBiNjItNDI2NC04OGI0LWNhZDRhZWYzZjM2NSJ9.vZtEDjovR_yEs4zbMa61NaFh7Rw3_tgJ0NjtUiZwmv8b0m3ZvOGedoX3oP3w4Ibvhj6_Ybt3sBFoBobM6BvrrXTHPx4bmGzGXBciZ56qe0oMSntGwPRfptRKKPdKzu6_H5VmqRv0I8BzdbSmdTfHbaQ3Hr9Hq9Uw74xc5Wdgy6gVGDPG7_zxBKod4b7IENHJcfug6oO6sSAAEG8Xo9N4cUKruWMRjEsZ3eBGJFHZsVBLaG_U5B3JXtFjzmK2COr1KUQwF7BrdAzLuwpTclJOa80qPSGSqXaMyoITc1bLP5WjX0Tjq7P5S6-w_hi1HooD1QLSAR7Y54qcpB7yD8jpWA"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["X-Amcn-Device-Ad-Id"] = []string{"!"}
   req.Header["X-Amcn-Language"] = []string{"en"}
   req.Header["X-Amcn-Network"] = []string{"bbca"}
   req.Header["X-Amcn-Platform"] = []string{"web"}
   req.Header["X-Amcn-Service-Id"] = []string{"bbca"}
   req.Header["X-Amcn-Tenant"] = []string{"amcn"}
   req.Header["X-Ccpa-Do-Not-Sell"] = []string{"passData"}
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
