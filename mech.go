package mech
// github.com/89z

import (
   "bytes"
   "encoding/json"
   "mime"
   "strconv"
)

func Extension_By_Type(typ string) (string, error) {
   media, _, err := mime.ParseMediaType(typ)
   if err != nil {
      return "", err
   }
   switch media {
   case "audio/mpeg":
      return ".mp3", nil
   case "audio/mp4":
      return ".m4a", nil
   case "audio/webm":
      return ".weba", nil
   case "video/mp4":
      return ".m4v", nil
   case "video/webm":
      return ".webm", nil
   }
   return "", not_found{typ}
}

type not_found struct {
   value string
}

func (n not_found) Error() string {
   var buf []byte
   buf = strconv.AppendQuote(buf, n.value)
   buf = append(buf, " is not found"...)
   return string(buf)
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
