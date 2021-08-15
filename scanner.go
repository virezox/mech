package mech

import (
   "golang.org/x/net/html"
   "io"
)

type Scanner struct {
   *html.Tokenizer
   html.Token
}

func NewScanner(r io.Reader) Scanner {
   return Scanner{
      Tokenizer: html.NewTokenizer(r),
   }
}

func (s Scanner) Attr(key string) string {
   for _, a := range s.Token.Attr {
      if a.Key == key {
         return a.Val
      }
   }
   return ""
}

func (s Scanner) Bytes() []byte {
   return []byte(s.Data)
}

func (s *Scanner) ScanAttr(key, val string) bool {
   for {
      if s.Next() == html.ErrorToken {
         break
      }
      t := s.Tokenizer.Token()
      for _, a := range t.Attr {
         if a.Key == key && a.Val == val {
            s.Token = t
            return true
         }
      }
   }
   return false
}

func (s *Scanner) ScanText() bool {
   for {
      n := s.Next()
      if n == html.ErrorToken {
         break
      }
      if n == html.TextToken {
         s.Token = s.Tokenizer.Token()
         return true
      }
   }
   return false
}

func (s Scanner) Text() string {
   return s.Data
}
