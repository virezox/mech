package mech

import (
   "fmt"
   "mime"
   "strings"
   "net/http"
   "net/http/httputil"
   "os"
)

var Verbose bool

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
   if exts == nil {
      return "", fmt.Errorf("%q has no associated extension", typ)
   }
   return exts[0], nil
}

func RoundTrip(req *http.Request) (*http.Response, error) {
   if Verbose {
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
      return nil, fmt.Errorf("status %q", res.Status)
   }
   return res, nil
}
