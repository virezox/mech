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
   "andybalholm/go-bit",
   "corona10/goimagehash",
   "disintegration/imaging",
   "golang/image",
   "mjibson/go-dsp",
   "myusuf3/imghash",
   "nfnt/resize",
   "r9y9/gossp",
   "umahmood/iter",
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
