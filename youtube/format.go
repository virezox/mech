package youtube

import (
   "fmt"
   "github.com/89z/mech"
   "io"
   "net/http"
   "time"
)

const chunk = 10_000_000

func bitrate(pos int64, begin time.Time) string {
   end := time.Since(begin).Seconds()
   if end < 1 {
      return ""
   }
   rate := float64(pos) / end
   metric := []string{" B/s", " kB/s", " MB/s", " GB/s"}
   return mech.NumberFormat(rate, metric)
}

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   URL string
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
      bytes := fmt.Sprintf("bytes=%d-%d", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      percent := mech.Percent(pos, f.ContentLength)
      fmt.Println(percent, bytes, bitrate(pos, begin))
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
