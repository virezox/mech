package youtube

import (
   "fmt"
   "testing"
)

func TestIosCreator(t *testing.T) {
   const (
      name = "IOS_CREATOR"
      version = "22.11.100"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIos(t *testing.T) {
   const (
      name = "IOS"
      version = "17.11.34"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}
