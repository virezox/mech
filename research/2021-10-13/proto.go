package youtube

import (
   "encoding/base64"
   "fmt"
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
   b, err := plenc.Marshal(nil, p)
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(b), nil
}

func (p *param) Relevance() {
   p.SortBy = pUint(0)
}

func (p *param) Rating() {
   p.SortBy = pUint(1)
}

func (p *param) UploadDate() {
   p.SortBy = pUint(2)
}

func (p *param) ViewCount() {
   p.SortBy = pUint(3)
}

func (p *param) LastHour() {
   p.Filter.UploadDate = pUint(1)
}

func (p *param) Today() {
   p.Filter.UploadDate = pUint(2)
}

func (p *param) ThisWeek() {
   p.Filter.UploadDate = pUint(3)
}

func (p *param) ThisMonth() {
   p.Filter.UploadDate = pUint(4)
}

func (p *param) ThisYear() {
   p.Filter.UploadDate = pUint(5)
}

func (p *param) Video() {
   p.Filter.Type = pUint(1)
}

func (p *param) Channel() {
   p.Filter.Type = pUint(2)
}

func (p *param) Playlist() {
   p.Filter.Type = pUint(3)
}

func (p *param) Movie() {
   p.Filter.Type = pUint(4)
}

func (p *param) UnderFourMinutes() {
   p.Filter.Duration = pUint(1)
}

func (p *param) OverTwentyMinutes() {
   p.Filter.Duration = pUint(2)
}

func (p *param) FourToTwentyMinutes() {
   p.Filter.Duration = pUint(3)
}

func (p *param) HD() {
   p.Filter.HD = pUint(1)
}

func (p *param) Subtitles() {
   p.Filter.Subtitles = pUint(1)
}

func (p *param) CreativeCommons() {
   p.Filter.CreativeCommons = pUint(1)
}

func (p *param) ThreeD() {
   p.Filter.ThreeD = pUint(1)
}

func (p *param) Live() {
   p.Filter.Live = pUint(1)
}

func (p *param) Purchased() {
   p.Filter.Purchased = pUint(1)
}

func (p *param) FourK() {
   p.Filter.FourK = pUint(1)
}

func (p *param) ThreeSixty() {
   p.Filter.ThreeSixty = pUint(1)
}

func (p *param) Location() {
   p.Filter.Location = pUint(1)
}

func (p *param) HDR() {
   p.Filter.HDR = pUint(1)
}

func (p *param) VR180() {
   p.Filter.VR180 = pUint(1)
}
