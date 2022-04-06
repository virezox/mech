package youtube

import (
   "testing"
)

func TestGoogleAssistant(t *testing.T) {
   const name = "GOOGLE_ASSISTANT"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}

func TestGoogleMedia(t *testing.T) {
   const name = "GOOGLE_MEDIA_ACTIONS"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}
