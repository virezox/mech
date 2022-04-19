package facebook

import (
   "github.com/89z/format"
   "github.com/headzoo/surf"
   "net/http"
   "testing"
)

type verbose struct{}

func (verbose) RoundTrip(req *http.Request) (*http.Response, error) {
   format.LogLevel(1).Dump(req)
   return new(http.Transport).RoundTrip(req)
}

func TestFBLogsInSuccessfully(t *testing.T) {
   bro := surf.NewBrowser()
   var ver verbose
   bro.SetTransport(ver)
   var web facebookWebservice
   err := web.login(bro, email, password)
   if err != nil {
      t.Fatal(err)
   }
}
