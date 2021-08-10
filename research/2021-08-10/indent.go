package main

import (
   "fmt"
   "golang.org/x/net/html"
   "io"
   "os"
   "strings"
)

type Encoder struct {
   io.Writer
   Indent string
}

func NewEncoder(w io.Writer) Encoder {
   return Encoder{Writer: w}
}

func (e Encoder) Encode(r io.Reader) error {
   var indent string
   z := html.NewTokenizer(r)
   for {
      tt := z.Next()
      if tt == html.ErrorToken {
         break
      }
      if tt == html.EndTagToken {
         indent = indent[len(e.Indent):]
      }
      t := z.Token().String()
      if tt == html.TextToken && strings.TrimSpace(t) == "" {
         continue
      }
      _, err := io.WriteString(e.Writer, indent + t + "\n")
      if err != nil {
         return err
      }
      // case html.StartTagToken, html.SelfClosingTagToken:
      if tt == html.StartTagToken {
         indent += e.Indent
      }
   }
   return nil
}

func (e *Encoder) SetIndent(indent string) {
   e.Indent = indent
}

func main() {
   f, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   z := html.NewTokenizer(f)
   for {
      tt := z.Next()
      if tt == html.ErrorToken {
         break
      }
      if tt == html.EndTagToken {
         fmt.Print("EndTagToken ")
      }
      if tt == html.StartTagToken {
         fmt.Print("StartTagToken ")
      }
      if tt == html.TextToken {
         fmt.Print("TextToken ")
      }
      if tt == html.SelfClosingTagToken {
         fmt.Print("SelfClosingTagToken ")
      }
      fmt.Println(z.Token())
   }
}
