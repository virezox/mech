package tiktok

import (
   "fmt"
   "net/url"
   "testing"
)

const addr =
   "https://api-h2.tiktokv.com" +
   "/aweme/v1/aweme/detail/?aweme_id=7038818332270808325"

var deviceParams = url.Values{
   "aid":[]string{"1180"},
   "app_name":[]string{"trill"},
   "channel":[]string{"apkpure"},
   "device_id":[]string{"7031670777339250182"},
   "device_platform":[]string{"android"},
   "device_type":[]string{"ONEPLUS A3010"},
   "iid":[]string{"7032045377013942018"},
   "os_api":[]string{"25"},
   "os_version":[]string{"7.1.1"},
   "version_code":[]string{"170804"},
   "version_name":[]string{"17.8.4"},
}

func TestTikTok(t *testing.T) {
   _, head, err := signURL(addr, deviceParams)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(head["x-gorgon"])
}
