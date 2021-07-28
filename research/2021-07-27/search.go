package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
)

func main() {
   s, err := youtube.NewSearch("oneohtrix point never along")
   if err != nil {
      panic(err)
   }
   for _, vid := range s.Videos() {
      fmt.Println(vid.VideoID)
   }
}
