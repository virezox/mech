package paramount

import (
   "os"
   "testing"
)

func TestParamount(t *testing.T) {
   res, err := master()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
