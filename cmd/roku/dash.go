package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/roku"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
)

func (f flags) do_DASH(content *roku.Content) error {
   f.Address = content.DASH().URL
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   play, err := site.Playback(f.id)
   if err != nil {
      return err
   }
   return f.DASH(content.Base(), play)
}
