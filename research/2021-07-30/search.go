package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/corona10/goimagehash"
)

/*
http://ia800709.us.archive.org/34/items/
mbid-10cc746f-786c-4307-b8de-92a687489cb4/
mbid-10cc746f-786c-4307-b8de-92a687489cb4-4958564206.jpg
*/
var other = goimagehash.NewImageHash(16638239206888408964, goimagehash.DHash)

func main() {
   s, err := youtube.NewSearch("oneohtrix point never along")
   if err != nil {
      panic(err)
   }
   for _, i := range s.Items() {
      d, err := i.Distance(other)
      if err != nil {
         panic(err)
      }
      fmt.Println(d, i.TvMusicVideoRenderer)
   }
}
