package tiktok

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
)

const Origin = "http://api2.musical.ly"

var LogLevel mech.LogLevel

type AwemeDetail struct {
   Author struct {
      Unique_ID string
   }
   Aweme_ID string
   Create_Time int
   // height field here is invalid
   Video struct {
      Duration int
      Play_Addr struct {
         Width int
         Height int
         URL_List []string
      }
   }
}

func NewAwemeDetail(id uint64) (*AwemeDetail, error) {
   req, err := http.NewRequest("GET", Origin + "/aweme/v1/aweme/detail/", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "aweme_id=" + strconv.FormatUint(id, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, mech.Response{res}
   }
   var detail struct {
      Aweme_Detail AwemeDetail
   }
   if err := json.NewDecoder(res.Body).Decode(&detail); err != nil {
      return nil, err
   }
   return &detail.Aweme_Detail, nil
}

func (a AwemeDetail) URL() (string, error) {
   sLen := len(a.Video.Play_Addr.URL_List)
   if sLen == 0 {
      return "", mech.InvalidSlice{}
   }
   return a.Video.Play_Addr.URL_List[0], nil
}
