package vimeo

import (
   "net/http"
   "net/http/httputil"
   "os"
   "testing"
)

func TestVimeo(t *testing.T) {
   req, err := newRequest("66531465")
   if err != nil {
      t.Fatal(err)
   }
   d, err := httputil.DumpRequest(req, true)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      t.Fatal(res)
   }
}
