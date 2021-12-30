package youtube

import (
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "strconv"
)

type Format struct {
   Itag int64
   URL string
   MimeType string
   Bitrate int64
   Width int64
   Height int64
   ContentLength int64 `json:"contentLength,string"`
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

func (f Format) Write(dst io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   LogLevel.Dump(req)
   pro := format.Content(f.ContentLength)
   for pro.Content < pro.ContentLength {
      req.Header.Set("Range", pro.Range())
      pro.Print()
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.Copy(dst, res.Body); err != nil {
         return err
      }
      pro.Content += pro.PartLength
   }
   return nil
}
