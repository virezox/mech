package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
)

func (f Format) Encode(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   pro := format.Progress_Bytes(w, f.Content_Length)
   var pos int64
   for pos < f.Content_Length {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      res, err := HTTP_Client.Level(0).Redirect().Status(206).Do(req)
      if err != nil {
         return err
      }
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}

type Formats []Format

func (f Formats) Audio(quality string) (*Format, bool) {
   for _, form := range f {
      if form.AudioQuality == quality {
         return &form, true
      }
   }
   return nil, false
}

func (f Formats) Video(height int) (*Format, bool) {
   distance := func(f *Format) int {
      if f.Height > height {
         return f.Height - height
      }
      return height - f.Height
   }
   var (
      ok bool
      output *Format
   )
   for i, input := range f {
      // since codecs are in this order avc1,vp9,av01,
      // do "<=" so we can get last one
      if output == nil || distance(&input) <= distance(output) {
         output = &f[i]
         ok = true
      }
   }
   return output, ok
}

func (f Format) Format(s fmt.State, verb rune) {
   if f.QualityLabel != "" {
      fmt.Fprint(s, "Quality:", f.QualityLabel)
   } else {
      fmt.Fprint(s, "Quality:", f.AudioQuality)
   }
   fmt.Fprint(s, " Bitrate:", f.Bitrate)
   if f.Content_Length >= 1 { // Tq92D6wQ1mg
      fmt.Fprint(s, " Size:", f.Content_Length)
   }
   fmt.Fprint(s, " Codec:", f.MimeType)
   if verb == 'a' {
      fmt.Fprint(s, " URL:", f.URL)
   }
}

func (f Formats) Media_Type() error {
   for i, form := range f {
      _, param, err := mime.ParseMediaType(form.MimeType)
      if err != nil {
         return err
      }
      f[i].MimeType = param["codecs"]
   }
   return nil
}

const chunk = 10_000_000

type Format struct {
   AudioQuality string
   Bitrate int
   Content_Length int64 `json:"contentLength,string"`
   Height int
   MimeType string
   QualityLabel string
   URL string
   Width int
}
