package youtube

import (
   "github.com/89z/rosso/protobuf"
)

var Duration = map[string]uint64{
   "Under 4 minutes": 1,
   "4 - 20 minutes": 3,
   "Over 20 minutes": 2,
}

var Features = map[string]protobuf.Number{
   "360°": 15,
   "3D": 7,
   "4K": 14,
   "Creative Commons": 6,
   "HD": 4,
   "HDR": 25,
   "Live": 8,
   "Location": 23,
   "Purchased": 9,
   "Subtitles/CC": 5,
   "VR180": 26,
}

var Sort_By = map[string]uint64{
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

var Upload_Date = map[string]uint64{
   "Last hour": 1,
   "Today": 2,
   "This week": 3,
   "This month": 4,
   "This year": 5,
}

type Filter struct {
   protobuf.Message
}

func New_Filter() Filter {
   var filter Filter
   filter.Message = make(protobuf.Message)
   return filter
}

type Params struct {
   protobuf.Message
}

func New_Params() Params {
   var par Params
   par.Message = make(protobuf.Message)
   return par
}

func (self Filter) Features(num protobuf.Number) {
   self.Message[num] = protobuf.Varint(1)
}

func (self Filter) Duration(val uint64) {
   self.Message[3] = protobuf.Varint(val)
}

func (self Filter) Type(val uint64) {
   self.Message[2] = protobuf.Varint(val)
}

func (self Filter) Upload_Date(val uint64) {
   self.Message[1] = protobuf.Varint(val)
}

func (self Params) Filter(val Filter) {
   self.Message[2] = val.Message
}

func (self Params) Sort_By(val uint64) {
   self.Message[1] = protobuf.Varint(val)
}
