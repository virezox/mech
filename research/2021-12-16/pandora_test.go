package pandora

import (
   "bytes"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func TestInfo(t *testing.T) {
   LogLevel = 1
   info, err := newPlaybackInfo()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", info)
}

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(assoc)
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

func TestEncrypt(t *testing.T) {
   buf := []byte(`{"hello":"world"}`)
   enc, err := encrypt(buf)
   if err != nil {
      t.Fatal(err)
   }
   str := hex.EncodeToString(enc)
   fmt.Println(str)
}
