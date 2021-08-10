package mech

import (
   "fmt"
   "strings"
   "testing"
)

const source = `
<!doctype html>
<html>
   <body>
      <table>
         <tr class="h">
            <th class="r">release</th>
            <th class="p">post</th>
         </tr>
         <tr class="d">
            <td class="r">2000</td>
            <td class="p">2021</td>
         </tr>
      </table>
   </body>
</html>
`

func TestNode(t *testing.T) {
   doc, err := Parse(strings.NewReader(source))
   if err != nil {
      t.Fatal(err)
   }
   // ByAttr
   if r := doc.ByAttr("class", "r"); r.Scan() {
      if r.Data != "th" {
         t.Fatal(r.Data)
      }
   }
   // ByAttrAll
   r := doc.ByAttr("class", "r")
   for r.Scan() {
      fmt.Printf("%+v\n", r.Node)
   }
   // can we scan the doc directly?
   for doc.Scan() {
      fmt.Println(doc.Data)
   }
}
