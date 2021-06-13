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
   for _, tr := range doc.byAttrAll("class", "lista2") {
      tds := tr.byTagAll("td")
      // <td> <a> #text
      title := tds[1].byTag("a").text()
      fmt.Println(title)
      // <td> <a> onmouseover
      img := tds[1].byTag("a").attr("onmouseover")
      fmt.Println(img)
      // <td> <a> href
      tor := tds[1].byTag("a").attr("href")
      fmt.Println(tor)
      // <td> #text <span> #text
      gen := tds[1].byTag("span").text()
      fmt.Println(gen)
      // <td> #text
      size := tds[3].text()
      fmt.Println(size)
      // <td> <font> #text
      seed := tds[4].byTag("font").text()
      fmt.Println(seed)
      // newline
      fmt.Println()
   }
}
