package youtube

import (
   "encoding/json"
   "fmt"
)

const VersionAndroid = "15.01"

type Android struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   VideoDetails `json:"videoDetails"`
}

func NewAndroid(id string) (Android, error) {
   res, err := post(id, "ANDROID", VersionAndroid)
   if err != nil {
      return Android{}, err
   }
   defer res.Body.Close()
   var and Android
   if err := json.NewDecoder(res.Body).Decode(&and); err != nil {
      return Android{}, err
   }
   return and, nil
}

func (a Android) NewFormat(itag int) (Format, error) {
   for _, format := range a.StreamingData.AdaptiveFormats {
      if format.Itag == itag {
         return format, nil
      }
   }
   return Format{}, fmt.Errorf("itag %v", itag)
}
