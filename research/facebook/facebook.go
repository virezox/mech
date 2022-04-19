package facebook

import (
   "github.com/89z/format"
   "github.com/headzoo/surf"
   "net/http"
   "strings"
)

type errorString string

func (e errorString) Error() string {
   return string(e)
}

type verbose struct{}

func (verbose) RoundTrip(req *http.Request) (*http.Response, error) {
   format.LogLevel(1).Dump(req)
   return new(http.Transport).RoundTrip(req)
}

type facebookWebservice struct{}

const endpoint = "https://m.facebook.com/login.php"

func (f *facebookWebservice) login(email, password string) error {
   bro := surf.NewBrowser()
   var ver verbose
   bro.SetTransport(ver)
   err := bro.Open(endpoint)
   if err != nil {
      return err
   }
   fm, err := bro.Form("form[id=login_form]")
   if err != nil {
      return err
   }
   if err := fm.Input("email", email); err != nil {
      return err
   }
   if err := fm.Input("pass", password); err != nil {
      return err
   }
   if fm.Submit() != nil {
      return errorString("ParseError login_form_button " + endpoint)
   }
   if strings.Contains(bro.Title(), "Log into Facebook") {
      return errorString("AccountError " + email + " " + endpoint)
   }
   return nil
}
