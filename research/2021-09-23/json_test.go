package instagram

import (
   "fmt"
   "testing"
   "time"
)

func TestJsonChannel(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonChannel(name)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestJsonGraphQL(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonGraphQL(shortcode)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestJsonP(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonP(shortcode)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestJsonTV(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonTV(shortcode)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestJsonUsers(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := jsonUsers(user)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
