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

func (p Player) Base() string {
   return format.Clean(p.VideoDetails.Author + "-" + p.VideoDetails.Title)
}

func (f Format) Ext() (string, error) {
   exts, err := mime.ExtensionsByType(f.MimeType)
   if err != nil {
      return "", err
   }
   for _, ext := range exts {
      return ext, nil
   }
   return "", notPresent{f.MimeType}
}

const partLength = 10_000_000

var LogLevel format.LogLevel

type Format struct {
   AudioQuality string
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int64
   MimeType string
   QualityLabel string
   URL string
   Width int
}

func (f Format) Format() (string, error) {
   buf := []byte("Itag:")
   buf = strconv.AppendInt(buf, f.Itag, 10)
   buf = append(buf, " Quality:"...)
   if f.QualityLabel != "" {
      buf = append(buf, f.QualityLabel...)
   } else {
      buf = append(buf, f.AudioQuality...)
   }
   buf = append(buf, " Bitrate:"...)
   buf = strconv.AppendInt(buf, f.Bitrate, 10)
   buf = append(buf, " Size:"...)
   buf = strconv.AppendInt(buf, f.ContentLength, 10)
   justType, _, err := mime.ParseMediaType(f.MimeType)
   if err != nil {
      return "", err
   }
   buf = append(buf, " Type:"...)
   buf = append(buf, justType...)
   if f.URL != "" {
      buf = append(buf, " URL:"...)
      buf = append(buf, f.URL...)
   }
   return string(buf), nil
}

func (f Format) Write(dst io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   LogLevel.Dump(req)
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
      end := time.Since(begin).Seconds()
      if end >= 1 {
         os.Stdout.WriteString(format.PercentInt64(content, f.ContentLength))
         os.Stdout.WriteString("\t")
         os.Stdout.WriteString(format.Size.GetInt64(content))
         os.Stdout.WriteString("\t")
         os.Stdout.WriteString(format.Rate.Get(float64(content)/end))
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

type notPresent struct {
   value string
}

func (n notPresent) Error() string {
   return strconv.Quote(n.value) + " is not present"
}
