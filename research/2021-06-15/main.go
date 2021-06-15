package main

import (
   "fmt"
   "strings"
)

const source = `
<!DOCTYPE html>
<html>
   <body>
      <header>
         <div class="ten">10</div>
         <div class="eleven">11</div>
      </header>
   </body>
</html>
`

func main() {
   s, err := NewScanner(strings.NewReader(source))
   if err != nil {
      panic(err)
   }
   s.Split("header")
   s.Scan()
   fmt.Printf("%+v\n", s.Node)
   s.Split("div")
   for s.Scan() {
      fmt.Printf("%+v\n", s.Node)
   }
   // this all works, but can I get back to <html>?
}
