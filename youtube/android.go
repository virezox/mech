package youtube

import (
   "encoding/json"
   "fmt"
)

type Android struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   VideoDetails struct {
      Author string
      Title string
   }
}

func NewAndroid(id string) (Android, error) {
   res, err := post(id, "ANDROID", "15.01")
   if err != nil {
      return Android{}, err
   }
   defer res.Body.Close()
   var and Android
   json.NewDecoder(res.Body).Decode(&and)
   return and, nil
}

func (a Android) Author() string {
   return a.VideoDetails.Author
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

func (a Android) Title() string {
   return a.VideoDetails.Title
}
