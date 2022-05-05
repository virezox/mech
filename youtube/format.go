package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "strings"
)

const chunk = 10_000_000

type Format struct {
   AudioQuality string
   Bitrate int
   ContentLength int64 `json:"contentLength,string"`
   Height int
   MimeType string
   QualityLabel string
   URL string
   Width int
}

func (f Format) Format(s fmt.State, verb rune) {
   if f.QualityLabel != "" {
      fmt.Fprint(s, "Quality:", f.QualityLabel)
   } else {
      fmt.Fprint(s, "Quality:", f.AudioQuality)
   }
   fmt.Fprint(s, " Bitrate:", f.Bitrate)
   if f.ContentLength >= 1 { // Tq92D6wQ1mg
      fmt.Fprint(s, " Size:", f.ContentLength)
   }
   fmt.Fprint(s, " Type:", f.MimeType)
   if verb == 'a' {
      fmt.Fprint(s, " URL:", f.URL)
   }
}

func (f Format) WriteTo(w io.Writer) (int64, error) {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return 0, err
   }
   LogLevel.Dump(req)
   var (
      pos int64
      pro = format.ProgressBytes(w, f.ContentLength)
   )
   for pos < f.ContentLength {
      req.Header.Set(
         "Range", fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1),
      )
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return 0, err
      }
      if _, err := io.Copy(pro, res.Body); err != nil {
         return 0, err
      }
      if err := res.Body.Close(); err != nil {
         return 0, err
      }
      pos += chunk
   }
   return f.ContentLength, nil
}

type Formats []Format

func (f Formats) Audio(quality string) *Format {
   for _, form := range f {
      if form.AudioQuality == quality {
         return &form
      }
   }
   return nil
}

func (f Formats) MediaType() error {
   for i, form := range f {
      typ, param, err := mime.ParseMediaType(form.MimeType)
      if err != nil {
         return err
      }
      param["codecs"], _, _ = strings.Cut(param["codecs"], ".")
      f[i].MimeType = mime.FormatMediaType(typ, param)
   }
   return nil
}

func (f Formats) Video(height int) *Format {
   distance := func(f *Format) int {
      if f.Height > height {
         return f.Height - height
      }
      return height - f.Height
   }
   var dst *Format
   for i, src := range f {
      if i == 0 || distance(&src) < distance(dst) {
         dst = &f[i]
      }
   }
   return dst
}
