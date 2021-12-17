package pandora

import (
   "bytes"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(login2)
   if err != nil {
      t.Fatal(err)
   }
   dec, err := decrypt(enc)
   if err != nil {
      t.Fatal(err)
   }
   buf := new(bytes.Buffer)
   json.Indent(buf, dec, "", " ")
   os.Stdout.ReadFrom(buf)
}

func TestLogin(t *testing.T) {
   LogLevel = 1
   part, err := newPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", part)
   user, err := part.userLogin("srpen6@gmail.com", password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", user)
   tLen := len(user.Result.UserAuthToken)
   if tLen != 58 {
      t.Fatal("userAuthToken", tLen)
   }
   info, err := new(userLogin).playbackInfo()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", info)
}

func TestEncrypt(t *testing.T) {
   buf := []byte(`{"hello":"world"}`)
   enc, err := encrypt(buf)
   if err != nil {
      t.Fatal(err)
   }
   str := hex.EncodeToString(enc)
   fmt.Println(str)
}
