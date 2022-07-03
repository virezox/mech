package main

import (
   "github.com/89z/mech/paramount"
)

func (f flags) do_DASH(preview *paramount.Preview) error {
   addr, err := paramount.New_Media(f.guid).DASH()
   if err != nil {
      return err
   }
   f.Address = addr.String()
   session, err := paramount.New_Session(f.guid)
   if err != nil {
      return err
   }
   return f.DASH(preview.Base(), session)
}
