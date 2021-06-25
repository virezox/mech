package deezer

import (
   "io"
   "net/http"
   "testing"
)

const sngID = "75498418"
const arl = "410b125e6b9d2600d2905364752a441972619cfc03ce3f745282c333ab3075cdbb2c2ca8602deb79611c50a0aa50fe026c79539b1cdd5403c4a8fb3368e63abdb68517259be7329bfb264f50a20f22f47c902c174098405325efb8b5db49ab9c"

func TestDecrypt(t *testing.T) {
   user, err := NewUserData(arl)
   if err != nil {
      t.Fatal(err)
   }
   track, err := NewTrack(sngID, user.Results.CheckForm, user.SID)
   if err != nil {
      t.Fatal(err)
   }
   source, err := track.Source(sngID, MP3_320)
   if err != nil {
      t.Fatal(err)
   }
   res, err := http.Get(source)
   if err != nil {
      t.Fatal(err)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      t.Fatal(err)
   }
   Decrypt(sngID, body)
   testHash := md5Hash(string(body))
   if testHash != "87207d3416377217f835b887c74f4300" {
      t.Fatal(testHash)
   }
}
