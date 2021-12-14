package twitter

import (
   "fmt"
   "testing"
)

const id = 1470124083547418624

func TestTwitter(t *testing.T) {
   act, err := NewActivate()
   if err != nil {
      t.Fatal(err)
   }
   stat, err := act.Status(id)
   if err != nil {
      t.Fatal(err)
   }
   for _, med := range stat.Extended_Entities.Media {
      for _, variant := range med.Video_Info.Variants {
         fmt.Printf("%+v\n", variant)
      }
   }
}
