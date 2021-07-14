package main

import (
   "fmt"
   "text/template"
   "strings"
)

type Format struct {
  Bitrate       int64
  ContentLength int64 `json:"contentLength,string"`
  Height        int
  Itag          int
  MimeType      string
}

func format(s string, v interface{}) string {
   t, err := new(template.Template).Parse(s)
   if err != nil {
      return ""
   }
   b := new(strings.Builder)
   t.Execute(b, v)
   return b.String()
}

func main() {
   f := Format{Bitrate:1, ContentLength:2, Height:3, Itag:4, MimeType:"MimeType"}
   g := struct{
      Format
      BitrateS, LengthS string
   }{f, "BitrateS", "LengthS"}
   fmt.Println(format(
      "itag {{.Itag}}, height {{.Height}}, " +
      "{{.BitrateS}}, {{.LengthS}}, {{.MimeType}}", g,
   ))
}
