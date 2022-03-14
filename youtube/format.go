package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "os"
   "time"
)

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
      req.Header.Set(
         "Range", fmt.Sprintf("bytes=%v-%v", content, content+partLength-1),
      )
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
   return fmt.Sprintf("%q is not present", n.value)
}

func (f Format) Format(s fmt.State, verb rune) {
   fmt.Fprint(s, "Itag:", f.Itag)
   if f.QualityLabel != "" {
      fmt.Fprint(s, " Quality:", f.QualityLabel)
   } else {
      fmt.Fprint(s, " Quality:", f.AudioQuality)
   }
   fmt.Fprint(s, " Bitrate:", f.Bitrate)
   fmt.Fprint(s, " Size:", f.ContentLength)
   fmt.Fprint(s, " Type:", f.MimeType)
   if verb == 'a' {
      fmt.Fprint(s, " URL:", f.URL)
   }
}

const partLength = 10_000_000

var LogLevel format.LogLevel

type Format struct {
   AudioQuality string
   Bitrate int
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   QualityLabel string
   URL string
   Width int
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

