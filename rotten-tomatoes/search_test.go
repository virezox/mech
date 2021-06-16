package tomato_test

import (
   "fmt"
   "github.com/89z/mech/rotten-tomatoes"
   "testing"
)

func TestSearch(t *testing.T) {
   s, err := tomato.NewSearch("inception")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", s)
}
