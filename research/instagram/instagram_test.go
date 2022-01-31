package instagram

import (
   "fmt"
   "testing"
)

const (
   id = 2762134734241678695
   shortcode = "CZVEugIPkVn"
)

func TestID(t *testing.T) {
   id, err := getID(shortcode)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(id)
}

func TestMedia(t *testing.T) {
   med, err := newMedia(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", med)
}
