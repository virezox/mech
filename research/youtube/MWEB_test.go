package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestPrint(t *testing.T) {
   for name, version := range names {
      fmt.Println(name + ";" + version)
   }
}

func TestPlayer(t *testing.T) {
   for name, version := range names {
      if version != "" {
         play, err := newPlayer(name, version)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(play.PlayabilityStatus.Status, name)
         time.Sleep(time.Second)
      }
   }
}

func TestMweb(t *testing.T) {
   const name = "MWEB"
   logLevel = 1
   version, err := newVersion("https://m.youtube.com", "iPad")
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}
