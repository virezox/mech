package imdb

import (
   "testing"
   "net/http/httputil"
   "os"
)

const rgconst = "rg2774637312"

func TestCred(t *testing.T) {
   cred, err := newCredentials()
   if err != nil {
      t.Fatal(err)
   }
   res, err := cred.Gallery(rgconst)
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
