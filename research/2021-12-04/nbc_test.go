package nbc

import (
   "fmt"
   "github.com/89z/mech"
   "net/http/httputil"
   "os"
   "testing"
)

const id = 9000194212

func TestNBC(t *testing.T) {
   mech.Verbose = true
   res, err := media()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(buf)
   return
   vod, err := newAccessVOD(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vod)
}
