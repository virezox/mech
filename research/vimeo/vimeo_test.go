package vimeo

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestVideo(t *testing.T) {
   web, err := newJsonWeb()
   if err != nil {
      t.Fatal(err)
   }
   videos := []string{"/videos/581039021:9603038895", "/videos/660408476"}
   for _, video := range videos {
      res, err := web.video(video)
      if err != nil {
         t.Fatal(err)
      }
      defer res.Body.Close()
      os.Stdout.ReadFrom(res.Body)
      fmt.Print("\n---\n")
      time.Sleep(time.Second)
   }
}
