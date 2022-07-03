package main

import (
   "github.com/89z/mech/paramount"
)

func (f flags) do_HLS(preview *paramount.Preview) error {
   addr, err := paramount.New_Media(f.guid).HLS()
   if err != nil {
      return err
   }
   f.Address = addr.String()
   return f.HLS(preview.Base())
}
