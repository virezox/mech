package main

import (
   "fmt"
   "os"
   "regexp"
)

func findSubmatch(file string) (string, error) {
   b, err := os.ReadFile(file)
   if err != nil {
      return "", err
   }
   re := regexp.MustCompile("/vi/([^/]+)/")
   find := re.FindSubmatch(b)
   if find == nil {
      return "", fmt.Errorf("FindSubmatch %v", re)
   }
   return string(find[1]), nil
}

func main() {
   s, err := findSubmatch("index.html")
   if err != nil {
      panic(err)
   }
   println(s)
}
