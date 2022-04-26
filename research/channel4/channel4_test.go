package channel4

import (
   "fmt"
   "testing"
)

var token = []byte(`
{"request_id":5273616,"token":"RmVycnNYTFVrXrl_6fPWvGgZOzj0h2U5X_TJv2BkerbH4CRhndeazwutTIT-HWs-c8Yn_zpvMEViC1wOOAel8nq_cjFXZITMgHkYtwbKZJsmHvkfPot44--PcU4uBtKYpTQwRy3odNLa4q5BCrKcsYmWcH5iyOeV","video":{"type":"ondemand","url":"https://ak-jos-c4assets-com.akamaized.net/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13630749547086512690&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1650990334&e=600&st=PEkooWh_S71X0m1tqNXdAjDeAd16bOHd70bAdjJ7mDM"},"message":"CAQ="}
`)

const kid = "00000000-0000-0000-0000-000004246624"

func TestWidevine(t *testing.T) {
   header, err := NewHeader(kid)
   if err != nil {
      t.Fatal(err)
   }
   payload, err := header.Payload(token)
   if err != nil {
      t.Fatal(err)
   }
   widevine, err := payload.Widevine()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", widevine)
   return
   key, err := widevine.Decrypt(header)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(string(key))
}
