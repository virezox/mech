package scan

import (
   "fmt"
   "strings"
   "testing"
)

func TestScan(t *testing.T) {
   doc, err := NewScanner(strings.NewReader(source))
   if err != nil {
      t.Fatal(err)
   }
   // make sure we only mutate if we want to:
   tr := doc.Split("tr")
   tr.Scan()
   if doc.Data != "" {
      t.Fatal(doc.Data)
   }
   // make sure we can only get valid children:
   td := tr.Split("td")
   if td.Scan() {
      t.Fatal("tr[0].td")
   }
   // make sure Scan actually works:
   if tr.Data != "tr" {
      t.Fatal(tr.Data)
   }
   // test Attr
   class := tr.Attr("class")
   if class != "h" {
      t.Fatal(class)
   }
   // ByTagAll
   tr = doc.Split("tr")
   for tr.Scan() {
      fmt.Printf("%+v\n", tr.Node)
   }
   // Text
   td = doc.Split("td")
   td.Scan()
   if td.Text() != "2000" {
      t.Fatal(td.Text())
   }
   // mutate
   doc = doc.Split("tr")
   doc.Scan()
   if doc.Data != "tr" {
      t.Fatal(doc.Data)
   }
}

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
