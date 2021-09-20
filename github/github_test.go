package github

import (
   "fmt"
   "testing"
)

func TestGitHub(t *testing.T) {
   Verbose = true
   s := NewSearch("stars:>999 pushed:>2020-09-19")
   s.Page(2)
   s.PerPage(9)
   r, err := s.Repos(nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", r)
}
