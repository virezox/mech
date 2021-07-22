package main

import (
   "fmt"
   "net/url"
)

const clientID =
   "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
   ".apps.googleusercontent.com"

func main() {
   val := url.Values{
      "client_id": {clientID},
      "redirect_uri": {"urn:ietf:wg:oauth:2.0:oob"},
      "response_type": {"code"},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }
   fmt.Printf(`1. Go to
https://accounts.google.com/o/oauth2/v2/auth?%v

2. Sign in to your Google Account
   `, val.Encode())
}

/*
curl -v `
-d client_id=861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com `
-d client_secret=SboVhoG9s0rNafixCSGGKXAT `
-d code=4/1AX4XfWgGbS8R7Fza-TojWaRuE0QFiN4asvmmc07VKlSjsH0ghn3Sm5... `
-d grant_type=authorization_code `
-d redirect_uri=urn:ietf:wg:oauth:2.0:oob `
https://oauth2.googleapis.com/token
*/
