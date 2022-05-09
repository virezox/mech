package ted

import (
   "fmt"
   "testing"
)

const slug = "rha_goddess_and_deepa_purushothaman_4_ways_to_redefine_power_at_work_to_include_women_of_color"

func TestSlug(t *testing.T) {
   talk, err := NewTalkResponse(slug)
   if err != nil {
      t.Fatal(err)
   }
   for _, vid := range talk.Downloads.Video {
      fmt.Printf("%a\n", vid)
   }
}
