package mech

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "strings"
)

type HtmlReader struct {
   *html.Lexer
   html.TokenType
   data []byte
   attr map[string]string
}

func NewHtmlReader(r io.Reader) HtmlReader {
   return HtmlReader{
      Lexer: html.NewLexer(parse.NewInput(r)),
   }
}

func (h HtmlReader) Attr(key string) (string, bool) {
   val, ok := h.attr[key]
   if !ok {
      return "", false
   }
   return strings.Trim(val, `'"`), true
}

func (h *HtmlReader) Bytes() []byte {
   for {
      switch h.TokenType {
      case html.ErrorToken:
         return nil
      case html.TextToken:
         return h.data
      }
      h.TokenType, h.data = h.Next()
   }
}

func (h *HtmlReader) NextAttr(key, val string) bool {
   for {
      switch h.TokenType, _ = h.Next(); h.TokenType {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         h.attr = make(map[string]string)
      case html.AttributeToken:
         h.attr[string(h.Text())] = string(h.AttrVal())
      case html.StartTagCloseToken, html.StartTagVoidToken:
         if v, ok := h.Attr(key); ok && v == val {
            return true
         }
      }
   }
}

func (h *HtmlReader) NextTag(name string) bool {
   for {
      switch h.TokenType, _ = h.Next(); h.TokenType {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         if string(h.Text()) == name {
            return true
         }
      }
   }
}
