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

var LogLevel = 1

func Dump(req *http.Request) error {
   switch LogLevel {
   case 1:
      fmt.Println(req.Method, req.URL)
   case 2:
      buf, err := httputil.DumpRequest(req, true)
      if err != nil {
         return err
      }
      os.Stdout.Write(buf)
      if buf[len(buf)-1] != '\n' {
         os.Stdout.WriteString("\n")
      }
   case 3:
      buf, err := httputil.DumpRequestOut(req, true)
      if err != nil {
         return err
      }
      os.Stdout.Write(buf)
   }
   return nil
}

func Clean(r rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}

// github.com/golang/go/issues/22318
func ExtensionByType(typ string) (string, error) {
   justType, _, err := mime.ParseMediaType(typ)
   if err != nil {
      return "", err
   }
   switch justType {
   case "audio/mp4":
      return ".m4a", nil
   case "audio/webm":
      return ".weba", nil
   case "video/mp4":
      return ".m4v", nil
   case "video/webm":
      return ".webm", nil
   }
   return "", NotFound{justType}
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

type String string

func (s String) At(i int) (byte, error) {
   high := len(s) - 1
   if i < 0 || i > high {
      return 0, outOfRange{i, 0, high}
   }
   return s[i], nil
}

func (s String) Slice(i, j int) (String, error) {
   if i < 0 || i > j {
      return "", outOfRange{i, 0, j}
   }
   high := len(s) - 1
   if j < i || j > high {
      return "", outOfRange{j, i, high}
   }
   return s[i:j], nil
}

type Strings []string

func (s Strings) At(i int) (string, error) {
   high := len(s) - 1
   if i < 0 || i > high {
      return "", outOfRange{i, 0, high}
   }
   return s[i], nil
}

func (s Strings) AtInt(i int) (int64, error) {
   str, err := s.At(i)
   if err != nil {
      return 0, err
   }
   return strconv.ParseInt(str, 10, 64)
}

func (s Strings) Slice(i, j int) ([]string, error) {
   if i < 0 || i > j {
      return nil, outOfRange{i, 0, j}
   }
   high := len(s) - 1
   if j < i || j > high {
      return nil, outOfRange{j, i, high}
   }
   return s[i:j], nil
}

type outOfRange struct {
   index, low, high int
}

func (o outOfRange) Error() string {
   index := int64(o.index)
   low := int64(o.low)
   high := int64(o.high)
   buf := []byte("index ")
   buf = strconv.AppendInt(buf, index, 10)
   buf = append(buf, " out of range "...)
   buf = strconv.AppendInt(buf, low, 10)
   buf = append(buf, ':')
   buf = strconv.AppendInt(buf, high, 10)
   return string(buf)
}
