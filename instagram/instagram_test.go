package instagram

import (
   "fmt"
   "testing"
)

const like = "CUrAS88Pr1G"

func TestLike(t *testing.T) {
   med, err := NewMedia(like)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", med)
}

func TestWrite(t *testing.T) {
   login, err := NewLogin("srpen6", password)
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Create("instagram.json"); err != nil {
      t.Fatal(err)
   }
}
