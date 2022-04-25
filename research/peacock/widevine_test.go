package widevine

import (
   "os"
   "testing"
)

const pssh = "AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAWylyu1Ra6Obew32huzP9I49yVmwY="

func TestWidevine(t *testing.T) {
   file, err := os.Open("2.txt")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   req, err := newRequest(pssh, file)
   if err != nil {
      t.Fatal(err)
   }
   LogLevel = 1
   res, err := req.post()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
