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
      d, err := httputil.DumpRequest(req, true)
      if err != nil {
         return err
      }
      os.Stdout.Write(d)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(v)
}
