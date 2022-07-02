package main

import (
   "flag"
   "github.com/89z/mech/amc"
   "os"
   "path/filepath"
)

func main() {
   f, err := new_flags()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.address, "a", "", "address")
   flag.Int64Var(&f.nid, "b", 0, "NID")
   flag.StringVar(&f.client_ID, "c", f.client_ID, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.IntVar(&f.bandwidth_video, "f", 1_999_999, "video bandwidth")
   flag.IntVar(&f.bandwidth_audio, "g", 127_000, "audio bandwidth")
   flag.BoolVar(&f.info, "i", false, "information")
   flag.StringVar(&f.private_key, "k", f.private_key, "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      amc.Client.Log_Level = 2
   }
   if f.email != "" {
      err := f.login()
      if err != nil {
         panic(err)
      }
   } else if f.nid >= 1 || f.address != "" {
      err := f.DASH()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) login() error {
   auth, err := amc.Unauth()
   if err != nil {
      return err
   }
   if err := auth.Login(f.email, f.password); err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return auth.Create(home + "/mech/amc.json")
}

type flags struct {
   address string
   bandwidth_audio int
   bandwidth_video int
   client_ID string
   email string
   info bool
   nid int64
   password string
   private_key string
   verbose bool
}

func new_flags() (*flags, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   var f flags
   f.client_ID = filepath.Join(home, "mech/client_id.bin")
   f.private_key = filepath.Join(home, "mech/private_key.pem")
   return &f, nil
}
