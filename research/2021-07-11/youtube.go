package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
)

var clients = []youtube.Client{
   {"ANDROID", "16.07.34"},
   {"ANDROID_CREATOR", "21.06.103"},
   {"ANDROID_EMBEDDED_PLAYER", "16.20"},
   {"ANDROID_KIDS", "6.02.3"},
   {"ANDROID_MUSIC", "4.32"},
   {"IOS", "16.05.7"},
   {"IOS_CREATOR", "20.47.100"},
   {"IOS_KIDS", "5.42.2"},
   {"IOS_MUSIC", "4.16.1"},
   {"MWEB", "2.19700101"},
   {"TVHTML5", "7.20210224.00.00"},
   {"WEB", "2.20210223.09.00"},
   {"WEB_CREATOR", "1.20210223.01.00"},
   {"WEB_EMBEDDED_PLAYER", "1.20210620.0.1"},
   {"WEB_KIDS", "2.1.3"},
   {"WEB_REMIX", "0.1"},
}

func pass(p *youtube.Player) bool {
   for _, f := range p.AdaptiveFormats {
      if f.Itag == 137 && f.URL != "" {
         return true
      }
   }
   return false
}

func main() {
   c := youtube.Client{"ANDROID", "16.05"}
   p, err := c.Player("54e6lBE3BoQ")
   if err != nil {
      panic(err)
   }
   for _, f := range p.AdaptiveFormats {
      fmt.Printf("%+v\n", f)
   }
   return
   for _, client := range clients {
      p, err := client.Player("54e6lBE3BoQ")
      if err != nil {
         fmt.Println(err, client)
         continue
      }
      if pass(p) {
         fmt.Println("pass", client)
      } else {
         fmt.Println("fail", client)
      }
   }
}

/*
pass {ANDROID 16.07.34}
pass {IOS 16.05.7}
*/
