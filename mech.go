package mech

import (
   "bytes"
   "encoding/json"
   "mime"
   "net/http"
   "strconv"
   "strings"
)

func Encode[T any](value T) (*bytes.Buffer, error) {
   buf := new(bytes.Buffer)
   enc := json.NewEncoder(buf)
   enc.SetIndent("", " ")
   err := enc.Encode(value)
   if err != nil {
      return nil, err
   }
   return buf, nil
}

func Clean(path string) string {
   fn := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return -1
      }
      return r
   }
   return strings.Map(fn, path)
}

func Ext(head http.Header) (string, error) {
   typ := head.Get("Content-Type")
   exts, err := mime.ExtensionsByType(typ)
   if err != nil {
      return "", err
   }
   for _, ext := range exts {
      return ext, nil
   }
   return "", notFound{typ}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}
