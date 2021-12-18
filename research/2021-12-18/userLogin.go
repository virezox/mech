package pandora

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "os"
   "strings"
)

type userLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

func main() {
   if len(os.Args) != 3 {
      fmt.Println("userLogin-user [password] [partnerAuthToken]")
      return
   }
   password, partnerAuthToken := os.Args[1], os.Args[2]
   body := fmt.Sprintf(`
   {
      "loginType": "user",
      "partnerAuthToken": "%v",
      "password": "%v",
      "syncTime": 2222222222,
      "username": "srpen6@gmail.com"
   }
   `, partnerAuthToken, password)
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
   val := make(mech.Values)
   // this can be empty, but must be included:
   val["auth_token"] = ""
   val["method"] = "auth.userLogin"
   val["partner_id"] = "42"
   req.URL.RawQuery = val.Encode()
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   var user userLogin
   if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", user)
   tLen := len(user.Result.UserAuthToken)
   if tLen != 58 {
      panic(tLen)
   }
}
