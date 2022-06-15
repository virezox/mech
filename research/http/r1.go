package http

import (
   "errors"
   "net/http"
   "os"
)

func One(s string) error {
   res, err := http.Get(s)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   if _, err := os.Stdout.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
