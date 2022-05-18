package bearer

import (
   "os"
   "testing"
)

func TestBearer(t *testing.T) {
   res, err := newBearer()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
