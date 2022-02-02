package instagram

import (
   "os"
   "testing"
)

func TestLogin(t *testing.T) {
   login, err := NewLogin("srpen6", password)
   if err != nil {
      t.Fatal(err)
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Create(cache + "/mech/instagram.json"); err != nil {
      t.Fatal(err)
   }
}
