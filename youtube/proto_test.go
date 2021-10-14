package youtube

import (
   "fmt"
   "testing"
)

type test struct {
   key string
   val func(*Param)
}

var tests = []test{
   {"Channel", (*Param).Channel},
   {"CreativeCommons", (*Param).CreativeCommons},
   {"FourK", (*Param).FourK},
   {"FourToTwentyMinutes", (*Param).FourToTwentyMinutes},
   {"HD", (*Param).HD},
   {"HDR", (*Param).HDR},
   {"LastHour", (*Param).LastHour},
   {"Live", (*Param).Live},
   {"Location", (*Param).Location},
   {"Movie", (*Param).Movie},
   {"OverTwentyMinutes", (*Param).OverTwentyMinutes},
   {"Playlist", (*Param).Playlist},
   {"Purchased", (*Param).Purchased},
   {"Rating", (*Param).Rating},
   {"Relevance", (*Param).Relevance},
   {"Subtitles", (*Param).Subtitles},
   {"ThisMonth", (*Param).ThisMonth},
   {"ThisWeek", (*Param).ThisWeek},
   {"ThisYear", (*Param).ThisYear},
   {"ThreeD", (*Param).ThreeD},
   {"ThreeSixty", (*Param).ThreeSixty},
   {"Today", (*Param).Today},
   {"UnderFourMinutes", (*Param).UnderFourMinutes},
   {"UploadDate", (*Param).UploadDate},
   {"VR180", (*Param).VR180},
   {"Video", (*Param).Video},
   {"ViewCount", (*Param).ViewCount},
}

func TestParam(t *testing.T) {
   for _, each := range tests {
      var p Param
      each.val(&p)
      s, err := p.Encode()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(each.key)
      fmt.Print("youtube.com/results?search_query=hello&sp=", s, "\n\n")
   }
}
