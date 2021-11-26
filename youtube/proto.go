package youtube

import (
   "encoding/base64"
   "encoding/json"
   "github.com/89z/parse/protobuf"
)

const (
   // UPLOAD DATE
   UploadDateLastHour = 1
   UploadDateToday = 2
   UploadDateThisWeek = 3
   UploadDateThisMonth = 4
   UploadDateThisYear = 5
   // TYPE
   TypeVideo = 1
   TypeChannel = 2
   TypePlaylist = 3
   TypeMovie = 4
   // DURATION
   DurationUnderFourMinutes = 1
   DurationOverTwentyMinutes = 2
   DurationFourToTwentyMinutes = 3
   // SORT BY
   SortByRelevance = 0
   SortByRating = 1
   SortByUploadDate = 2
   SortByViewCount = 3
)

type Filter struct {
   UploadDate int `json:"1,omitempty"`
   Type int `json:"2,omitempty"`
   Duration int `json:"3,omitempty"`
   HD bool `json:"4,omitempty"`
   Subtitles bool `json:"5,omitempty"`
   CreativeCommons bool `json:"6,omitempty"`
   ThreeD bool `json:"7,omitempty"`
   Live bool `json:"8,omitempty"`
   Purchased bool `json:"9,omitempty"`
   FourK bool `json:"14,omitempty"`
   ThreeSixty bool `json:"15,omitempty"`
   Location bool `json:"23,omitempty"`
   HDR bool `json:"25,omitempty"`
   VR180 bool `json:"26,omitempty"`
}

type Params struct {
   SortBy int `json:"1,omitempty"`
   Filter *Filter `json:"2,omitempty"`
}

func (p Params) Encode() (string, error) {
   buf, err := json.Marshal(p)
   if err != nil {
      return "", err
   }
   mes := make(protobuf.Message)
   if err := mes.UnmarshalJSON(buf); err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(mes.Marshal()), nil
}
