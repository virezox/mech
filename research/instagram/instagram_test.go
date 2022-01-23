package instagram

import (
   "net/http/httputil"
   "os"
   "testing"
)

func TestInstagram(t *testing.T) {
   res, err := post()
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
