package pandora

import (
   "encoding/hex"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "strings"
)

func (u userLogin) valueExchange() error {
   body := fmt.Sprintf(`
   {
      "offerName": "premium_access",
      "syncTime": 2222222222,
      "userAuthToken": "%v"
   }
   `, u.Result.UserAuthToken)
   enc, err := encrypt([]byte(body))
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "http://android-tuner.pandora.com/services/json/",
      strings.NewReader(hex.EncodeToString(enc)),
   )
   if err != nil {
      return err
   }
   val := make(mech.Values)
   // this can be empty, but must be included:
   val["auth_token"] = ""
   val["method"] = "user.startValueExchange"
   val["partner_id"] = "42"
   // this can be empty, but it must be included:
   val["user_id"] = ""
   req.URL.RawQuery = val.Encode()
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}
