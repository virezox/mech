package scan

import (
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
}
