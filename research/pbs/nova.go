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
   for _, series := range data.Props.PageProps.IsSeries {
      for _, asset := range series.Episode.Assets {
         fmt.Println(asset)
      }
   }
}

func (a Asset) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Object_Type:", a.Object_Type)
   fmt.Fprint(f, "Slug: ", a.Slug)
   if verb == 'a' {
      fmt.Fprint(f, "\nPlayer_Code: ", a.Player_Code)
   }
}

type Asset struct {
   Object_Type string
   Slug string
   Player_Code string
}

type nextData struct {
   Props struct {
      PageProps struct {
         IsSeries []struct {
            Episode struct {
               Assets []Asset
            }
         }
         Data struct {
            Episodes []struct {
               Episode struct {
                  Assets []struct {
                     Player_Code string
                  }
               }
            }
         }
         Video struct {
            Data struct {
               Episodes []struct {
                  Episode struct {
                     Assets []struct {
                        Player_Code string
                     }
                  }
               }
            }
         }
      }
   }
}
