package http

import (
   "errors"
   "net/http"
   "os"
   "testing/iotest"
   "time"
)

func Three(s string) error {
   req, err := http.NewRequest("GET", s, nil)
   if err != nil {
      return err
   }
   tr := http.Transport{
      ExpectContinueTimeout: 9 * time.Second,
      IdleConnTimeout:       9 * time.Second,
      TLSHandshakeTimeout:   9 * time.Second,
   }
   res, err := tr.RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   if _, err := os.Stdout.ReadFrom(iotest.ErrReader(nil)); err != nil {
      return err
   }
   return nil
}
