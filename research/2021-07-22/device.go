package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
)

func main() {
   a, err := youtube.NewAuth()
   if err != nil {
      panic(err)
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Sign in to your Google Account

4. Press Enter to continue`, a.Verification_URL, a.User_Code)
   fmt.Scanln()
   x, err := a.Exchange()
   if err != nil {
      panic(err)
   }
   // at this point we would want to save the refresh_token
   fmt.Printf("%+v\n", x)
   if err := x.Refresh(); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", x)
}
