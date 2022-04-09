package twitter

import (
   "net/http/httputil"
   "os"
   "testing"
)

func TestTwitter(t *testing.T) {
   LogLevel = 1
   guest, err := NewGuest()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := guest.xauth(identifier, password)
   if err != nil {
      t.Fatal(err)
   }
   res, err := auth.search()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(buf)
}
