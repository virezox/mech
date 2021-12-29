package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "strconv"
   "time"
)

type Format struct {
   Itag int64
   URL string
   MimeType string
   Bitrate int64
   Width int64
   Height int64
   ContentLength int64 `json:"contentLength,string"`
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
   pro := newProgress(f.ContentLength)
   for pro.value < pro.total {
      pro.setBytes()
      pro.meter()
      req.Header.Set("Range", pro.bytes)
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.Copy(dst, res.Body); err != nil {
         return err
      }
      pro.value += pro.chunk
   }
   return nil
}

type progress struct {
   begin time.Time
   bytes string
   value, chunk, total int64
}

func newProgress(total int64) progress {
   pro := progress{chunk: 10_000_000, total: total}
   pro.begin = time.Now()
   return pro
}

func (p progress) meter() {
   end := time.Since(p.begin).Milliseconds()
   if end > 0 {
      meter := format.PercentInt64(p.value, p.total)
      meter += "\t" + p.bytes
      meter += "\t" + format.Rate.LabelInt(1000 * p.value / end)
      fmt.Println(meter)
   }
}

func (p *progress) setBytes() {
   buf := []byte("bytes=")
   buf = strconv.AppendInt(buf, p.value, 10)
   buf = append(buf, '-')
   buf = strconv.AppendInt(buf, p.value+p.chunk-1, 10)
   p.bytes = string(buf)
}
