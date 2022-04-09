package twitter

import (
   "os"
   "testing"
)

func TestTwitter(t *testing.T) {
   res, err := xauth()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
