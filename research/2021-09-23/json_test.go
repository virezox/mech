package instagram

import (
   "fmt"
   "testing"
   "time"
)

func TestJsonGraphQL(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonGraphQL(id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestJsonP(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonP(id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestJsonTV(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonTV(id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
