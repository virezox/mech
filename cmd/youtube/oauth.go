package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

// write auth to file
func authWrite() error {
   o, err := youtube.NewOAuth()
   if err != nil {
      return err
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Sign in to your Google Account

4. Press Enter to continue`, o.Verification_URL, o.User_Code)
   fmt.Scanln()
   x, err := o.Exchange()
   if err != nil {
      return err
   }
   c, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   c += "/mech"
   os.Mkdir(c, os.ModeDir)
   f, err := os.Create(c + "/youtube.json")
   if err != nil {
      return err
   }
   defer f.Close()
   e := json.NewEncoder(f)
   e.SetIndent("", " ")
   return e.Encode(x)
}

// read auth from file
func authRead() (*youtube.Exchange, error) {
   c, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }
   f, err := os.Open(c + "/mech/youtube.json")
   if err != nil {
      return nil, err
   }
   defer f.Close()
   x := new(youtube.Exchange)
   if err := json.NewDecoder(f).Decode(x); err != nil {
      return nil, err
   }
   return x, nil
}
