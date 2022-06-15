package http

import (
   "errors"
   "fmt"
   "net/http"
   "os"
   "time"
)

type bastard struct{}

func (bastard) Read([]byte) (int, error) {
   fmt.Println("bastard")
   time.Sleep(time.Second)
   return 0, nil
}

func (bastard) Close() error {
   return nil
}

func four(s string) error {
   req, err := http.NewRequest("GET", s, nil)
   if err != nil {
      return err
   }
   client := http.Client{Timeout: 9 * time.Second}
   res, err := client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   res.Body = bastard{}
   if _, err := os.Stdout.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
