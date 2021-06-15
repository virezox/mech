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

func TestMutate(t *testing.T) {
   doc, err := Parse(strings.NewReader(source))
   if err != nil {
      t.Fatal(err)
   }
   if doc = doc.ByTag("tr"); doc.Scan() {
      if doc.Data != "tr" {
         t.Fatal(doc.Data)
      }
   }
}

func TestNode(t *testing.T) {
   doc, err := Parse(strings.NewReader(source))
   if err != nil {
      t.Fatal(err)
   }
   // make sure we only mutate if we want to:
   if tr := doc.ByTag("tr"); tr.Scan() {
      if doc.Data != "" {
         t.Fatal(doc.Data)
      }
   }
   // make sure Scan actually works:
   if tr := doc.ByTag("tr"); tr.Scan() {
      if tr.Data != "tr" {
         t.Fatal(tr.Data)
      }
   }
   // make sure we can only get valid children:
   if tr := doc.ByTag("tr"); tr.Scan() {
      if td := tr.ByTag("td"); td.Scan() {
         t.Fatal("tr[0].td")
      }
   }
   // test Attr
   if tr := doc.ByTag("tr"); tr.Scan() {
      if class := tr.Attr("class"); class != "h" {
         t.Fatal(class)
      }
   }
   // Text
   if td := doc.ByTag("td"); td.Scan() {
      if td.Text() != "2000" {
         t.Fatal(td.Text())
      }
   }
   // ByAttr
   if r := doc.ByAttr("class", "r"); r.Scan() {
      if r.Data != "th" {
         t.Fatal(r.Data)
      }
   }
   // ByTagAll
   tr := doc.ByTag("tr")
   for tr.Scan() {
      fmt.Printf("%+v\n", tr.Node)
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
