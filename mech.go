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

func Dump(req *http.Request) error {
   if Verbose {
      buf, err := httputil.DumpRequest(req, true)
      if err != nil {
         return err
      }
      os.Stdout.Write(buf)
      if buf[len(buf)-1] != '\n' {
         os.Stdout.WriteString("\n")
      }
   } else {
      fmt.Println(req.Method, req.URL)
   }
   return nil
}

func Clean(r rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}

func Compare(a, b int) int {
   if a < b {
      return -1
   }
   if a == b {
      return 0
   }
   return 1
}

// github.com/golang/go/issues/22318
func ExtensionByType(typ string) (string, error) {
   justType, _, err := mime.ParseMediaType(typ)
   if err != nil {
      return "", err
   }
   ext, ok := map[string]string{
      "audio/mp4": ".m4a",
      "audio/webm": ".weba",
      "video/mp4": ".m4v",
      "video/webm": ".webm",
   }[justType]
   if !ok {
      return "", NotFound{justType}
   }
   return ext, nil
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
   // no space here, as some number are unitless
   return strconv.FormatFloat(val, 'f', 3, 64) + metric[key]
}

func Percent(pos, length int64) string {
   return strconv.FormatInt(100*pos/length, 10) + "%"
}

type NotFound struct {
   Input string
}

func (n NotFound) Error() string {
   return strconv.Quote(n.Input) + " not found"
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
      metric: []string{" B", " kB", " MB", " GB"},
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
