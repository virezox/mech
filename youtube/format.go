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

type Format struct {
   Bitrate bitrate
   ContentLength mech.ContentLength `json:"contentLength,string"`
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
   var pos mech.ContentLength
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%d-%d", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      speed := newBitrate(pos, begin)
      fmt.Printf("%d%% %v %v\n", 100*pos/f.ContentLength, bytes, speed)
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

type bitrate int64

func newBitrate(c mech.ContentLength, t time.Time) bitrate {
   // this is float64
   end := time.Since(t).Seconds()
   if end < 1 {
      return 0
   }
   speed := c / mech.ContentLength(end)
   return bitrate(speed)
}

func (b bitrate) String() string {
   met := []string{"B/s", "kB/s", "MB/s", "GB/s"}
   return mech.NumberFormat(float64(b), met)
}
