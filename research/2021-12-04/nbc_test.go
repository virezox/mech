package nbc

import (
   "fmt"
   "github.com/89z/mech"
   "net/http/httputil"
   "os"
   "testing"
)

// POST /access/vod/nbcuniversal/9000194212 HTTP/1.1

const (
   // nbc.com/la-brea/video/pilot/9000194212
   res540 = 9000194212
   // nbc.com/the-blacklist/video/the-skinner/9000210182
   res1080 = 9000210182
)

func TestVOD(t *testing.T) {
   mech.Verbose = true
   vod, err := newAccessVOD(res540)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vod)
}

func TestMedia(t *testing.T) {
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
}
