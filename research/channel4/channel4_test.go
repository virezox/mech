package channel4

import (
   "fmt"
   "testing"
)

const payload = `
{"request_id":5273616,"token":"amJBaGREdUnQ37eeW6Ur2QH3z37Mm6XTJxcWxn2oRu6_IFowZAzBuQ3PCoJiVXvDze8E8t0HMEM86VYt4ExpAVW01d_mtJHrsJEUDCjXcEUphoWCOXmldn__zkonMSzrQvi0Dtgor28EsaZH0bUAdFDDajNeZA8b","video":{"type":"ondemand","url":"https://cf.jos.c4assets.com/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13497048883755456981&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1650945199&e=600&st=KN6vSjmXWCHpzKWAgkNkJaLuUJiqU7kX11JS9XZuOBo"},"message":"CAQ="}
`

func TestWidevine(t *testing.T) {
   vine, err := newWidevine(payload)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vine)
}

const (
   in = "00000000-0000-0000-0000-000004246624"
   out = "AAAAMnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABISEAAAAAAAAAAAAAAAAAQkZiQ="
)

func TestPSSH(t *testing.T) {
   pssh, err := createPSSH(in)
   if err != nil {
      t.Fatal(err)
   }
   if pssh != out {
      t.Fatal(pssh)
   }
}
