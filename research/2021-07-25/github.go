package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "time"
)

type repo struct {
   Pushed_At string
   Size int
   Stargazers_Count int
}

var addrs = []string{
   "Nr90/imgsim",
   "ajdnik/imghash",
   "andybalholm/dhash",
   "atijust/ahash",
   "corona10/goimagehash",
   "devedge/imagehash",
   "g3vxy/dhash",
   "myusuf3/imghash",
   "olegfedoseev/image-diff",
   "teran/imgsum",
   "umahmood/perceptive",
}

func main() {
   for _, addr := range addrs {
      res, err := http.Get("https://api.github.com/repos/" + addr)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      var r repo
      json.NewDecoder(res.Body).Decode(&r)
      fmt.Println(r.Size, r.Stargazers_Count, r.Pushed_At, addr)
      time.Sleep(time.Second)
   }
}
