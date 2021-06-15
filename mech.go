package mech

import (
   "golang.org/x/net/html"
   "io"
)

type Node struct {
   *html.Node
   todo []*html.Node
   callback func(*html.Node) bool
}

func Parse(r io.Reader) (Node, error) {
   n, err := html.Parse(r)
   if err != nil {
      return Node{}, err
   }
   return Node{
      n, []*html.Node{n}, func(*html.Node) bool {
         return true
      },
   }, nil
}

func (n Node) Attr(key string) string {
   for _, attr := range n.Node.Attr {
      if attr.Key == key {
         return attr.Val
      }
   }
   return ""
}

// keep source as is, return modified copy
func (n Node) ByAttr(key, val string) Node {
   n.todo = []*html.Node{n.Node}
   n.callback = func(c *html.Node) bool {
      for _, attr := range c.Attr {
         if attr.Key == key && attr.Val == val {
            return true
         }
      }
      return false
   }
   return n
}

func (n Node) ByTag(tag string) Node {
   n.todo = []*html.Node{n.Node}
   n.callback = func(c *html.Node) bool {
      return c.Data == tag
   }
   return n
}

// this can modify the struct now, as we are working with a copy
func (n *Node) Scan() bool {
   for len(n.todo) > 0 {
      t := n.todo[0]
      n.todo = n.todo[1:]
      for c := t.FirstChild; c != nil; c = c.NextSibling {
         n.todo = append(n.todo, c)
      }
      if n.callback(t) {
         n.Node = t
         return true
      }
   }
   return false
}

func (n Node) Text() string {
   for c := n.FirstChild; c != nil; c = c.NextSibling {
      if c.Type == html.TextNode {
         return c.Data
      }
   }
   return ""
}
