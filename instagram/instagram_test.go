package instagram

import (
   "fmt"
   "os"
   "testing"
)

const like = "CUrAS88Pr1G"

func TestLike(t *testing.T) {
   graph, err := Login{}.GraphQL(like)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", graph)
}

func TestWrite(t *testing.T) {
   login, err := NewLogin("srpen6", password)
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Encode(os.Stdout); err != nil {
      t.Fatal(err)
   }
}
