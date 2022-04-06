package youtube

import (
   "fmt"
   "testing"
)

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

func TestIosEmbed(t *testing.T) {
   const (
      name = "IOS_EMBEDDED_PLAYER"
      version = "2.0"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIosKids(t *testing.T) {
   const (
      name = "IOS_KIDS"
      version = "7.10.3"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIosLive(t *testing.T) {
   const (
      name = "IOS_LIVE_CREATION_EXTENSION"
      version = "17.11.34"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIosMessage(t *testing.T) {
   const (
      name = "IOS_MESSAGES_EXTENSION"
      version = "17.11.34"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIosMusic(t *testing.T) {
   const (
      name = "IOS_MUSIC"
      version = "4.70.50"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIosProduce(t *testing.T) {
   const (
      name = "IOS_PRODUCER"
      version = "0.1"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIosUnplug(t *testing.T) {
   const (
      name = "IOS_UNPLUGGED"
      version = "6.12.1"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestIosUptime(t *testing.T) {
   const (
      name = "IOS_UPTIME"
      version = "1.0"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}
