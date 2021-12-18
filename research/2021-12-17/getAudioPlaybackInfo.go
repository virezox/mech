package main

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

var (
   LogLevel mech.LogLevel
   key = []byte("6#26FRL$ZWD")
)

func encrypt(src []byte) ([]byte, error) {
   for len(src) % blowfish.BlockSize >= 1 {
      src = append(src, 0)
   }
   dst := make([]byte, len(src))
   blow, err := blowfish.NewCipher(key)
   if err != nil {
      return nil, err
   }
   for low := 0; low < len(src); low += blow.BlockSize() {
      blow.Encrypt(dst[low:], src[low:])
   }
   return dst, nil
}

type playbackInfo struct {
   Result struct {
      AudioUrlMap struct {
         HighQuality struct {
            AudioURL string
         }
      }
   }
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("getAudioPlaybackInfo [userAuthToken]")
      return
   }
   userAuthToken := os.Args[1]
   dec := fmt.Sprintf(`
   {
      "deviceCode": "",
      "includeAudioToken": true,
      "pandoraId": "TR:1168891",
      "syncTime": 2222222222,
      "userAuthToken": "%v"
   }
   `, userAuthToken)
   enc, err := encrypt([]byte(dec))
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest(
      "POST",
      "http://android-tuner.pandora.com/services/json/",
      strings.NewReader(hex.EncodeToString(enc)),
   )
   if err != nil {
      panic(err)
   }
   val := make(mech.Values)
   // this can be empty, but it must be included:
   val["auth_token"] = ""
   val["method"] = "onDemand.getAudioPlaybackInfo"
   val["partner_id"] = "42"
   // this can be empty, but it must be included:
   val["user_id"] = ""
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(append(buf, '\n'))
   var info playbackInfo
   if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", info)
}
