package youtube

import (
   "errors"
   "github.com/89z/rosso/os"
   "io"
   "mime"
   "net/http"
   "strconv"
)

func (self Format) Ext() (string, error) {
   media, _, err := mime.ParseMediaType(self.MimeType)
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
   return "", errors.New(self.MimeType)
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

func (self Format) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, "Quality:"...)
   if self.QualityLabel != "" {
      b = append(b, self.QualityLabel...)
   } else {
      b = append(b, self.AudioQuality...)
   }
   b = append(b, " Bitrate:"...)
   b = strconv.AppendInt(b, self.Bitrate, 10)
   if self.ContentLength >= 1 { // Tq92D6wQ1mg
      b = append(b, " ContentLength:"...)
      b = strconv.AppendInt(b, self.ContentLength, 10)
   }
   b = append(b, "\n\tMimeType:"...)
   b = append(b, self.MimeType...)
   b = append(b, '\n')
   return b, nil
}

type Formats []Format

func (self Format) Encode(w io.Writer) error {
   req, err := http.NewRequest("GET", self.URL, nil)
   if err != nil {
      return err
   }
   pro := os.Progress_Bytes(w, self.ContentLength)
   var pos int64
   for pos < self.ContentLength {
      b := []byte("bytes=")
      b = strconv.AppendInt(b, pos, 10)
      b = append(b, '-')
      b = strconv.AppendInt(b, pos+chunk-1, 10)
      req.Header.Set("Range", string(b))
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

func (self Formats) Audio(quality string) (*Format, bool) {
   for _, form := range self {
      if form.AudioQuality == quality {
         return &form, true
      }
   }
   return nil, false
}

func (self Formats) Video(height int) (*Format, bool) {
   distance := func(self *Format) int {
      if self.Height > height {
         return self.Height - height
      }
      return height - self.Height
   }
   var (
      ok bool
      output *Format
   )
   for i, input := range self {
      // since codecs are in this order avc1,vp9,av01,
      // do "<=" so we can get last one
      if output == nil || distance(&input) <= distance(output) {
         output = &self[i]
         ok = true
      }
   }
   return output, ok
}

const chunk = 10_000_000
