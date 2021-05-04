package youtube

import (
   "errors"
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
   for _, match := range regexp.MustCompile(`\d+`).FindAll(line, -1) {
      pos, err := strconv.Atoi(string(match))
      if err != nil { return err }
      pos %= len(sig)
      sig[0], sig[pos] = sig[pos], sig[0]
   }
   return nil
}

type Format struct {
   Bitrate int
   Height int
   Itag int
   MimeType string
   SignatureCipher string
}

func (v Video) NewFormat(itag int) (Format, error) {
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { return format, nil }
   }
   return Format{}, errors.New("itag not found")
}

// NewRequest returns the url for a specific format
func (f Format) NewRequest() (*http.Request, error) {
   val, err := url.ParseQuery(f.SignatureCipher)
   if err != nil { return nil, err }
   sig := []byte(val.Get("s"))
   // get player
   body, err := getPlayer()
   if err != nil { return nil, err }
   // decrypt
   err = decrypt(sig, body)
   if err != nil { return nil, err }
   req, err := http.NewRequest("GET", val.Get("url"), nil)
   if err != nil { return nil, err }
   val = req.URL.Query()
   val.Set("sig", string(sig))
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Range", "bytes=0-")
   return req, nil
}

type Video struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
      }
   }
   VideoDetails struct {
      Author string
      ShortDescription string
      Title string
      ViewCount int `json:"viewCount,string"`
   }
}


func (v Video) Description() string { return v.VideoDetails.ShortDescription }

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }
