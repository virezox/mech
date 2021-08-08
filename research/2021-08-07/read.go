package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "os"
)

func main() {
   f, err := os.Open("returnal.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   var tracks []struct {
      musicbrainz.Track
      Items []struct {
         youtube.Item
      }
   }
   json.NewDecoder(f).Decode(&tracks)
   for _, t := range tracks {
      // track index
      indexes := []float64{0}
      // track length
      lengths := []float64{t.Length}
      var index float64
      for _, i := range t.Items {
         // item index
         indexes = append(indexes, index)
         // item length
         d, err := i.Duration()
         if err != nil {
            panic(err)
         }
         lengths = append(lengths, d.Seconds() * 1000)
         index++
      }
      fmt.Print(indexes, "\n", lengths, "\n\n")
   }
}
