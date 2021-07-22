package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
   "time"
)

func TestOAuth(t *testing.T) {
   o, err := youtube.NewOAuth()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Sign in to your Google Account
`, o.Verification_URL, o.User_Code)
   var a youtube.Auth
   for range [9]struct{}{} {
      time.Sleep(9 * time.Second)
      x, err := o.Exchange()
      if err != nil {
         t.Fatal(err)
      }
      if x.Access_Token != "" {
         a = youtube.Auth{"Authorization", "Bearer " + x.Access_Token}
         break
      }
   }
   p, err := youtube.NewPlayer("Cr381pDsSsA", a, youtube.Android)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", p)
}
