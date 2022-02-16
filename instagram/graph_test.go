package instagram

import (
   "fmt"
   "testing"
   "os"
   "time"
)

var usernames = []string{
   "karajewelll",
   "lokalist.id",
}

func TestUser(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   login, err := OpenLogin(cache + "/mech/instagram.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, username := range usernames {
      user, err := login.User(username)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(user)
      time.Sleep(time.Second)
   }
}
