package http

import (
   "errors"
   "net/http"
   "os"
   "testing/iotest"
)

func Two(s string) error {
   res, err := http.Get(s)
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
