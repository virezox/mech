package vimeo

import (
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/json"
   "io"
)

var Client = http.Default_Client

type Embed struct {
   API_Data struct {
      API_Token string
   }
   Config_URL string
}

func New_Embed(ref string) (*Embed, error) {
   res, err := Client.Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var scan json.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte(".OTTData =")
   scan.Scan()
   scan.Sep = nil
   emb := new(Embed)
   if err := scan.Decode(emb); err != nil {
      return nil, err
   }
   return emb, nil
}
