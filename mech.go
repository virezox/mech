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

var VoidElement = map[string]bool{
   "br": true,
   "img": true,
   "input": true,
   "meta": true,
}

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

type Encoder struct {
   io.Writer
   indent string
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
         indent = indent[len(e.indent):]
      }
      s := string(z.Raw())
      if tt == html.TextToken && strings.TrimSpace(s) == "" {
         continue
      }
      _, err := io.WriteString(e.Writer, indent + s + "\n")
      if err != nil {
         return err
      }
      if tt == html.StartTagToken && !VoidElement[z.Token().Data] {
         indent += e.indent
      }
   }
   return nil
}

func (e *Encoder) SetIndent(indent string) {
   e.indent = indent
}

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
