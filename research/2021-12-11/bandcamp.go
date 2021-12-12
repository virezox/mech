package bandcamp

import (
   "encoding/json"
   "github.com/89z/mech"
   "github.com/89z/parse/net"
   "html"
   "io"
)

type dataTralbum struct {
   TrackInfo []struct {
      File struct {
         MP3_128 string `json:"mp3-128"`
      }
   }
}

func newDataTralbum(src io.Reader) (*dataTralbum, error) {
   for _, node := range net.ReadHTML(src, "script") {
      data, ok := node.Attr["data-tralbum"]
      if ok {
         data = html.UnescapeString(data)
         tra := new(dataTralbum)
         err := json.Unmarshal([]byte(data), tra)
         if err != nil {
            return nil, err
         }
         return tra, nil
      }
   }
   return nil, mech.NotFound{"data-tralbum"}
}
