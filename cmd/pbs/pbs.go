package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech"
   "github.com/89z/mech/pbs"
   "net/http"
   "os"
)

func doAsset(address string, info bool) error {
   slug, err := pbs.Slug(address)
   if err != nil {
      return err
   }
   asset, err := pbs.NewAsset(slug)
   if err != nil {
      return err
   }
   if info {
      fmt.Printf("%+v\n", asset)
   } else {
      for _, video := range asset.Resource.MP4_Videos {
         fmt.Println("GET", video.URL)
         res, err := http.Get(video.URL)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusOK {
            return pbs.ErrorString(res.Status)
         }
         name := asset.Resource.Title + "-" + video.Profile + ".mp4"
         file, err := os.Create(mech.Clean(name))
         if err != nil {
            return err
         }
         defer file.Close()
         pro := format.NewProgress(res)
         if _, err := file.ReadFrom(pro); err != nil {
            return err
         }
      }
   }
   return nil
}
