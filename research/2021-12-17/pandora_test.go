package pandora

import (
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

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
   info, err := user.playbackInfo()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", info)
}

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(audioEnc)
   if err != nil {
      t.Fatal(err)
   }
   dec, err := decrypt(enc)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(dec)
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
