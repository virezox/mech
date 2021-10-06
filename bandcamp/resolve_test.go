package bandcamp

import (
   "fmt"
   "net/http"
   "os"
   "testing"
)

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
