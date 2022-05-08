package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   // <p id="report-account-vm" data-tou-report-params="
   "https://schnaussandmunk.bandcamp.com/music",
   // <p id="report-account-vm" data-tou-report-params="
   "https://schnaussandmunk.bandcamp.com",
   // <p id="report-account-vm" data-tou-report-params="
   "https://schnaussandmunk.bandcamp.com/album/passage-2",
   // <p id="report-account-vm" data-tou-report-params="
   "https://schnaussandmunk.bandcamp.com/track/amaris-2",
}

func TestData(t *testing.T) {
   for _, test := range tests {
      data, err := NewData(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", data)
      time.Sleep(99 * time.Millisecond)
   }
}
