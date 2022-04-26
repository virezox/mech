package channel4

import (
   "fmt"
   "strings"
   "testing"
)

var token = strings.NewReader(`
{"request_id":5273616,"token":"UmJmV092VmjsUfEOs8Brlf_gaikoKil7Df9cR8qFbiEgiQz6-laMJuNWSiKyeJcgCVJdSlPPohyr1Evg04ZhzPcAbxo3shda7hp-NaMIDUFUdMBn8mNxb2nD_au8j0BbAHSEU2-IQFgLJqMF3C_WZL1CRgV6R2V6","video":{"type":"ondemand","url":"https://ak-jos-c4assets-com.akamaized.net/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13636097608999027874&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1651013628&e=600&st=pAvTinlirekGcg08lMDAcgf6n7PBOJj_U3DsvIVXMMk"},"message":"CAQ="}
`)

func TestDecrypt(t *testing.T) {
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
