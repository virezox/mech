package html

import (
   "github.com/tdewolff/parse/v2/html"
   "strings"
   stdhtml "html"
)

// textContent
func (l *Lexer) Bytes() []byte {
   for {
      switch l.TokenType {
      case html.ErrorToken:
         return nil
      case html.TextToken:
         return l.data
      }
      l.TokenType, l.data = l.Next()
   }
}

// getAttribute
func (l Lexer) GetAttr(key string) string {
   val := l.attr[key]
   trim := strings.Trim(val, `'"`)
   return stdhtml.UnescapeString(trim)
}

// hasAttribute
func (l Lexer) HasAttr(key string) bool {
   _, ok := l.attr[key]
   return ok
}

// getElementsByClassName
func (l *Lexer) NextAttr(key, val string) bool {
   for {
      switch l.TokenType, _ = l.Next(); l.TokenType {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         l.attr = make(map[string]string)
      case html.AttributeToken:
         l.attr[string(l.Text())] = string(l.AttrVal())
      case html.StartTagCloseToken, html.StartTagVoidToken:
         if l.HasAttr(key) && l.GetAttr(key) == val {
            return true
         }
      }
   }
}

// getElementsByTagName
func (l *Lexer) NextTag(name string) bool {
   for {
      switch l.TokenType, _ = l.Next(); l.TokenType {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         if string(l.Text()) == name {
            return true
         }
      }
   }
}
