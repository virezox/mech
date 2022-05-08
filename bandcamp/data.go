package bandcamp

import (
   "encoding/json"
   "github.com/89z/format/xml"
   "net/http"
)

/*
2809477874
https://schnaussandmunk.bandcamp.com/track/amaris-2
{"i_id":2809477874,"a_id":3454424886,"i_type":"t"}

1670971920
https://schnaussandmunk.bandcamp.com/album/passage-2
{"i_id":1670971920,"a_id":3454424886,"i_type":"a"}

3454424886
https://schnaussandmunk.bandcamp.com
{"i_id":3596433032,"a_id":3454424886,"i_type":"a"}

3454424886
https://schnaussandmunk.bandcamp.com/music
{"a_id":3454424886}
*/

type DataBand struct {
   Name string
}

type DataTralbum struct {
   Art_ID int64
   Album_Release_Date string
   Current struct {
      Title string
   }
   Artist string
   TrackInfo []struct {
      Track_Num int64
      Title string
      File *struct {
         MP3_128 string `json:"mp3-128"`
      }
   }
}

func NewDataTralbum(addr string) (*DataTralbum, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   scan, err := xml.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte(" data-tralbum=")
   scan.Scan()
   scan.Split = []byte("<script data-tralbum=")
   var script struct {
      DataTralbum []byte `xml:"data-tralbum,attr"`
   }
   if err := scan.Decode(&script); err != nil {
      return nil, err
   }
   data := new(DataTralbum)
   if err := json.Unmarshal(script.DataTralbum, data); err != nil {
      return nil, err
   }
   return data, nil
}
