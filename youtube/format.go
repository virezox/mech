package youtube

import (
   "errors"
   "github.com/89z/rosso/os"
   "io"
   "mime"
   "net/http"
   "strconv"
)

func (f Format) Ext() (string, error) {
   media, _, err := mime.ParseMediaType(f.MimeType)
   if err != nil {
      return "", err
   }
   switch media {
   case "audio/mp4":
      return ".m4a", nil
   case "audio/webm":
      return ".weba", nil
   case "video/mp4":
      return ".m4v", nil
   case "video/webm":
      return ".webm", nil
   }
   return "", errors.New(f.MimeType)
}

type Format struct {
   AudioQuality string
   QualityLabel string
   Width int
   Height int
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   MimeType string
   URL string
}

func (f Format) MarshalText() ([]byte, error) {
   var buf []byte
   buf = append(buf, "Quality:"...)
   if f.QualityLabel != "" {
      buf = append(buf, f.QualityLabel...)
   } else {
      buf = append(buf, f.AudioQuality...)
   }
   buf = append(buf, " Bitrate:"...)
   buf = strconv.AppendInt(buf, f.Bitrate, 10)
   if f.ContentLength >= 1 { // Tq92D6wQ1mg
      buf = append(buf, " ContentLength:"...)
      buf = strconv.AppendInt(buf, f.ContentLength, 10)
   }
   buf = append(buf, "\n\tMimeType:"...)
   buf = append(buf, f.MimeType...)
   buf = append(buf, '\n')
   return buf, nil
}

type Formats []Format

func (f Format) Encode(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   pro := os.Progress_Bytes(w, f.ContentLength)
   var pos int64
   for pos < f.ContentLength {
      buf := []byte("bytes=")
      buf = strconv.AppendInt(buf, pos, 10)
      buf = append(buf, '-')
      buf = strconv.AppendInt(buf, pos+chunk-1, 10)
      req.Header.Set("Range", string(buf))
      res, err := HTTP_Client.Level(0).Redirect(nil).Status(206).Do(req)
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

const chunk = 10_000_000
