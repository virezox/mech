package mech

import (
   "golang.org/x/net/html"
   "io"
)

type Decoder struct {
   *html.Tokenizer
   html.Token
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      Tokenizer: html.NewTokenizer(r),
   }
}

func (d *Decoder) AttrSelector(key, val string) bool {
   for {
      if d.Next() == html.ErrorToken {
         break
      }
      t := d.Tokenizer.Token()
      for _, a := range t.Attr {
         if a.Key == key && a.Val == val {
            d.Token = t
            return true
         }
      }
   }
   return false
}

func (d *Decoder) TextSelector() bool {
   for {
      n := d.Next()
      if n == html.ErrorToken {
         break
      }
      if n == html.TextToken {
         d.Token = d.Tokenizer.Token()
         return true
      }
   }
   return false
}

func (d Decoder) Attr(key string) string {
   for _, a := range d.Token.Attr {
      if a.Key == key {
         return a.Val
      }
   }
   return ""
}

func (d Decoder) Text() string {
   return d.Data
}
