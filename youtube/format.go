package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "time"
)

var LogLevel format.LogLevel

func (f Format) Write(dst io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   LogLevel.Dump(req)
   var (
      begin = time.Now()
      content int64
   )
   for content < f.ContentLength {
      req.Header.Set(
         "Range", fmt.Sprintf("bytes=%v-%v", content, content+partLength-1),
      )
      end := time.Since(begin).Seconds()
      if end >= 1 {
         fmt.Print(format.Percent(content, f.ContentLength), "\t")
         fmt.Print(format.LabelSize(content), "\t")
         fmt.Println(format.LabelRate(float64(content)/end))
      }
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.Copy(dst, res.Body); err != nil {
         return err
      }
      content += partLength
   }
   return nil
}

func (f *Format) ParseMediaType() error {
   typ, _, err := mime.ParseMediaType(f.MimeType)
   if err != nil {
      return err
   }
   f.MimeType = typ
   return nil
}

func (f Format) Format(s fmt.State, verb rune) {
   fmt.Fprint(s, "Itag:", f.Itag)
   if f.QualityLabel != "" {
      fmt.Fprint(s, " Quality:", f.QualityLabel)
   } else {
      fmt.Fprint(s, " Quality:", f.AudioQuality)
   }
   fmt.Fprint(s, " Bitrate:", f.Bitrate)
   fmt.Fprint(s, " Size:", f.ContentLength)
   fmt.Fprint(s, " Type:", f.MimeType)
   if verb == 'a' {
      fmt.Fprint(s, " URL:", f.URL)
   }
}

type notPresent struct {
   value string
}

func (n notPresent) Error() string {
   return fmt.Sprintf("%q is not present", n.value)
}

const partLength = 10_000_000

func (f Format) Ext() (string, error) {
   exts, err := mime.ExtensionsByType(f.MimeType)
   if err != nil {
      return "", err
   }
   for _, ext := range exts {
      return ext, nil
   }
   return "", notPresent{f.MimeType}
}

type Format struct {
   AudioQuality string
   Bitrate int
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   QualityLabel string
   URL string
   Width int
}

// We cannot do bitrate. If we look at this video:
//
// ID          | 720 low | 720 mid | 720 high | 1080 low
// ------------|---------|---------|----------|---------
// p-P5-7eV9GE | 1052477 | 1325265 | 1350187  | 2078318
//
// then the target would be 1350187. Then if we look at this video, 480 would
// be chosen:
//
// ID          | 480     | 720 low
// ------------|---------|--------
// qqiC88f9ogU | 1158788 | 2097952
type Height struct {
   StreamingData
   Target int
}

func (h Height) Less(i, j int) bool {
   return h.distance(i) < h.distance(j)
}

func (h Height) distance(i int) int {
   diff := h.AdaptiveFormats[i].Height - h.Target
   if diff >= 0 {
      return diff
   }
   return -diff
}
