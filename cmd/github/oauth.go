package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/github"
   "os"
)

func authConstruct() (*youtube.Exchange, error) {
   cac, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }
   fil, err := os.Open(cac + "/mech/youtube.json")
   if err != nil {
      return nil, err
   }
   defer fil.Close()
   exc := new(youtube.Exchange)
   if err := json.NewDecoder(fil).Decode(exc); err != nil {
      return nil, err
   }
   return exc, nil
}

func authExchange() error {
   oau, err := github.NewOAuth()
   if err != nil {
      return err
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Press Enter to continue`, oau.Verification_URI, oau.User_Code)
   fmt.Scanln()
   exc, err := oau.Exchange()
   if err != nil {
      return err
   }
   cac, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cac += "/mech"
   os.Mkdir(cac, os.ModeDir)
   fil, err := os.Create(cac + "/github.json")
   if err != nil {
      return err
   }
   defer fil.Close()
   enc := json.NewEncoder(fil)
   enc.SetIndent("", " ")
   return enc.Encode(exc)
}
