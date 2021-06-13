package main

import (
   "errors"
   "golang.org/x/net/html"
   "os"
   "regexp"
)

func selector(n *html.Node, f func(n *html.Node) bool) []*html.Node {
   var (
      in = []*html.Node{n}
      out []*html.Node
   )
   for len(in) > 0 {
      n := in[0]
      if f(n) {
         out = append(out, n)
      }
      for c := n.FirstChild; c != nil; c = c.NextSibling {
         in = append(in, c)
      }
      in = in[1:]
   }
   return out
}

var notFound = errors.New("not found")

func parse(file string) (string, error) {
   f, err := os.Open(file)
   if err != nil {
      return "", err
   }
   defer f.Close()
   n1, err := html.Parse(f)
   if err != nil {
      return "", err
   }
   for _, n2 := range selector(n1, func(n *html.Node) bool {
      return n.Data == "ytd-video-renderer"
   }) {
      for _, n3 := range selector(n2, func(n *html.Node) bool {
         return n.Data == "a"
      }) {
         for _, a := range n3.Attr {
            if a.Key == "href" {
               return a.Val, nil
            }
         }
         break
      }
      break
   }
   return "", notFound
}

func findSubmatch(file string) (string, error) {
   b, err := os.ReadFile(file)
   if err != nil {
      return "", err
   }
   re := regexp.MustCompile("/vi/([^/]+)/")
   find := re.FindSubmatch(b)
   if find == nil {
      return "", notFound
   }
   return string(find[1]), nil
}

func main() {
   {
      s, err := findSubmatch("begin.html")
      if err != nil {
         panic(err)
      }
      println(s)
   }
   {
      s, err := parse("begin.html")
      if err != nil {
         panic(err)
      }
      println(s)
   }
}
