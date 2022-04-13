package mech

import (
   "bytes"
   "encoding/json"
   "mime"
   "strconv"
   "strings"
)

func Clean(path string) string {
   fn := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return -1
      }
      return r
   }
   return strings.Map(fn, path)
}

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

func ExtensionByType(typ string) (string, error) {
   media, _, err := mime.ParseMediaType(typ)
   if err != nil {
      return "", err
   }
   switch media {
   case "audio/webm":
      return ".weba", nil
   case "video/webm":
      return ".webm", nil
   case "audio/mp4":
      return ".m4a", nil
   case "video/mp4":
      return ".m4v", nil
   }
   return "", notFound{typ}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}
