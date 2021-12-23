package tiktok

import (
   "net/http"
   "strings"
)

func register() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST",
      "https://log2.musical.ly/service/2/device_register/",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/json; charset=utf-8"},
      "sdk-version": {"1"},
      "user-agent": {"com.zhiliaoapp.musically/2019092901 (Linux; U; Android 7.1.2 en; L83BLStylo3LTE-A; Build/L83BLStylo3LTE-A; Cronet/58.0.2991.0)"},
   }
   return new(http.Transport).RoundTrip(req)
}

const body = `
{
  "magic_tag": "ss_app_log",
  "header": {
    "access": "wifi",
    "aid": 1233,
    "app_version": "13.2.11",
    "appkey": "5559e28267e58eb4c1000012",
    "carrier": "Sprint Spectrum",
    "channel": "googleplay",
    "clientudid": "025fb2a9-54d4-4845-8706-10d3c2384bbf",
    "cpu_abi": "arm64-v8a",
    "density_dpi": "424",
    "device_brand": "Alcatel",
    "device_model": "OneTouchIdol5LTEUS6060C",
    "display_density": "mdpi",
    "display_name": "TikTok",
    "google_aid": "50a4edff-f434-42b1-be13-d1c87b8c7838",
    "language": "en",
    "manifest_version_code": 2019092901,
    "mc": "54:52:00:95:19:95",
    "mcc_mnc": "310120",
    "not_request_sender": 0,
    "openudid": "f8r3ikpjgznbexi",
    "os": "Android",
    "os_api": "25",
    "os_version": "7.1.2",
    "package": "com.zhiliaoapp.musically",
    "release_build": "eaeeb2f_20190929",
    "resolution": "1080x1920",
    "sdk_version": "2.5.5.8",
    "sim_region": "us",
    "sim_serial_number": [],
    "timezone": 1,
    "tz_name": "America/New_York",
    "tz_offset": 0,
    "update_version_code": 2019092901,
    "version_code": 130211
  },
  "_gen_time": 1640232924927
}
`
