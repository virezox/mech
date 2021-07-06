package main

import (
   "encoding/json"
   "fmt"
   "net/http"
)

type cover struct {
   Images []struct {
      Image string
   }
}

func newCover(id string) (string, error) {
   addr := fmt.Sprintf("http://archive.org/download/mbid-%v/index.json", id)
   res, err := http.Get(addr)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   var cov cover
   json.NewDecoder(res.Body).Decode(&cov)
   for _, img := range cov.Images {
      return img.Image, nil
   }
   return "", fmt.Errorf("%q fail", id)
}

func main() {
   s, err := newCover("a40cb6e9-c766-37c4-8677-7eb51393d5a1")
   if err != nil {
      panic(err)
   }
   println(s)
}
