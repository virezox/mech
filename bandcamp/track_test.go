package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const trackID = "2809477874"

func TestTrack(t *testing.T) {
   Verbose(true)
   tra := new(Track)
   if err := tra.Get(trackID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
   time.Sleep(100 * time.Millisecond)
   if err := tra.Post(trackID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
   time.Sleep(100 * time.Millisecond)
   if err := tra.PostForm(trackID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
}
