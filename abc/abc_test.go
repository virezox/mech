package abc

import (
   "fmt"
   "testing"
)

const grey = "https://abc.com" +
   "/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you"

func TestMech(t *testing.T) {
   route, err := NewRoute(grey)
   if err != nil {
      t.Fatal(err)
   }
   video, err := route.Video()
   if err != nil {
      t.Fatal(err)
   }
   if err := video.Authorize(); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", video)
}
