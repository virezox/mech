package nbc

import (
   "github.com/89z/mech"
   "net/http/httputil"
   "os"
   "testing"
)

const (
   // nbc.com/la-brea/video/pilot/9000194212
   res540 = 9000194212
   // nbc.com/the-blacklist/video/the-skinner/9000210182
   res1080 = 9000210182
)

func TestWeb(t *testing.T) {
   mech.Verbose = true
   res, err := media(res540)
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
