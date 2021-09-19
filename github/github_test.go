package github

import (
   "fmt"
   "testing"
)

func TestGitHub(t *testing.T) {
   Verbose = true
   rs, err := NewRepoSearch("stars:>999 pushed:>2020-09-19", "1")
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", rs)
}
