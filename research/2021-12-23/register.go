package tiktok

import (
   "encoding/hex"
   "fmt"
   "net/http"
   "net/url"
   "strings"
)

const openudid = "f8r3ikpjgznbexj"

var body = fmt.Sprintf(`
{
  "magic_tag": "ss_app_log",
  "header": {
    "channel": "googleplay",
    "device_model": "US601XSeriesXChargeAmazonPrimeExclusiveLTE-A/XPower2",
    "openudid": "ezj94b2tx6kpg8b",
    "os_version": "7.1.2",
    "version_code": 130211,
    "access": "wifi",
    "aid": 1233,
    "app_version": "13.2.11",
    "appkey": "5559e28267e58eb4c1000012",
    "carrier": "AT&T Wireless Inc.",
    "clientudid": "f13af0c2-be43-499b-9a12-106fdb0fd92c",
    "cpu_abi": "arm64-v8a",
    "density_dpi": "267",
    "device_brand": "LG",
    "display_density": "mdpi",
    "display_name": "TikTok",
    "google_aid": "0585f9dd-ee8e-400d-aa9a-663a6ca79976",
    "language": "en",
    "manifest_version_code": 2019092901,
    "mc": "54:52:00:4c:d8:11",
    "mcc_mnc": "310980",
    "not_request_sender": 0,
    "os": "Android",
    "os_api": "25",
    "package": "com.zhiliaoapp.musically",
    "release_build": "eaeeb2f_20190929",
    "resolution": "720x1280",
    "sdk_version": "2.5.5.8",
    "sim_region": "us",
    "sim_serial_number": [],
    "timezone": 1,
    "tz_name": "America\\/New_York",
    "tz_offset": 0,
    "update_version_code": 2019092901,
  },
  "_gen_time": 1640287897348
}
`, device_model, openudid)

func register() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST",
      "https://log2.musical.ly/service/2/device_register/",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   gorgon, err := genXGorgon(body)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/json; charset=utf-8"},
      "sdk-version": {"1"},
      "user-agent": {"com.zhiliaoapp.musically/2019092901 (Linux; U; Android 7.1.2 en; L83BLStylo3LTE-A; Build/L83BLStylo3LTE-A; Cronet/58.0.2991.0)"},
      "x-gorgon": {hex.EncodeToString(gorgon)},
      "x-khronos": {"0"},
   }
   return new(http.Transport).RoundTrip(req)
}

func TestRegister(t *testing.T) {
   res, err := register()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(append(buf, '\n'))
}
