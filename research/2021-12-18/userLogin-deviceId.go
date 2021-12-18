package main

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "os"
   "strings"
)

var (
   LogLevel mech.LogLevel
   key = []byte("6#26FRL$ZWD")
)

type userLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

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

func main() {
   if len(os.Args) != 2 {
      fmt.Println("userLogin-deviceId [partnerAuthToken]")
      return
   }
   partnerAuthToken := os.Args[1]
   deviceID := "7db1bef0-1ea2-4ba5-8c0a-079274c81b75"
   body := fmt.Sprintf(`
   {
      "deviceId": "%v",
      "loginType": "deviceId",
      "partnerAuthToken": "%v",
      "syncTime": 2222222222
   }
   `, deviceID, partnerAuthToken)
   enc, err := encrypt([]byte(body))
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
   req.Header.Set("User-Agent", "Pandora/2110.1 Android/7.0 generic_x86")
   val := make(mech.Values)
   val["auth_token"] = ""
   val["method"] = "auth.userLogin"
   val["partner_id"] = "42"
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   var user userLogin
   if err := json.NewDecoder(res.Body).Decode(user); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", user)
   tLen := len(user.Result.UserAuthToken)
   if tLen != 58 {
      panic(tLen)
   }
}
