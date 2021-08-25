package youtube

import (
   "encoding/base64"
   "google.golang.org/protobuf/testing/protopack"
)

const (
   BytesType = protopack.BytesType
   VarintType = protopack.VarintType
)

var Params = map[string]map[string]Message{
   "SORT BY": {
      "Relevance": {
         Tag{1, VarintType}, Varint(0),
      },
      "Rating": {
         Tag{1, VarintType}, Varint(1),
      },
      "Upload date": {
         Tag{1, VarintType}, Varint(2),
      },
      "View count": {
         Tag{1, VarintType}, Varint(3),
      },
   },
   "UPLOAD DATE": {
      "Last hour": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{1, VarintType}, Varint(1),
         },
      },
      "Today": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{1, VarintType}, Varint(2),
         },
      },
      "This week": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{1, VarintType}, Varint(3),
         },
      },
      "This month": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{1, VarintType}, Varint(4),
         },
      },
      "This year": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{1, VarintType}, Varint(5),
         },
      },
   },
   "TYPE": {
      "Video": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{2, VarintType}, Varint(1),
         },
      },
      "Channel": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{2, VarintType}, Varint(2),
         },
      },
      "Playlist": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{2, VarintType}, Varint(3),
         },
      },
      "Movie": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{2, VarintType}, Varint(4),
         },
      },
   },
   "DURATION": {
      "Under 4 minutes": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{3, VarintType}, Varint(1),
         },
      },
      "Over 20 minutes": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{3, VarintType}, Varint(2),
         },
      },
      "4 - 20 minutes": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{3, VarintType}, Varint(3),
         },
      },
   },
   "FEATURES": {
      "HD": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{4, VarintType}, Varint(1),
         },
      },
      "Subtitles/CC": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{5, VarintType}, Varint(1),
         },
      },
      "Creative Commons": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{6, VarintType}, Varint(1),
         },
      },
      "3D": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{7, VarintType}, Varint(1),
         },
      },
      "Live": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{8, VarintType}, Varint(1),
         },
      },
      "Purchased": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{9, VarintType}, Varint(1),
         },
      },
      "4K": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{14, VarintType}, Varint(1),
         },
      },
      "360°": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{15, VarintType}, Varint(1),
         },
      },
      "Location": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{23, VarintType}, Varint(1),
         },
      },
      "HDR": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{25, VarintType}, Varint(1),
         },
      },
      "VR180": {
         Tag{2, BytesType}, LengthPrefix{
            Tag{26, VarintType}, Varint(1),
         },
      },
   },
}

type LengthPrefix = protopack.LengthPrefix

type Message protopack.Message

func Continuation(videoID string) Message {
   return Message{
      Tag{2, BytesType}, LengthPrefix{
         Tag{2, BytesType}, String(videoID),
      },
      Tag{3, VarintType}, Varint(6),
      Tag{6, BytesType}, LengthPrefix{
         Tag{4, BytesType}, LengthPrefix{
            Tag{4, BytesType}, String(videoID),
         },
      },
   }
}

func (m Message) Encode() string {
   b := protopack.Message(m).Marshal()
   return base64.StdEncoding.EncodeToString(b)
}

type String = protopack.String

type Tag = protopack.Tag

type Varint = protopack.Varint
