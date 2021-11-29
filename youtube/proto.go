package youtube

import (
   "encoding/base64"
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
   protobuf.Message
}

func NewFilter() Filter {
   return Filter{
      make(protobuf.Message),
   }
}

func (f *Filter) CreativeCommons(v uint64) {
   f.Message[6] = v
}

func (f *Filter) Duration(v uint64) {
   f.Message[3] = v
}

func (f *Filter) FourK(v uint64) {
   f.Message[14] = v
}

func (f *Filter) HD(v uint64) {
   f.Message[4] = v
}

func (f *Filter) HDR(v uint64) {
   f.Message[25] = v
}

func (f *Filter) Live(v uint64) {
   f.Message[8] = v
}

func (f *Filter) Location(v uint64) {
   f.Message[23] = v
}

func (f *Filter) Purchased(v uint64) {
   f.Message[9] = v
}

func (f *Filter) Subtitles(v uint64) {
   f.Message[5] = v
}

func (f *Filter) ThreeD(v uint64) {
   f.Message[7] = v
}

func (f *Filter) ThreeSixty(v uint64) {
   f.Message[15] = v
}

func (f *Filter) Type(v uint64) {
   f.Message[2] = v
}

func (f *Filter) UploadDate(v uint64) {
   f.Message[1] = v
}

func (f *Filter) VR180(v uint64) {
   f.Message[26] = v
}

type Params struct {
   protobuf.Message
}

func NewParams() Params {
   return Params{
      make(protobuf.Message),
   }
}

func (p Params) Encode() string {
   return base64.StdEncoding.EncodeToString(p.Marshal())
}

func (p *Params) Filter(v Filter) {
   p.Message[2] = v.Message
}

func (p *Params) SortBy(v uint64) {
   p.Message[1] = v
}
