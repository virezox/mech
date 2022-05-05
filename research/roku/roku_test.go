package roku

import (
   "testing"
)

func TestRoku(t *testing.T) {
   con, err := NewContent()
   if err != nil {
      t.Fatal(err)
   }
   if con.EpisodeNumber == "" {
      t.Fatal(con)
   }
   if con.ReleaseDate == "" {
      t.Fatal(con)
   }
   if con.RunTimeSeconds == 0 {
      t.Fatal(con)
   }
   if con.SeasonNumber == "" {
      t.Fatal(con)
   }
   if con.Series.Title == "" {
      t.Fatal(con)
   }
   if con.Title == "" {
      t.Fatal(con)
   }
   if len(con.ViewOptions) == 0 {
      t.Fatal(con)
   }
   if len(con.ViewOptions[0].Media.Videos) == 0 {
      t.Fatal(con)
   }
   if con.ViewOptions[0].Media.Videos[0].URL == "" {
      t.Fatal(con)
   }
}
