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

func TestNBC(t *testing.T) {
   mech.LogLevel = 3
   med, err := newMedia(guid)
   if err != nil {
      t.Fatal(err)
   }
   forms, err := med.video()
   if err != nil {
      t.Fatal(err)
   }
   form, ok := forms.Get(0)
   if ok {
      fmt.Println("GET", form.URI)
      res, err := http.Get(form.URI.String())
      if err != nil {
         t.Fatal(err)
      }
      defer res.Body.Close()
      forms, err := m3u.Decode(res.Body, "")
      if err != nil {
         t.Fatal(err)
      }
      form, ok := forms.Get(1)
      if ok {
         fmt.Println("GET", form.URI)
      }
   }
}
