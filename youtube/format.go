package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "strconv"
   "strings"
   "time"
)

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int64
   Itag int64
   MimeType string
   URL string
}

func (f Format) String() string {
   buf := []byte("Itag:")
   buf = strconv.AppendInt(buf, f.Itag, 10)
   if f.Height >= 1 {
      buf = append(buf, " Height:"...)
      buf = strconv.AppendInt(buf, f.Height, 10)
   }
   buf = append(buf, " Bitrate:"...)
   buf = strconv.AppendInt(buf, f.Bitrate, 10)
   buf = append(buf, " ContentLength:"...)
   buf = strconv.AppendInt(buf, f.ContentLength, 10)
   justType, _, err := mime.ParseMediaType(f.MimeType)
   if err == nil {
      buf = append(buf, " MimeType:"...)
      buf = append(buf, justType...)
   }
   return string(buf)
}

func (f Format) Write(dst io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   LogLevel.Dump(req)
   par := newPartial(f.ContentLength)
   for par.value < par.total {
      fmt.Println(par.progress())
      req.Header.Set("Range", par.bytes)
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.Copy(dst, res.Body); err != nil {
         return err
      }
      par.value += par.chunk
   }
   return nil
}

type partial struct {
   begin time.Time
   bytes string
   value, chunk, total int64
}

func newPartial(total int64) partial {
   par := partial{chunk: 10_000_000, total: total}
   par.begin = time.Now()
   return par
}

func (p *partial) progress() string {
   var str strings.Builder
   percent := format.PercentInt64(p.value, p.total)
   str.WriteString(percent)
   str.WriteByte(' ')
   buf := []byte("bytes=")
   buf = strconv.AppendInt(buf, p.value, 10)
   buf = append(buf, '-')
   buf = strconv.AppendInt(buf, p.value+p.chunk-1, 10)
   p.bytes = string(buf)
   str.Write(buf)
   end := time.Since(p.begin).Milliseconds()
   if end > 0 {
      rate := format.Rate.LabelInt(1000 * p.value / end)
      str.WriteByte(' ')
      str.WriteString(rate)
   }
   return str.String()
}
