package main

import (
   "fmt"
   "github.com/89z/mech/paramount"
   "net/http"
)

func doDASH(guid string, bandwidth int64, info bool) error {
   addr, err := paramount.NewMedia(guid).DASH()
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return nil
}
