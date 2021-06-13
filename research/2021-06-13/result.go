package main

import (
   "fmt"
   "golang.org/x/net/html"
   "io"
   "os"
)

func (n node) queryAll(f func(n node) bool) []node {
   var (
      in = []node{n}
      out []node
   )
   for len(in) > 0 {
      n := in[0]
      if f(n) {
         out = append(out, n)
      }
      for c := n.FirstChild; c != nil; c = c.NextSibling {
         in = append(in, node{c})
      }
      in = in[1:]
   }
   return out
}

type node struct { *html.Node }

func newNode(r io.Reader) (node, error) {
   d, err := html.Parse(r)
   if err != nil {
      return node{}, err
   }
   return node{d}, nil
}

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
   rows := doc.queryAll(func(n node) bool {
      if n.Data != "tr" {
         return false
      }
      if len(n.Attr) == 0 {
         return false
      }
      return n.Attr[0].Val == "lista2"
   })
   fmt.Println(rows)
   /*
   var results []Result
   for _, row := range rows {
      cols := queryAll(row, func(n *html.Node) bool {
         return n.Data == "td"
      })
      var r Result
      //        <td>    <a>        #text
      r.Title = cols[1].FirstChild.FirstChild.Data
      //        <td>    <a>        onmouseover
      r.Image = cols[1].FirstChild.Attr[0].Val
      //                   <td>    <a>        href
      r.Torrent = Origin + cols[1].FirstChild.Attr[2].Val
      //              <td>    #text     <span>      #text
      r.Genre = cols[1].LastChild.PrevSibling.FirstChild.Data
      //       <td>    #text
      r.Size = cols[3].FirstChild.Data
      //          <td>    <font>     #text
      r.Seeders = cols[4].FirstChild.FirstChild.Data
      // append
      results = append(results, r)
   }
   */
}
