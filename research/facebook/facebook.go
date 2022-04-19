package facebook

import (
   "github.com/headzoo/surf/browser"
   "net/url"
   "strings"
)

const (
	httpscheme       = "https"
	facebookhost     = "facebook.com"
	facebooklogin    = "/login.php"
	facebookchpasswd = "/settings/security/password/"
)

type facebookWebservice struct{}

func (f *facebookWebservice) login(browsr *browser.Browser, email, password string) error {
   fUrl := buildFBUrl(facebooklogin)
   err := browsr.Open(fUrl.String())
   if err != nil {
      return err
   }
   fm, err := browsr.Form("form[id=login_form]")
   if err != nil {
      return err
   }
   err = fm.Input("email", email)
   if err != nil {
      return err
   }
   err = fm.Input("pass", password)
   if err != nil {
      return err
   }
   if fm.Submit() != nil {
      return errorString("ParseError login_form_button " + fUrl.String())
   }
   if strings.Contains(browsr.Title(), "Log into Facebook") {
      return errorString("AccountError " + email + " " + facebookhost)
   }
   return nil
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

func buildFBUrl(path string) *url.URL {
	return &url.URL{
		Scheme: httpscheme,
		Host:   "m." + facebookhost,
		Path:   path,
	}
}
