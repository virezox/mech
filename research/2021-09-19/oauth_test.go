package github

import (
   "fmt"
   "testing"
)

func TestOAuth(t *testing.T) {
   o, err := NewOAuth()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Sign in to your GitHub Account
`, o.Get("verification_uri"), o.Get("user_code"))
}
