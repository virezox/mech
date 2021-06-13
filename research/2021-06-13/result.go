package main

import (
   "golang.org/x/net/html"
   "io"
)

func tag(t string) func(node) bool {
   return func(n node) bool {
      return n.Data == t
   }
}

type node struct {
   *html.Node
}

func newNode(r io.Reader) (node, error) {
   d, err := html.Parse(r)
   if err != nil {
      return node{}, err
   }
   return node{d}, nil
}

func (n node) attr(key string) string {
   for _, a := range n.Attr {
      if a.Key == key {
         return a.Val
      }
   }
   return ""
}

func (n node) byAttrAll(key, val string) []node {
   return n.selector(true, func(n node) bool {
      for _, a := range n.Attr {
         if a.Key == key && a.Val == val {
            return true
         }
      }
      return false
   })
}

func (n node) byTag(name string) node {
   for _, n := range n.selector(false, tag(name)) {
      return n
   }
   return node{}
}

func (n node) byTagAll(name string) []node {
   return n.selector(true, tag(name))
}

func (n node) selector(all bool, f func(n node) bool) []node {
   var (
      in = []node{n}
      out []node
   )
   for len(in) > 0 {
      n := in[0]
      if f(n) {
         out = append(out, n)
         if ! all {
            break
         }
      }
      for c := n.FirstChild; c != nil; c = c.NextSibling {
         in = append(in, node{c})
      }
      in = in[1:]
   }
   return out
}

func (n node) text() string {
   for _, n := range n.selector(false, func(n node) bool {
      return n.Type == html.TextNode
   }) {
      return n.Data
   }
   return ""
}
