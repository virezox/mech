package channel4

import (
   "fmt"
   "testing"
)

var token = []byte(`
{"request_id":5273616,"token":"aWF1Z3FMUWbml_L9fdnu6hG8bu1RM0dE6Zm23ij9s_vavpqlQzL14guwdGr_AMqR7LUrSGIWhZ0-Z2BAGpDKoFq1TvXnd6c3mHmOafXHm5ghP8EaO6VevMwIVP4dRN84w41SxSGophZyJuAR2JOs2nYfN2-uOMEM","video":{"type":"ondemand","url":"https://ak-jos-c4assets-com.akamaized.net/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13502678456665898701&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1651004420&e=600&st=xkC9-eODpIT3iDusei87mJvl7H2QcPbu9enzWkF9JAQ"},"message":"CAQ="}
`)

func TestWidevine(t *testing.T) {
   payload, err := NewPayload(token)
   if err != nil {
      t.Fatal(err)
   }
   widevine, err := payload.Widevine()
   if err != nil {
      t.Fatal(err)
   }
   key, err := widevine.Decrypt()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(string(key))
}
