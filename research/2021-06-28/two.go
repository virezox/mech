package main

var clients = []string{
   "ANDROID",
   "ANDROID_CREATOR",
   "ANDROID_EMBEDDED_PLAYER",
   "ANDROID_KIDS",
   "ANDROID_MUSIC",
   "IOS",
   "IOS_CREATOR",
   "IOS_KIDS",
   "IOS_MUSIC",
   "MWEB",
   "TVHTML5",
   "WEB",
   "WEB_CREATOR",
   "WEB_KIDS",
   "WEB_REMIX",
}

func main() {
   for _, one := range clients {
      for _, two := range clients {
         println(one, two)
      }
   }
}
