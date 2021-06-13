package main

import (
   "fmt"
   "os"
)

func main() {
   file, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   doc, err := newNode(file)
   if err != nil {
      panic(err)
   }
   for _, tr := range doc.queryAll(class("lista2")) {
      fmt.Printf("%+v\n", tr.Node)
   }
   /*
   for _, tr := range doc.allClass("lista2") {
      tds := tr.allTag("td")
      // <td> <a> #text
      title := tds[1].tag("a").text()
      //        <td>    <a>        onmouseover
      r.Image = tds[1].FirstChild.Attr[0].Val
      //                   <td>    <a>        href
      r.Torrent = Origin + tds[1].FirstChild.Attr[2].Val
      //              <td>    #text     <span>      #text
      r.Genre = tds[1].LastChild.PrevSibling.FirstChild.Data
      //       <td>    #text
      r.Size = tds[3].FirstChild.Data
      //          <td>    <font>     #text
      r.Seeders = tds[4].FirstChild.FirstChild.Data
   }
   */
}
