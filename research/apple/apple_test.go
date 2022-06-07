package apple

import (
   "fmt"
   "os"
   "testing"
)

/*
master

https://play.itunes.apple.com/WebObjects/MZPlay.woa/hls/subscription/
playlist.m3u8?cc=SI&svcId=tvs.vds.4093&a=1524726231&isExternal=true&
brandId=tvs.sbd.4000&id=377684155&l=en-GB&aec=UHD

video_gr203 SD    508x254          avc1.64001f  avg:  258.4 kbps

https://play.itunes.apple.com/WebObjects/MZPlay.woa/hls/subscription/
stream/playlist.m3u8?cc=SI&g=203&cdn=vod-ap3-aoc.tv.apple.com&a=1484589502&
p=461374806&st=1821682630&a=1524726231&p=377684155&st=1491016092&a=1524197777&
p=368330428&st=1449517808&a=1524197722&p=368330432&st=1449518320&a=1524198082&
p=368330370&st=1449517650&a=1525078430&p=368329706&st=1449510065&a=1524197604&
p=368330236&st=1449518888&a=1524197554&p=368330322&st=1449518516&a=1524197773&
p=368330253&st=1449517978&a=1539152595&p=368283705&st=1449199972

https://vod-ap3-aoc.tv.apple.com/itunes-assets/HLSVideo126/v4/25/a3/dd/
25a3ddc3-9b9e-cfc2-97e2-5dcb03ffd255/
P377684155_A1524726231_FF_video_gr203_sdr_508x254_cbcs_--0.mp4

https://vod-ap3-aoc.tv.apple.com/itunes-assets/HLSVideo126/v4/25/a3/dd/
25a3ddc3-9b9e-cfc2-97e2-5dcb03ffd255/
P377684155_A1524726231_FF_video_gr203_sdr_508x254_cbcs_--1.m4s
*/
const (
   // tv.apple.com/us/episode/biscuits/umc.cmc.45cu44369hb2qfuwr3fihnr8e
   contentID = "umc.cmc.45cu44369hb2qfuwr3fihnr8e"
   // 22bdb0063805260307ee5045c0f3835a
   pssh = "data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgwC7YzAgICAgICBI88aJmwY="
)

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := OpenAuth(home, "mech/apple.json")
   if err != nil {
      t.Fatal(err)
   }
   privateKey, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   clientID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   request, err := auth.Request(privateKey, clientID, pssh)
   if err != nil {
      t.Fatal(err)
   }
   env, err := NewEnvironment()
   if err != nil {
      t.Fatal(err)
   }
   episode, err := NewEpisode(contentID)
   if err != nil {
      t.Fatal(err)
   }
   license, err := request.License(env, episode)
   if err != nil {
      t.Fatal(err)
   }
   key, err := license.Key()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestCreate(t *testing.T) {
   con, err := NewConfig()
   if err != nil {
      t.Fatal(err)
   }
   sign, err := con.Signin(email, password)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := sign.Auth()
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Create(home, "mech/apple.json"); err != nil {
      t.Fatal(err)
   }
}
