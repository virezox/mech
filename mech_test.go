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
   // make sure we only mutate if we want to:
   if tr := doc.ByTag("tr"); tr.Next() {
      if doc.Data != "" {
         t.Fatal(doc.Data)
      }
   }
   // make sure Next actually works:
   if tr := doc.ByTag("tr"); tr.Next() {
      if tr.Data != "tr" {
         t.Fatal(tr.Data)
      }
   }
   // make sure we can only get valid children:
   if tr := doc.ByTag("tr"); tr.Next() {
      if td := tr.ByTag("td"); td.Next() {
         t.Fatal("tr[0].td")
      }
   }
   // test Attr
   if tr := doc.ByTag("tr"); tr.Next() {
      if class := tr.Attr("class"); class != "h" {
         t.Fatal(class)
      }
   }
   // Text
   if td := doc.ByTag("td"); td.Next() {
      if td.Text() != "2000" {
         t.Fatal(td.Text())
      }
   }
   // ByAttr
   if r := doc.ByAttr("class", "r"); r.Next() {
      if r.Data != "th" {
         t.Fatal(r.Data)
      }
   }
   // ByTagAll
   tr := doc.ByTag("tr")
   for tr.Next() {
      fmt.Printf("%+v\n", tr.Node)
   }
   // ByAttrAll
   r := doc.ByAttr("class", "r")
   for r.Next() {
      fmt.Printf("%+v\n", r.Node)
   }
   // mutate
   if doc = doc.ByTag("tr"); doc.Next() {
      if doc.Data != "tr" {
         t.Fatal(doc.Data)
      }
   }
}
