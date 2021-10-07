package main

import (
   "github.com/89z/mech/googleplay"
   "net/url"
   "os"
)

func main() {
   txt, err := os.ReadFile("ac2dm.txt")
   if err != nil {
      panic(err)
   }
   val, err := url.ParseQuery(string(txt))
   if err != nil {
      panic(err)
   }
   ac2 := googleplay.Ac2dm{val}
   googleplay.Verbose(true)
   auth, err := ac2.OAuth2()
   if err != nil {
      panic(err)
   }
   data, err := auth.Details("38B5418D8683ADBB", "com.google.android.youtube")
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(data)
}
