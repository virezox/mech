package bandcamp

import (
   "fmt"
   "net/http"
   "os"
   "testing"
)

const (
   addr = "https://schnaussandmunk.bandcamp.com/track/amaris-2"
   id = 2809477874
)

func TestTrack(t *testing.T) {
   Verbose(true)
   inf, err := NewInfo(addr)
   if err != nil {
      t.Fatal(err)
   }
   tra, err := NewTrack(inf.Track_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
}

// 0.095s
func TestInfo(t *testing.T) {
   req, err := http.NewRequest("GET", ApiUrl, nil)
   if err != nil {
      t.Fatal(err)
   }
   q := req.URL.Query()
   q.Set("key", key)
   q.Set("url", addr)
   req.URL.RawQuery = q.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}

// 0.405s
func TestHead(t *testing.T) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      t.Fatal(err)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(res.Cookies())
}
