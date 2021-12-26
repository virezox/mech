package twitter

import (
   "fmt"
   "testing"
)

const spaceID = "1OdKrBnaEPXKX"

func TestSpace(t *testing.T) {
   space, err := NewSpace(guest, spaceID)
   if err != nil {
      t.Fatal(err)
   }
   stream, err := space.Stream(guest)
   if err != nil {
      t.Fatal(err)
   }
   chunks, err := stream.Chunks()
   if err != nil {
      t.Fatal(err)
   }
   for _, chunk := range chunks {
      fmt.Println(chunk)
   }
}
