package instagram

import (
   "text/scanner"
   "strings"
)

func textScanner(address string) string {
   var buf scanner.Scanner
   buf.Init(strings.NewReader(address))
   buf.Mode = scanner.ScanIdents
   buf.IsIdentRune = func(r rune, i int) bool {
      return r != '/' && r != scanner.EOF
   }
   for buf.Scan() != scanner.EOF {
      if buf.TokenText() == "p" {
         buf.Scan()
         buf.Scan()
         return buf.TokenText()
      }
   }
   return ""
}
