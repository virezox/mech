package mech

import (
   "bufio"
   "golang.org/x/net/html"
   "io"
   "net/http"
   "net/textproto"
   "net/url"
   "strings"
)

func ReadRequest(r io.Reader) (*http.Request, error) {
   t := textproto.NewReader(bufio.NewReader(r))
   s, err := t.ReadLine()
   if err != nil {
      return nil, err
   }
   h, err := t.ReadMIMEHeader()
   if err != nil {
      return nil, err
   }
   f := strings.Fields(s)
   p, err := url.Parse(f[1])
   if err != nil {
      return nil, err
   }
   p.Host = h.Get("Host")
   return &http.Request{
      Body: io.NopCloser(t.R),
      Header: http.Header(h),
      Method: f[0],
      URL: p,
   }, nil
}

type Node struct {
   *html.Node
   todo []*html.Node
   callback func(*html.Node) bool
}

func Parse(r io.Reader) (*Node, error) {
   n, err := html.Parse(r)
   if err != nil {
      return nil, err
   }
   return &Node{
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

func (n Node) ByAttr(key, val string) *Node {
   n.todo = []*html.Node{n.Node}
   n.callback = func(c *html.Node) bool {
      for _, attr := range c.Attr {
         if attr.Key == key && attr.Val == val {
            return true
         }
      }
      return false
   }
   return &n
}

// keep source as is, return modified copy
func (n Node) ByTag(tag string) *Node {
   n.todo = []*html.Node{n.Node}
   n.callback = func(c *html.Node) bool {
      // x/net/html lowercases the tags
      return strings.EqualFold(c.Data, tag)
   }
   return &n
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
