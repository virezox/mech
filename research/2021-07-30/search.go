package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
)

const returnal =
   "http://ia800306.us.archive.org/0/items" +
   "/mbid-a96ee369-9d38-4b13-a8c4-dab190519fc0" +
   "/mbid-a96ee369-9d38-4b13-a8c4-dab190519fc0-4753528871.jpg"

func main() {
   musicbrainz.Verbose = true
   other, err := musicbrainz.Hash(returnal)
   if err != nil {
      panic(err)
   }
   s, err := youtube.NewSearch("oneohtrix point never describing bodies")
   for _, i := range s.Items() {
      d, err := i.Distance(other)
      if err != nil {
         panic(err)
      }
      fmt.Println(d, i.TvMusicVideoRenderer)
   }
}
