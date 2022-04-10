package twitter

import (
   "strings"
   "testing"
)

func TestSearch(t *testing.T) {
   search, err := NewSearch("filter:spaces")
   if err != nil {
      t.Fatal(err)
   }
   if len(search.GlobalObjects.Tweets) != 20 {
      t.Fatal(search)
   }
   for _, tweet := range search.GlobalObjects.Tweets {
      var ok bool
      for _, addr := range tweet.Entities.URLs {
         if strings.HasPrefix(addr.Expanded_URL, "https://twitter.com/i/spaces/") {
            ok = true
         }
      }
      if !ok {
         t.Fatal(tweet)
      }
   }
}
