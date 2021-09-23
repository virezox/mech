package instagram

import (
   "fmt"
   "testing"
   "time"
)

const id = "CT-cnxGhvvO"

func TestGraphQL(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := graphql(id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestP(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := p(id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
