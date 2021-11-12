package youtube

import (
   "fmt"
   "github.com/89z/mech"
   "io"
   "net/http"
   "sort"
   "strings"
   "time"
)

const chunk = 10_000_000

func bitrate(pos int64, begin time.Time) string {
   end := time.Since(begin).Seconds()
   if end < 1 {
      return ""
   }
   rate := float64(pos) / end
   metric := []string{"B/s", "kB/s", "MB/s", "GB/s"}
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
   fmt.Println(req.Method, req.URL)
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
      if res.StatusCode != http.StatusPartialContent {
         return fmt.Errorf("status %v", res.Status)
      }
      if _, err := io.Copy(w, res.Body); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}

type FormatSlice []Format

func (f FormatSlice) Filter(keep func(Format)bool) FormatSlice {
   var forms FormatSlice
   for _, form := range f {
      if keep(form) {
         forms = append(forms, form)
      }
   }
   return forms
}

func (f FormatSlice) Sort(less ...func(a, b Format) bool) {
   if less == nil {
      less = []func(a, b Format) bool{
         func(a, b Format) bool {
            return b.Height < a.Height
         },
         func(a, b Format) bool {
            f, s := strings.Index, "/mp4;"
            return f(a.MimeType, s) < f(b.MimeType, s)
         },
         func(a, b Format) bool {
            return b.Bitrate < a.Bitrate
         },
      }
   }
   sort.Slice(f, func(a, b int) bool {
      fa, fb := f[a], f[b]
      for _, fn := range less {
         if fn(fa, fb) {
            return true
         }
         if fn(fb, fa) {
            break
         }
      }
      return false
   })
}
