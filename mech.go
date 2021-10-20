package mech

import (
   "fmt"
   "mime"
   "strings"
   "net/http"
   "net/http/httputil"
   "os"
)

var extensions = map[string][]string{
   "audio/mp4": {".m4a"},
   "audio/webm": {".weba"},
   "video/mp4": {".m4v", ".mp4", ".mp4v"},
   "video/webm": {".webm"},
}

var verbose bool

func Clean(r rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}

// github.com/golang/go/issues/22318
func ExtensionsByType(typ string) ([]string, error) {
   justType, _, err := mime.ParseMediaType(typ)
   if err != nil {
      return nil, err
   }
   return extensions[justType], nil
}

func RoundTrip(req *http.Request) (*http.Response, error) {
   if verbose {
      dum, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      if dum[len(dum)-1] != '\n' {
         dum = append(dum, '\n')
      }
      os.Stdout.Write(dum)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, true)
      if err != nil {
         return nil, err
      }
      return nil, fmt.Errorf("%s", dum)
   }
   return res, nil
}

func Verbose(v bool) {
   verbose = v
}
