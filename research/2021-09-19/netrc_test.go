package github

import (
   "fmt"
   "net/http"
   "testing"
)

func TestNetrc(t *testing.T) {
   user, pass, err := netrc()
   if err != nil {
      t.Fatal(err)
   }
   req, err := http.NewRequest("HEAD", "https://api.github.com/rate_limit", nil)
   if err != nil {
      t.Fatal(err)
   }
   req.SetBasicAuth(user, pass)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", res)
}
