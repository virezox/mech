package tiktok

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
   "time"
)

func (a AwemeDetail) URL() string {
   for _, addr := range a.Video.Play_Addr.URL_List {
      return addr
   }
   return ""
}

var LogLevel format.LogLevel

type AwemeDetail struct {
   Author struct {
      Unique_ID string
   }
   Aweme_ID string
   Create_Time int64
   // height field here is invalid
   Video struct {
      Duration int64
      Play_Addr struct {
         Width int
         Height int
         URL_List []string
      }
   }
}

func NewAwemeDetail(id int64) (*AwemeDetail, error) {
   req, err := http.NewRequest(
      "GET", "http://api2.musical.ly/aweme/v1/aweme/detail/", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "aweme_id=" + strconv.FormatInt(id, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var detail struct {
      Aweme_Detail AwemeDetail
   }
   if err := json.NewDecoder(res.Body).Decode(&detail); err != nil {
      return nil, err
   }
   return &detail.Aweme_Detail, nil
}

func (a AwemeDetail) Duration() time.Duration {
   return time.Duration(a.Video.Duration) * time.Millisecond
}

func (a AwemeDetail) Time() time.Time {
   return time.Unix(a.Create_Time, 0)
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
