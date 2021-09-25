package instagram

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
)

func htmlEmbed(id string) error {
   req, err := http.NewRequest("GET", origin + "/p/" + id + "/embed/", nil)
   if err != nil {
      return err
   }
   req.Header.Set("User-Agent", "Mozilla")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

func htmlP(id string) error {
   req, err := http.NewRequest("GET", origin + "/p/" + id + "/", nil)
   if err != nil {
      return err
   }
   req.Header.Set("User-Agent", "Mozilla")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}
