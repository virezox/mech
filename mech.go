package mech

import (
   "fmt"
   "mime"
   "strings"
)

func Clean(r rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}

func Ext(typ string) (string, error) {
   exts, err := mime.ExtensionsByType(typ)
   if err != nil {
      return "", err
   }
   for _, ext := range exts {
      return ext, nil
   }
   return "", fmt.Errorf("%q has no associated extension", typ)
}
