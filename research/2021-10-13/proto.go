package youtube

import (
   "encoding/base64"
   "github.com/segmentio/encoding/proto"
)

type param struct {
   SortBy uint `protobuf:"varint,1"`
   Filter struct {
      UploadDate uint `protobuf:"varint,1"`
      Type uint `protobuf:"varint,2"`
      Duration uint `protobuf:"varint,3"`
      HD uint `protobuf:"varint,4"`
      Subtitles uint `protobuf:"varint,5"`
      CreativeCommons uint `protobuf:"varint,6"`
      ThreeD uint `protobuf:"varint,7"`
      Live uint `protobuf:"varint,8"`
      Purchased uint `protobuf:"varint,9"`
      FourK uint `protobuf:"varint,14"`
      ThreeSixty uint `protobuf:"varint,15"`
      Location uint `protobuf:"varint,23"`
      HDR uint `protobuf:"varint,25"`
      VR180 uint `protobuf:"varint,26"`
   } `protobuf:"bytes,2"`
}

func (p param) Encode() (string, error) {
   data, err := proto.Marshal(p)
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(data), nil
}

func (p *param) Relevance() {
   p.SortBy = 0
}

func (p *param) Rating() {
   p.SortBy = 1
}

func (p *param) UploadDate() {
   p.SortBy = 2
}

func (p *param) ViewCount() {
   p.SortBy = 3
}

func (p *param) LastHour() {
   p.Filter.UploadDate = 1
}

func (p *param) Today() {
   p.Filter.UploadDate = 2
}

func (p *param) ThisWeek() {
   p.Filter.UploadDate = 3
}

func (p *param) ThisMonth() {
   p.Filter.UploadDate = 4
}

func (p *param) ThisYear() {
   p.Filter.UploadDate = 5
}

func (p *param) Video() {
   p.Filter.Type = 1
}

func (p *param) Channel() {
   p.Filter.Type = 2
}

func (p *param) Playlist() {
   p.Filter.Type = 3
}

func (p *param) Movie() {
   p.Filter.Type = 4
}

func (p *param) UnderFourMinutes() {
   p.Filter.Duration = 1
}

func (p *param) OverTwentyMinutes() {
   p.Filter.Duration = 2
}

func (p *param) FourToTwentyMinutes() {
   p.Filter.Duration = 3
}

func (p *param) HD() {
   p.Filter.HD = 1
}

func (p *param) Subtitles() {
   p.Filter.Subtitles = 1
}

func (p *param) CreativeCommons() {
   p.Filter.CreativeCommons = 1
}

func (p *param) ThreeD() {
   p.Filter.ThreeD = 1
}

func (p *param) Live() {
   p.Filter.Live = 1
}

func (p *param) Purchased() {
   p.Filter.Purchased = 1
}

func (p *param) FourK() {
   p.Filter.FourK = 1
}

func (p *param) ThreeSixty() {
   p.Filter.ThreeSixty = 1
}

func (p *param) Location() {
   p.Filter.Location = 1
}

func (p *param) HDR() {
   p.Filter.HDR = 1
}

func (p *param) VR180() {
   p.Filter.VR180 = 1
}
