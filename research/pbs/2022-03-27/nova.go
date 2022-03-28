package main

import (
   "encoding/json"
   "fmt"
   "os"
)

func main() {
   file, err := os.Open("ignore.json")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   var data nextData
   if err := json.NewDecoder(file).Decode(&data); err != nil {
      panic(err)
   }
   for _, episode := range data.Props.PageProps.IsSeries {
      if episode.Slug == "nova-universe-revealed-milky-way" {
         for _, asset := range episode.Episode.Assets {
            if asset.Object_Type == "full_length" {
               fmt.Printf("%a\n", asset)
            }
         }
      }
   }
}

type nextData struct {
   Props struct {
      PageProps struct {
         IsSeries []struct {
            Episode struct {
               Assets []Asset
            }
            Slug string
         }
      }
   }
}

type Asset struct {
   Object_Type string
   Slug string
   Player_Code string
}

func (a Asset) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Object_Type:", a.Object_Type)
   fmt.Fprint(f, "Slug: ", a.Slug)
   if verb == 'a' {
      fmt.Fprint(f, "\nPlayer_Code: ", a.Player_Code)
   }
}
