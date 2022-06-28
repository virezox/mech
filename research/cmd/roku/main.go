package main

import (
   "flag"
   "github.com/89z/format/dash"
   "github.com/89z/mech/roku"
   "net/url"
   "os"
   "path/filepath"
)

type downloader struct {
   *roku.Content
   client string
   info bool
   key []byte
   pem string
   url *url.URL
   media dash.Media
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var down downloader
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var id string
   flag.StringVar(&id, "b", "", "ID")
   // c
   down.client = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&down.client, "c", down.client, "client ID")
   // d
   var is_dash bool
   flag.BoolVar(&is_dash, "d", false, "DASH download")
   // f
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   var video int64
   flag.Int64Var(&video, "f", 1920832, "video bandwidth")
   // g
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   var audio int64
   flag.Int64Var(&audio, "g", 128000, "audio bandwidth")
   // i
   flag.BoolVar(&down.info, "i", false, "information")
   // k
   down.pem = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&down.pem, "k", down.pem, "private key")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      roku.Client.Log_Level = 2
   }
   if id != "" || address != "" {
      if id == "" {
         id = roku.Content_ID(address)
      }
      down.Content, err = roku.New_Content(id)
      if err != nil {
         panic(err)
      }
      if is_dash {
         err := down.DASH(video, audio)
         if err != nil {
            panic(err)
         }
      } else {
         err := down.HLS(video)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
