package youtube

import (
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "os"
   "strconv"
   "time"
)

func (f Format) Write(dst io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   format.Log.Dump(req)
   var (
      begin = time.Now()
      content int64
   )
   for content < f.ContentLength {
      buf := []byte("bytes=")
      buf = strconv.AppendInt(buf, content, 10)
      buf = append(buf, '-')
      buf = strconv.AppendInt(buf, content+partLength-1, 10)
      req.Header.Set("Range", string(buf))
      end := time.Since(begin).Milliseconds()
      if end >= 1 {
         format.PercentInt64(os.Stdout, content, f.ContentLength)
         os.Stdout.WriteString("\t")
         format.Size.LabelInt64(os.Stdout, content)
         os.Stdout.WriteString("\t")
         format.Rate.LabelInt64(os.Stdout, 1000*content/end)
         os.Stdout.WriteString("\n")
      }
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.Copy(dst, res.Body); err != nil {
         return err
      }
      content += partLength
   }
   return nil
}

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

const partLength = 10_000_000


