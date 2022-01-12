package youtube

import (
   "encoding/base64"
   "github.com/89z/format/protobuf"
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

type Params struct {
   protobuf.Message
}

func (p Params) Encode() string {
   buf := p.Marshal()
   return base64.StdEncoding.EncodeToString(buf)
}

type Filter struct {
   protobuf.Message
}

func NewFilter() Filter {
   var filter Filter
   filter.Message = make(protobuf.Message)
   return filter
}

func NewParams() Params {
   var par Params
   par.Message = make(protobuf.Message)
   return par
}

func (f Filter) CreativeCommons(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 6}] = val
   return f
}

func (f Filter) Duration(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 3}] = val
   return f
}

func (f Filter) FourK(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 14}] = val
   return f
}

func (f Filter) HD(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 4}] = val
   return f
}

func (f Filter) HDR(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 25}] = val
   return f
}

func (f Filter) Live(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 8}] = val
   return f
}

func (f Filter) Location(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 23}] = val
   return f
}

func (f Filter) Purchased(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 9}] = val
   return f
}

func (f Filter) Subtitles(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 5}] = val
   return f
}

func (f Filter) ThreeD(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 7}] = val
   return f
}

func (f Filter) ThreeSixty(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 15}] = val
   return f
}

func (f Filter) Type(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 2}] = val
   return f
}

func (f Filter) UploadDate(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 1}] = val
   return f
}

func (f Filter) VR180(val uint64) Filter {
   f.Message[protobuf.Tag{Number: 26}] = val
   return f
}

func (p Params) SortBy(val uint64) Params {
   p.Message[protobuf.Tag{Number: 1}] = val
   return p
}

func (p Params) Filter(val Filter) Params {
   p.Message[protobuf.Tag{Number: 2}] = val.Message
   return p
}
