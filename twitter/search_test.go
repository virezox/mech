package twitter

import (
   "fmt"
   "strings"
   "testing"
   "time"
)

const prefix = "https://twitter.com/i/spaces/"

func TestSearch(t *testing.T) {
   // until:2022-04-13
   search, err := NewSearch("filter:spaces")
   if err != nil {
      t.Fatal(err)
   }
   if len(search.GlobalObjects.Tweets) != 20 {
      t.Fatal(search)
   }
   guest, err := NewGuest()
   if err != nil {
      t.Fatal(err)
   }
   for _, tweet := range search.GlobalObjects.Tweets {
      for _, addr := range tweet.Entities.URLs {
         if strings.HasPrefix(addr.Expanded_URL, prefix) {
            id, err := SpaceID(addr.Expanded_URL)
            if err != nil {
               t.Fatal(err)
            }
            space, err := guest.AudioSpace(id)
            if err != nil {
               t.Fatal(err)
            }
            dur := space.Duration()
            if dur >= 9*time.Minute && dur <= 99*time.Minute {
               fmt.Println(tweet)
               return
            }
            time.Sleep(99 * time.Millisecond)
         }
      }
   }
}
