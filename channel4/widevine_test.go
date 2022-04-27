package channel4

import (
   "fmt"
   "strings"
   "testing"
)

var token = strings.NewReader(`
{"request_id":5273616,"token":"bU1kZ2pJc0Zz7LQ4cnDVX8JyJjVzqvI0q1gWvhaqKTHXiZbfr5ZL1P3NsxgBKjhSOjVZyCEw9Y7EjDW8WoCz7ZiyMa4U-Z5VDdem-Wpwkbpq-itKcrC3_HWkkQ0TlIzINNjdGXc6e1FM82SS-17jY4FGLqfFSb7X","video":{"type":"ondemand","url":"https://cf.jos.c4assets.com/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13636097608999265113&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1651026452&e=600&st=sxlm1NjF2vmIkCg36_EQAHg5OKk8_O0H6vmUTRVo4UA"},"message":"CAQ="}
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
   fmt.Println(key)
}
