package youtube

import (
   "mime"
   "strings"
)

func getCodec(v string) (string, error) {
   _, param, err := mime.ParseMediaType(v)
   if err != nil {
      return "", err
   }
   before, _, _ := strings.Cut(param["codecs"], ".")
   return before, nil
}
