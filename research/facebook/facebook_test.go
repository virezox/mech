package facebook

import (
   "testing"
)

func TestFBLogsInSuccessfully(t *testing.T) {
   var web facebookWebservice
   err := web.login(email, password)
   if err != nil {
      t.Fatal(err)
   }
}
