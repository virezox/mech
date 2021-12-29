package mech

import (
   "bytes"
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
)

func Clean(char rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, char) {
      return -1
   }
   return char
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

type InvalidSlice struct {
   Index, Length int
}

func (i InvalidSlice) Error() string {
   index, length := int64(i.Index), int64(i.Length)
   var buf []byte
   buf = append(buf, "index out of range ["...)
   buf = strconv.AppendInt(buf, index, 10)
   buf = append(buf, "] with length "...)
   buf = strconv.AppendInt(buf, length, 10)
   return string(buf)
}

type LogLevel int

func (l LogLevel) Dump(req *http.Request) error {
   switch l {
   case 0:
      fmt.Println(req.Method, req.URL)
   case 1:
      buf, err := httputil.DumpRequest(req, true)
      if err != nil {
         return err
      }
      os.Stdout.Write(buf)
      if !bytes.HasSuffix(buf, []byte{'\n'}) {
         os.Stdout.WriteString("\n")
      }
   case 2:
      buf, err := httputil.DumpRequestOut(req, true)
      if err != nil {
         return err
      }
      os.Stdout.Write(buf)
   }
   return nil
}

type NotFound struct {
   Input string
}

func (n NotFound) Error() string {
   return strconv.Quote(n.Input) + " not found"
}

type Values map[string]string

func (v Values) Encode() string {
   vals := make(url.Values)
   for key, val := range v {
      vals.Set(key, val)
   }
   return vals.Encode()
}

func (v Values) Header() http.Header {
   vals := make(http.Header)
   for key, val := range v {
      vals.Set(key, val)
   }
   return vals
}

func (v Values) Reader() io.Reader {
   enc := v.Encode()
   return strings.NewReader(enc)
}

type Progress struct {
   *http.Response
   begin time.Time
   document int64
   part int
}

// Read method has pointer receiver
func NewProgress(res *http.Response) *Progress {
   pro := Progress{Response: res}
   pro.begin = time.Now()
   return &pro
}

func (p *Progress) Read(buf []byte) (int, error) {
   if p.part == 0 {
      p.meter()
   }
   read, err := p.Body.Read(buf)
   if err != nil {
      return 0, err
   }
   p.document += int64(read)
   p.part += read
   if p.part >= 10_000_000 {
      p.part = 0
   }
   return read, nil
}

func (p Progress) meter() {
   end := time.Since(p.begin).Milliseconds()
   if end > 0 {
      meter := format.PercentInt64(p.document, p.ContentLength)
      meter += "\t" + format.Size.LabelInt(p.document)
      meter += "\t" + format.Rate.LabelInt(1000 * p.document / end)
      fmt.Println(meter)
   }
}
