package youtube

import (
   "fmt"
   "testing"
)

type test struct {
   key string
   val func(*param)
}

var tests = []test{
   {"Channel", (*param).Channel},
   {"CreativeCommons", (*param).CreativeCommons},
   {"FourK", (*param).FourK},
   {"FourToTwentyMinutes", (*param).FourToTwentyMinutes},
   {"HD", (*param).HD},
   {"HDR", (*param).HDR},
   {"LastHour", (*param).LastHour},
   {"Live", (*param).Live},
   {"Location", (*param).Location},
   {"Movie", (*param).Movie},
   {"OverTwentyMinutes", (*param).OverTwentyMinutes},
   {"Playlist", (*param).Playlist},
   {"Purchased", (*param).Purchased},
   {"Rating", (*param).Rating},
   {"Relevance", (*param).Relevance},
   {"Subtitles", (*param).Subtitles},
   {"ThisMonth", (*param).ThisMonth},
   {"ThisWeek", (*param).ThisWeek},
   {"ThisYear", (*param).ThisYear},
   {"ThreeD", (*param).ThreeD},
   {"ThreeSixty", (*param).ThreeSixty},
   {"Today", (*param).Today},
   {"UnderFourMinutes", (*param).UnderFourMinutes},
   {"UploadDate", (*param).UploadDate},
   {"VR180", (*param).VR180},
   {"Video", (*param).Video},
   {"ViewCount", (*param).ViewCount},
}

func TestParam(t *testing.T) {
   for _, each := range tests {
      var p param
      each.val(&p)
      s, err := p.Encode()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(each.key)
      fmt.Print("youtube.com/results?search_query=hello&sp=", s, "\n\n")
   }
}
