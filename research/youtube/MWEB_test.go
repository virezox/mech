package youtube

import (
   "testing"
   "time"
)

func TestMweb(t *testing.T) {
   const name = "MWEB"
   version, err := newVersion("https://m.youtube.com", "iPad")
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestPost(t *testing.T) {
   for name, version := range names {
      if version != "" {
         err := post(name, version)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
      }
   }
}
