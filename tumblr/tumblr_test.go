package tumblr

import (
   "fmt"
   "testing"
)

const post = "https://lyssafreyguy.tumblr.com/post/187741823636"

func TestTumblr(t *testing.T) {
   link, err := NewPermalink(post)
   if err != nil {
      t.Fatal(err)
   }
   post, err := link.BlogPost()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", post)
}
