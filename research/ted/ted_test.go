package ted

import (
   "os"
   "testing"
)

const slug = "rha_goddess_and_deepa_purushothaman_4_ways_to_redefine_power_at_work_to_include_women_of_color"

func TestSlug(t *testing.T) {
   res, err := get(slug)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
