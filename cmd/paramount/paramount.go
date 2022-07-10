package main

import (
   "github.com/89z/mech/paramount"
)

func (f flags) do_DASH(preview *paramount.Preview) error {
   addr, err := paramount.New_Media(f.guid).DASH()
   if err != nil {
      return err
   }
   f.Poster, err = paramount.New_Session(f.guid)
   if err != nil {
      return err
   }
   reps, err := f.DASH(addr.String(), preview.Base())
   if err != nil {
      return err
   }
   return f.DASH_Get(reps, 0)
}

func (f flags) do_HLS(preview *paramount.Preview) error {
   addr, err := paramount.New_Media(f.guid).HLS()
   if err != nil {
      return err
   }
   master, err := f.HLS(addr.String(), preview.Base())
   if err != nil {
      return err
   }
   return f.HLS_Stream(master.Streams, 0)
}
