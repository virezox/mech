package youtube

import (
   "fmt"
   "testing"
)

func TestGoogleMedia(t *testing.T) {
   const (
      name = "GOOGLE_MEDIA_ACTIONS"
      version = "0.1"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestGoogleAssistant(t *testing.T) {
   const (
      name = "GOOGLE_ASSISTANT"
      version = "0.1"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}
