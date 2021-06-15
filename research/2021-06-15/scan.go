package scan

import (
   "golang.org/x/net/html"
   "io"
)

type Scanner struct {
   *html.Node
   todo []*html.Node
   callback func(*html.Node) bool
}

func NewScanner(r io.Reader) (Scanner, error) {
   n, err := html.Parse(r)
   if err != nil {
      return Scanner{}, err
   }
   return Scanner{Node: n}, nil
}

// keep source as is, return modified copy
func (s Scanner) Split(tag string) Scanner {
   s.todo = []*html.Node{s.Node}
   s.callback = func(n *html.Node) bool {
      return n.Data == tag
   }
   return s
}

// this can modify the struct now, as we are working with a copy
func (s *Scanner) Scan() bool {
   for len(s.todo) > 0 {
      n := s.todo[0]
      s.todo = s.todo[1:]
      if s.callback(n) {
         s.Node = n
         return true
      }
      for c := n.FirstChild; c != nil; c = c.NextSibling {
         s.todo = append(s.todo, c)
      }
   }
   return false
}

func (s Scanner) Attr(key string) string {
   for _, attr := range s.Node.Attr {
      if attr.Key == key {
         return attr.Val
      }
   }
   return ""
}

func (s Scanner) Text() string {
   for c := s.FirstChild; c != nil; c = c.NextSibling {
      if c.Type == html.TextNode {
         return c.Data
      }
   }
   return ""
}
