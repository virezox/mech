package tiktok

import (
   "net/http"
   "strings"
)

const body = `
{
  "magic_tag": "ss_app_log",
  "header": {
    "display_name": "TikTok",
    "update_version_code": 2019092901,
    "manifest_version_code": 2019092901,
    "aid": 1233,
    "channel": "googleplay",
    "appkey": "5559e28267e58eb4c1000012",
    "package": "com.zhiliaoapp.musically",
    "app_version": "13.2.11",
    "version_code": 130211,
    "sdk_version": "2.5.5.8",
    "os": "Android",
    "os_version": "7.1.2",
    "os_api": "25",
    "device_model": "OneTouchIdol5LTEUS6060C",
    "device_brand": "Alcatel",
    "cpu_abi": "arm64-v8a",
    "release_build": "eaeeb2f_20190929",
    "density_dpi": "424",
    "display_density": "mdpi",
    "resolution": "1080x1920",
    "language": "en",
    "mc": "54:52:00:95:19:95",
    "timezone": 1,
    "access": "wifi",
    "not_request_sender": 0,
    "carrier": "Sprint Spectrum",
    "mcc_mnc": "310120",
    "google_aid": "50a4edff-f434-42b1-be13-d1c87b8c7838",
    "openudid": "f8r3ikpjgznbexi",
    "clientudid": "025fb2a9-54d4-4845-8706-10d3c2384bbf",
    "sim_serial_number": [],
    "tz_name": "America/New_York",
    "tz_offset": 0,
    "sim_region": "us"
  },
  "_gen_time": 1640232924927
}
`

func register() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST",
      "https://log2.musical.ly/service/2/device_register/",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "ac=wifi&channel=googleplay&aid=1233&app_name=musical_ly&version_code=130211&version_name=13.2.11&device_platform=android&ab_version=13.2.11&ssmix=a&device_type=L83BLStylo3LTE-A&device_brand=LG&language=en&os_api=25&os_version=7.1.2&uuid=331596668660397&openudid=d24qyq54nzhkl9o&manifest_version_code=2019092901&resolution=720x1280&dpi=258&update_version_code=2019092901&app_type=normal&sys_region=US&is_my_cn=0&pass-route=1&mcc_mnc=311880&pass-region=1&timezone_name=America%2FNew_York&carrier_region_v2=311&timezone_offset=0&build_number=13.2.11&region=US&uoo=0&app_language=en&carrier_region=US&locale=en&ac2=wifi5g&_rticket=1640225728462000&ts=1640225728462"
   req.Header = http.Header{
      "content-type": {"application/json; charset=utf-8"},
      "sdk-version": {"1"},
      "user-agent": {"com.zhiliaoapp.musically/2019092901 (Linux; U; Android 7.1.2 en; L83BLStylo3LTE-A; Build/L83BLStylo3LTE-A; Cronet/58.0.2991.0)"},
   }
   return new(http.Transport).RoundTrip(req)
}
