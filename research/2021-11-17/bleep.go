package main

import (
   "github.com/89z/parse/html"
)

func main() {
   file, err := os.Open("bleep.html")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   lex := html.NewLexer(file)
   // get date
   // <meta property="music:release_date" content="2001-05-01 00:00:00.0">
   lex.NextAttr("property", "music:release_date")
   // get image
   // get title
}
