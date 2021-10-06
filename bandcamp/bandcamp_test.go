package bandcamp

import (
   "fmt"
   "net/http"
   "os"
   "testing"
   "time"
)

type test struct {
   in string
   typ string
   id int
}

var tests = []test{
   {"https://schnaussandmunk.bandcamp.com/album/passage-2", "a", 1670971920},
   {"https://schnaussandmunk.bandcamp.com/track/amaris-2", "t", 2809477874},
}

var details = []Detail{
   {1, 79940049, "a"},
   {1, 2809477874, "t"},
}

func TestDetail(t *testing.T) {
   Verbose(true)
   for _, test := range tests {
      d, err := TralbumDetail(test.in)
      if err != nil {
         t.Fatal(err)
      }
      if d.Tralbum_Type != test.typ {
         t.Fatal(d.Tralbum_Type)
      }
      if d.Tralbum_ID != test.id {
         t.Fatal(d.Tralbum_ID)
      }
      time.Sleep(99 * time.Millisecond)
   }
}

func TestTralbum(t *testing.T) {
   Verbose(true)
   for _, detail := range details {
      tra, err := detail.Tralbum()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", tra)
   }
}

// 0.802s
func TestNew(t *testing.T) {
   req, err := http.NewRequest(
      "HEAD", "https://schnaussandmunk.bandcamp.com", nil,
   )
   if err != nil {
      t.Fatal(err)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(res.Cookies())
}

// 0.078s
func TestOld(t *testing.T) {
   req, err := http.NewRequest("GET", "http://bandcamp.com/api/url/2/info", nil)
   if err != nil {
      t.Fatal(err)
   }
   q := req.URL.Query()
   q.Set("key", "veidihundr")
   q.Set("url", "schnaussandmunk.bandcamp.com")
   req.URL.RawQuery = q.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
