package youtube

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "regexp"
   "strconv"
)

/*
Current logic is based on this input:

var uy={VP:function(a){a.reverse()},
eG:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
li:function(a,b){a.splice(0,b)}};
vy=function(a){a=a.split("");uy.eG(a,50);uy.eG(a,48);uy.eG(a,23);uy.eG(a,31);return a.join("")};

if this fails in the future, we should keep a record of all failed cases, to
keep from repeating a mistake.
*/
func decrypt(sig, body []byte) error {
   // get line
   line := regexp.MustCompile(`\.split\(""\);[^\n]+`).Find(body)
   // get swaps
   matches := regexp.MustCompile(`\d+`).FindAll(line, -1)
   for _, match := range matches {
      pos, err := strconv.Atoi(string(match))
      if err != nil { return err }
      pos %= len(sig)
      sig[0], sig[pos] = sig[pos], sig[0]
   }
   return nil
}

func httpGet(addr url.URL) ([]byte, error) {
   get := addr.String()
   println("Get", get)
   res, err := http.Get(get)
   if err != nil { return nil, err }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

type Video struct {
   StreamingData struct {
      AdaptiveFormats []struct {
         Bitrate int
         Height int
         Itag int
         MimeType string
         SignatureCipher string
      }
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
      }
   }
   VideoDetails struct {
      ShortDescription string
      Title string
      VideoId string
      ViewCount int `json:",string"`
   }
}

func (v Video) Description() string { return v.VideoDetails.ShortDescription }

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }

func (v Video) signatureCipher(itag int) (string, error) {
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { return format.SignatureCipher, nil }
   }
   return "", errors.New("itag not found")
}
