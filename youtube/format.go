package youtube

import (
   "fmt"
   "github.com/89z/mech"
   "io"
   "mime"
   "net/http"
   "strconv"
   "time"
)

const chunk = 10_000_000

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

func (f Format) Write(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   LogLevel.Dump(req)
   begin := time.Now()
   var pos int64
   for pos < f.ContentLength {
      percent := strconv.FormatInt(100*pos/f.ContentLength, 10) + "% "
      bytes := fmt.Sprintf("bytes=%d-%d", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Print(percent, bytes)
      if end := time.Since(begin).Milliseconds(); end > 0 {
         f, symbol := mech.FormatRate(1000 * pos / end)
         fmt.Printf(" %.3f%v", f, symbol)
      }
      fmt.Println()
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.Copy(w, res.Body); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}
