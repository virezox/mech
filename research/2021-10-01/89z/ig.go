package main

import (
   "github.com/Davincible/goinsta"
   "os"
)

func main() {  
   insta := goinsta.New("srpen6", os.Args[1])
   err := insta.Login()
   if err != nil {
      panic(err)
   }
   insta.Export("ig.json")
}
