package tiktok

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
)

const Origin = "https://api2.musical.ly"

var LogLevel mech.LogLevel

type Detail struct {
   Aweme_Detail struct {
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
}

func NewDetail(id uint64) (*Detail, error) {
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
   det := new(Detail)
   if err := json.NewDecoder(res.Body).Decode(det); err != nil {
      return nil, err
   }
   return det, nil
}
