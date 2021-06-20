package youtube

import (
   "encoding/json"
   "fmt"
   "net/http"
   "strings"
)

type Android struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
}

func NewAndroid(id string) (Android, error) {
   body := fmt.Sprintf(`
   {
      "videoId": %q, "context": {
         "client": {"clientName": "ANDROID", "clientVersion": "15.01"}
      }
   }
   `, id)
   req, err := http.NewRequest(
      "POST", PlayerAPI, strings.NewReader(body),
   )
   if err != nil {
      return Android{}, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Android{}, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return Android{}, fmt.Errorf("status %v", res.Status)
   }
   var and Android
   json.NewDecoder(res.Body).Decode(&and)
   return and, nil
}

func (a Android) Formats() []Format {
   var formats []Format
   for _, format := range a.StreamingData.AdaptiveFormats {
      if format.ContentLength > 0 {
         formats = append(formats, format)
      }
   }
   return formats
}

func (a Android) NewFormat(itag int) (Format, error) {
   for _, format := range a.Formats() {
      if format.Itag == itag {
         return format, nil
      }
   }
   return Format{}, fmt.Errorf("itag %v", itag)
}
