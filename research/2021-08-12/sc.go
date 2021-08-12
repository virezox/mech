package main

import (
   "fmt"
   "github.com/zackradisic/soundcloud-api"
)

func main() {
   var api *soundcloudapi.API
   dlURL, err := api.GetDownloadURL(
      "https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the",
      "hls",
   )
   if err != nil {
      panic(err)
   }
   fmt.Println(dlURL)
}
