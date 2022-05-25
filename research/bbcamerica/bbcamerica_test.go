package bbcamerica

import (
   "os"
   "testing"
)

func TestUnauth(t *testing.T) {
   res, err := unauth()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
