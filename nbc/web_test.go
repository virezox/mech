package nbc

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "testing"
)

// nbc.com/la-brea/video/pilot/9000194212
const guid = 9000194212

func TestWeb(t *testing.T) {
   mech.LogLevel = 3
   med, err := newMedia(guid)
   if err != nil {
      t.Fatal(err)
   }
   forms, err := med.video()
   if err != nil {
      t.Fatal(err)
   }
   form := forms[0]
   fmt.Println("GET", form)
   res, err := http.Get(form["URI"])
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   files, err := m3u.Decode(res.Body, "")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(files)
}
