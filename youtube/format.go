// YouTube
package youtube

import (
   "fmt"
   "io"
   "mime"
   "net/http"
   "sort"
)

const chunk = 10_000_000

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   URL string
}

func (f Format) Ext() string {
   exts, err := mime.ExtensionsByType(f.MimeType)
   if err != nil {
      return ""
   }
   return exts[0]
}

func (f Format) Write(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   var pos int64
   fmt.Println(invert, req.Method, reset, req.URL)
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Printf("%v%% %v\n", 100*pos/f.ContentLength, bytes)
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
            exts := map[string]int{".m4v": 1, ".m4a": 1}
            return exts[a.Ext()] < exts[b.Ext()]
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