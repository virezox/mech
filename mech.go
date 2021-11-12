package mech

import (
   "fmt"
   "mime"
   "strings"
   "net/http"
   "net/http/httputil"
   "os"
   "strconv"
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

func NumberFormat(val float64, met []string) string {
   var key int
   for val >= 1000 {
      val /= 1000
      key++
   }
   if key >= len(met) {
      return ""
   }
   return fmt.Sprintf("%.1f ", val) + met[key]
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

type ContentLength int64

func (c ContentLength) String() string {
   met := []string{"B", "kB", "MB", "GB"}
   return NumberFormat(float64(c), met)
}

type NotFound struct {
   String string
}

func (n NotFound) Error() string {
   return strconv.Quote(n.String) + " not found"
}

type Progress struct {
   *http.Response
   x, xMax int
   y ContentLength
}

func NewProgress(res *http.Response) *Progress {
   return &Progress{Response: res, xMax: 10_000_000}
}

func (p *Progress) Read(buf []byte) (int, error) {
   if p.x == 0 {
      percent := 100 * int64(p.y) / p.ContentLength
      fmt.Printf("%v%% %v\n", percent, p.y)
   }
   num, err := p.Body.Read(buf)
   if err != nil {
      return 0, err
   }
   p.y += ContentLength(num)
   p.x += num
   if p.x >= p.xMax {
      p.x = 0
   }
   return num, nil
}
