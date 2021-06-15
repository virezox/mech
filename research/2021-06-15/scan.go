package main

import (
   "golang.org/x/net/html"
   "io"
)

type Scanner struct {
   *html.Node
   todo []*html.Node
   callback func(*html.Node) bool
}

func NewScanner(r io.Reader) (*Scanner, error) {
   n, err := html.Parse(r)
   if err != nil {
      return nil, err
   }
   return &Scanner{Node: n}, nil
}

func (s *Scanner) Scan() bool {
   for len(s.todo) > 0 {
      n := s.todo[0]
      s.todo = s.todo[1:]
      if s.callback(n) {
         s.Node = n
         return true
      }
      for n = n.FirstChild; n != nil; n = n.NextSibling {
         s.todo = append(s.todo, n)
      }
   }
   return false
}

func (s *Scanner) Split(tag string) {
   s.todo = []*html.Node{s.Node}
   s.callback = func(n *html.Node) bool {
      return n.Data == tag
   }
}
