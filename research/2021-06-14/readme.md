# June 14 2021

If you dont mind something a little simpler, I wrote my own module for doing
this. This one does what youre asking:

~~~
package main

import (
   "github.com/89z/mech"
   "strings"
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
   r := strings.NewReader(source)
   doc, err := mech.NewNode(r)
   if err != nil {
      panic(err)
   }
   div := doc.ByAttr("class", "funcname").Text()
   println(div) // One
}
~~~

Other examples:

~~~
for _, div := range doc.ByAttrAll("class", "funcname") {
   println(div.Text()) // One, Two
}
~~~

~~~
doc.ByTag("div").Text() // One
~~~

~~~
for _, div := range doc.ByTagAll("div") {
   println(div.Text()) // One, Two
}
~~~

~~~
doc.ByTag("div").Attr("class") // funcname
~~~

- https://github.com/89z/mech/blob/2b68f17d/old/youtube.go
- https://golang.org/pkg/encoding/xml
- https://stackoverflow.com/questions/12883079/go-parse-html-table
- https://stackoverflow.com/questions/27998747/parse-html-page-and-get-value
