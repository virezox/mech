package youtube

import (
   "encoding/base64"
   "github.com/segmentio/encoding/proto"
)

type Param struct {
   SortBy int `protobuf:"varint,1"`
   Filter struct {
      UploadDate int `protobuf:"varint,1"`
      Type int `protobuf:"varint,2"`
      Duration int `protobuf:"varint,3"`
      HD int `protobuf:"varint,4"`
      Subtitles int `protobuf:"varint,5"`
      CreativeCommons int `protobuf:"varint,6"`
      ThreeD int `protobuf:"varint,7"`
      Live int `protobuf:"varint,8"`
      Purchased int `protobuf:"varint,9"`
      FourK int `protobuf:"varint,14"`
      ThreeSixty int `protobuf:"varint,15"`
      Location int `protobuf:"varint,23"`
      HDR int `protobuf:"varint,25"`
      VR180 int `protobuf:"varint,26"`
   } `protobuf:"bytes,2"`
}

// "EgIQAg=="
func (p *Param) Channel() {
   p.Filter.Type = 2
}

// "EgIwAQ==
func (p *Param) CreativeCommons() {
   p.Filter.CreativeCommons = 1
}

func (p Param) Encode() (string, error) {
   data, err := proto.Marshal(p)
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(data), nil
}

// "EgJwAQ=="
func (p *Param) FourK() {
   p.Filter.FourK = 1
}

// "EgIYAw=="
func (p *Param) FourToTwentyMinutes() {
   p.Filter.Duration = 3
}

// "EgIgAQ=="
func (p *Param) HD() {
   p.Filter.HD = 1
}

// "EgPIAQE="
func (p *Param) HDR() {
   p.Filter.HDR = 1
}

// "EgIIAQ=="
func (p *Param) LastHour() {
   p.Filter.UploadDate = 1
}

// "EgJAAQ=="
func (p *Param) Live() {
   p.Filter.Live = 1
}

// "EgO4AQE="
func (p *Param) Location() {
   p.Filter.Location = 1
}

// "EgIQBA=="
func (p *Param) Movie() {
   p.Filter.Type = 4
}

// "EgIYAg=="
func (p *Param) OverTwentyMinutes() {
   p.Filter.Duration = 2
}

// "EgIQAw=="
func (p *Param) Playlist() {
   p.Filter.Type = 3
}

// "EgJIAQ=="
func (p *Param) Purchased() {
   p.Filter.Purchased = 1
}

// "CAE="
func (p *Param) Rating() {
   p.SortBy = 1
}

// ""
func (p *Param) Relevance() {
   p.SortBy = 0
}

// "EgIoAQ=="
func (p *Param) Subtitles() {
   p.Filter.Subtitles = 1
}

// "EgIIBA=="
func (p *Param) ThisMonth() {
   p.Filter.UploadDate = 4
}

// "EgIIAw=="
func (p *Param) ThisWeek() {
   p.Filter.UploadDate = 3
}

// "EgIIBQ=="
func (p *Param) ThisYear() {
   p.Filter.UploadDate = 5
}

// "EgI4AQ=="
func (p *Param) ThreeD() {
   p.Filter.ThreeD = 1
}

// "EgJ4AQ=="
func (p *Param) ThreeSixty() {
   p.Filter.ThreeSixty = 1
}

// "EgIIAg=="
func (p *Param) Today() {
   p.Filter.UploadDate = 2
}

// "EgIYAQ=="
func (p *Param) UnderFourMinutes() {
   p.Filter.Duration = 1
}

// "CAI="
func (p *Param) UploadDate() {
   p.SortBy = 2
}

// "EgPQAQE="
func (p *Param) VR180() {
   p.Filter.VR180 = 1
}

// "EgIQAQ=="
func (p *Param) Video() {
   p.Filter.Type = 1
}

// "CAM="
func (p *Param) ViewCount() {
   p.SortBy = 3
}
