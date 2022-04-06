package youtube

import (
   "testing"
)

func TestIosEmbed(t *testing.T) {
   const name = "IOS_EMBEDDED_PLAYER"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestIosKids(t *testing.T) {
   const name = "IOS_KIDS"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestIosLive(t *testing.T) {
   const name = "IOS_LIVE_CREATION_EXTENSION"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestIosMessage(t *testing.T) {
   const name = "IOS_MESSAGES_EXTENSION"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestIosMusic(t *testing.T) {
   const name = "IOS_MUSIC"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestIosProduce(t *testing.T) {
   const name = "IOS_PRODUCER"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestIosUnplug(t *testing.T) {
   const name = "IOS_UNPLUGGED"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestIosUptime(t *testing.T) {
   const name = "IOS_UPTIME"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}
