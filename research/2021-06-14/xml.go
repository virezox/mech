package main

import (
   "encoding/xml"
   "fmt"
)

const source = `
<!DOCTYPE html>
<html>
   <body>
      <header>
         <div class="funcname">One</div>
         <div class="funcname">Two</div>
      </header>
   </body>
</html>
`

func main() {
   m := make(map[string]interface{})
   if err := xml.Unmarshal([]byte(source), &m); err != nil {
      panic(err)
   }
   fmt.Println(m)
}
