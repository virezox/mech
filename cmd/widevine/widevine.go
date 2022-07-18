package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/widevine"
   "github.com/89z/rosso/http"
   "os"
   "path/filepath"
)

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // a
   flag.StringVar(&f.address, "a", "", "license address")
   // b
   flag.StringVar(&f.key_ID, "b", "", "key ID")
   // c
   f.client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.client_ID, "c", f.client_ID, "client ID")
   // h
   flag.StringVar(&f.header, "h", "", "header")
   // k
   f.private_key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.private_key, "k", f.private_key, "private key")
   // v
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      widevine.Client.Log_Level = 2
   }
   if f.key_ID != "" {
      contents, err := f.contents()
      if err != nil {
         panic(err)
      }
      fmt.Println(contents.Content())
   } else {
      flag.Usage()
   }
}

var http_client = http.Default_Client

func (f flags) contents() (widevine.Containers, error) {
   private_key, err := os.ReadFile(f.private_key)
   if err != nil {
      return nil, err
   }
   client_ID, err := os.ReadFile(f.client_ID)
   if err != nil {
      return nil, err
   }
   key_ID, err := widevine.Key_ID(f.key_ID)
   if err != nil {
      return nil, err
   }
   module, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      return nil, err
   }
   return module.Post(f)
}
