package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/amc"
   "os"
   "path/filepath"
)

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.Int64Var(&f.nid, "b", 0, "NID")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // e
   flag.StringVar(&f.email, "e", "", "email")
   // f
   flag.IntVar(&f.Video_Bandwidth, "f", 1_999_999, "video bandwidth")
   // g
   flag.IntVar(&f.Audio_Bandwidth, "g", 127_000, "audio bandwidth")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
   // p
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   if f.email != "" {
      err := f.login()
      if err != nil {
         panic(err)
      }
   } else if f.nid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.Open_Auth(home + "/mech/amc.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Create(home + "/mech/amc.json"); err != nil {
      return err
   }
   play, err := auth.Playback(f.nid)
   if err != nil {
      return err
   }
   f.Address = play.Source().Src
   f.Base = play.Base()
   f.Poster = play
   return f.DASH()
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
