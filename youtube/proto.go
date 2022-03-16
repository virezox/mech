package youtube

import (
   "encoding/base64"
   "github.com/89z/format/protobuf"
)

var Duration = map[string]uint64{
   "Under 4 minutes": 1,
   "4 - 20 minutes": 3,
   "Over 20 minutes": 2,
}

var Features = map[string]protobuf.Tag{
   "360°": {Number: 15},
   "3D": {Number: 7},
   "4K": {Number: 14},
   "Creative Commons": {Number: 6},
   "HD": {Number: 4},
   "HDR": {Number: 25},
   "Live": {Number: 8},
   "Location": {Number: 23},
   "Purchased": {Number: 9},
   "Subtitles/CC": {Number: 5},
   "VR180": {Number: 26},
}

var SortBy = map[string]uint64{
   "Relevance": 0,
   "Upload date": 2,
   "View count": 3,
   "Rating": 1,
}

var Type = map[string]uint64{
   "Video": 1,
   "Channel": 2,
   "Playlist": 3,
   "Movie": 4,
}

var UploadDate = map[string]uint64{
   "Last hour": 1,
   "Today": 2,
   "This week": 3,
   "This month": 4,
   "This year": 5,
}

type Filter struct {
   protobuf.Message
}

func NewFilter() Filter {
   var filter Filter
   filter.Message = make(protobuf.Message)
   return filter
}

func (f Filter) Duration(val uint64) {
   key := protobuf.Tag{Number: 3}
   f.Message[key] = val
}

func (f Filter) Features(key protobuf.Tag) {
   f.Message[key] = 1
}

func (f Filter) Type(val uint64) {
   key := protobuf.Tag{Number: 2}
   f.Message[key] = val
}

func (f Filter) UploadDate(val uint64) {
   key := protobuf.Tag{Number: 1}
   f.Message[key] = val
}

type Params struct {
   protobuf.Message
}

func NewParams() Params {
   var par Params
   par.Message = make(protobuf.Message)
   return par
}

func (p Params) Encode() string {
   buf := p.Marshal()
   return base64.StdEncoding.EncodeToString(buf)
}

func (p Params) Filter(val Filter) {
   key := protobuf.Tag{Number: 2}
   p.Message[key] = val.Message
}

func (p Params) SortBy(val uint64) {
   key := protobuf.Tag{Number: 1}
   p.Message[key] = val
}
