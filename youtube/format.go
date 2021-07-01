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
   fmt.Println(invert, "GET", reset, req.URL)
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Println(bytes)
      res, err := new(http.Transport).RoundTrip(req)
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

type Formats []Format

func (s Formats) Filter(keep func(Format)bool) Formats {
   var forms Formats
   for _, form := range s {
      if keep(form) {
         forms = append(forms, form)
      }
   }
   return forms
}

func (s Formats) Sort() {
   formatFns := []formatFn{
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
   sort.Slice(s, func(a, b int) bool {
      sa, sb := s[a], s[b]
      for _, fn := range formatFns {
         if fn(sa, sb) {
            return true
         }
         if fn(sb, sa) {
            break
         }
      }
      return false
   })
}

type formatFn func(a, b Format) bool
