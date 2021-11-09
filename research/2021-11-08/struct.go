package youtube

import (
   "encoding/base64"
   "github.com/89z/parse/protobuf"
)

type sParam struct {
   SortBy int `json:"1"`
   Filter struct {
      UploadDate int `json:"1"`
      Type int `json:"2"`
      Duration int `json:"3"`
      HD int `json:"4"`
      Subtitles int `json:"5"`
      CreativeCommons int `json:"6"`
      ThreeD int `json:"7"`
      Live int `json:"8"`
      Purchased int `json:"9"`
      FourK int `json:"14"`
      ThreeSixty int `json:"15"`
      Location int `json:"23"`
      HDR int `json:"25"`
      VR180 int `json:"26"`
   } `json:"2"`
}

func (s sParam) encode() (string, error) {
   enc, err := protobuf.NewEncoder(s)
   if err != nil {
      return "", err
   }
   buf, err := enc.Encode()
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(buf), nil
}

// "EgIIAQ=="
func (s *sParam) lastHour() {
   s.Filter.UploadDate = 1
}
