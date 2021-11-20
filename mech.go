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

var Verbose bool

var extensions = map[string][]string{
   "audio/mp4": {".m4a"},
   "audio/webm": {".weba"},
   "video/mp4": {".m4v", ".mp4", ".mp4v"},
   "video/webm": {".webm"},
}

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

func NumberFormat(val float64, metric []string) string {
   var key int
   for val >= 1000 {
      val /= 1000
      key++
   }
   if key >= len(metric) {
      return ""
   }
   return strconv.FormatFloat(val, 'f', 3, 64) + " " + metric[key]
}

func Percent(pos, length int64) string {
   return strconv.FormatInt(100*pos/length, 10) + "%"
}

func RoundTrip(req *http.Request) (*http.Response, error) {
   if Verbose {
      buf, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      if buf[len(buf)-1] != '\n' {
         buf = append(buf, '\n')
      }
      os.Stdout.Write(buf)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, status{res}
   }
   return res, nil
}

type Invalid struct {
   Input string
}

func (i Invalid) Error() string {
   return strconv.Quote(i.Input) + " invalid"
}

type NotFound struct {
   Input string
}

func (n NotFound) Error() string {
   return strconv.Quote(n.Input) + " not found"
}

type status struct {
   *http.Response
}

func (s status) Error() string {
   buf, err := httputil.DumpResponse(s.Response, true)
   if err != nil {
      return err.Error()
   }
   return string(buf)
}

type Progress struct {
   *http.Response
   metric []string
   x, xMax int
   y int64
}

func NewProgress(res *http.Response) *Progress {
   return &Progress{
      Response: res,
      metric: []string{"B", "kB", "MB", "GB"},
      xMax: 10_000_000,
   }
}

func (p *Progress) Read(buf []byte) (int, error) {
   if p.x == 0 {
      bytes := NumberFormat(float64(p.y), p.metric)
      fmt.Println(Percent(p.y, p.ContentLength), bytes)
   }
   num, err := p.Body.Read(buf)
   if err != nil {
      return 0, err
   }
   p.y += int64(num)
   p.x += num
   if p.x >= p.xMax {
      p.x = 0
   }
   return num, nil
}
