package main

import (
   "golang.org/x/net/html"
   "io"
)

type node struct { *html.Node }

func newNode(r io.Reader) (node, error) {
   d, err := html.Parse(r)
   if err != nil {
      return node{}, err
   }
   return node{d}, nil
}

func (n node) query(f func(n node) bool) node {
   in := []node{n}
   for len(in) > 0 {
      n := in[0]
      if f(n) {
         return n
      }
      for c := n.FirstChild; c != nil; c = c.NextSibling {
         in = append(in, node{c})
      }
      in = in[1:]
   }
   return node{}
}

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

func (n node) allTag(t string) []node {
   return n.queryAll(func(n node) bool {
      return n.Data == t
   })
}

func (n node) tag(t string) node {
   return n.query(func(n node) bool {
      return n.Data == t
   })
}

func (n node) text() string {
   n = n.query(func(n node) bool {
      return n.Type == html.TextNode
   })
   return n.Data
}

func class(c string) func(n node) bool {
   return func(n node) bool {
      for _, a := range n.Attr {
         if a.Key == "class" && a.Val == c {
            return true
         }
      }
      return false
   }
}

func (n node) class(c string) node {
   return n.query(class(c))
}

func (n node) allClass(c string) []node {
   return n.queryAll(func(n node) bool {
      for _, a := range n.Attr {
         if a.Key == "class" && a.Val == c {
            return true
         }
      }
      return false
   })
}
