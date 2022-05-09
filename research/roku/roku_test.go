package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestRoku(t *testing.T) {
   site, err := newCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(site.cookie)
   fmt.Println(site.token)
   res, err := site.playback()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
