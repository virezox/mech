package bandcamp

import (
   "encoding/json"
   "net/http"
   "net/http/httputil"
   "os"
)

const Origin = "http://bandcamp.com"

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

var Verbose bool

func roundTrip(req *http.Request, v interface{}) error {
   if Verbose {
      dum, err := httputil.DumpRequest(req, true)
      if err != nil {
         return err
      }
      if dum[len(dum)-1] != '\n' {
         dum = append(dum, '\n')
      }
      os.Stdout.Write(dum)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(v)
}
