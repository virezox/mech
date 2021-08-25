package youtube

import (
   "encoding/base64"
   "google.golang.org/protobuf/testing/protopack"
)

const (
   bytesType = protopack.BytesType
   varintType = protopack.VarintType
)

var Params = map[string]map[string]Message{
   "SORT BY": {
      "Relevance": {
         tag{1, varintType}, varint(0),
      },
      "Rating": {
         tag{1, varintType}, varint(1),
      },
      "Upload date": {
         tag{1, varintType}, varint(2),
      },
      "View count": {
         tag{1, varintType}, varint(3),
      },
   },
   "UPLOAD DATE": {
      "Last hour": {
         tag{2, bytesType}, lengthPrefix{
            tag{1, varintType}, varint(1),
         },
      },
      "Today": {
         tag{2, bytesType}, lengthPrefix{
            tag{1, varintType}, varint(2),
         },
      },
      "This week": {
         tag{2, bytesType}, lengthPrefix{
            tag{1, varintType}, varint(3),
         },
      },
      "This month": {
         tag{2, bytesType}, lengthPrefix{
            tag{1, varintType}, varint(4),
         },
      },
      "This year": {
         tag{2, bytesType}, lengthPrefix{
            tag{1, varintType}, varint(5),
         },
      },
   },
   "TYPE": {
      "Video": {
         tag{2, bytesType}, lengthPrefix{
            tag{2, varintType}, varint(1),
         },
      },
      "Channel": {
         tag{2, bytesType}, lengthPrefix{
            tag{2, varintType}, varint(2),
         },
      },
      "Playlist": {
         tag{2, bytesType}, lengthPrefix{
            tag{2, varintType}, varint(3),
         },
      },
      "Movie": {
         tag{2, bytesType}, lengthPrefix{
            tag{2, varintType}, varint(4),
         },
      },
   },
   "DURATION": {
      "Under 4 minutes": {
         tag{2, bytesType}, lengthPrefix{
            tag{3, varintType}, varint(1),
         },
      },
      "Over 20 minutes": {
         tag{2, bytesType}, lengthPrefix{
            tag{3, varintType}, varint(2),
         },
      },
      "4 - 20 minutes": {
         tag{2, bytesType}, lengthPrefix{
            tag{3, varintType}, varint(3),
         },
      },
   },
   "FEATURES": {
      "HD": {
         tag{2, bytesType}, lengthPrefix{
            tag{4, varintType}, varint(1),
         },
      },
      "Subtitles/CC": {
         tag{2, bytesType}, lengthPrefix{
            tag{5, varintType}, varint(1),
         },
      },
      "Creative Commons": {
         tag{2, bytesType}, lengthPrefix{
            tag{6, varintType}, varint(1),
         },
      },
      "3D": {
         tag{2, bytesType}, lengthPrefix{
            tag{7, varintType}, varint(1),
         },
      },
      "Live": {
         tag{2, bytesType}, lengthPrefix{
            tag{8, varintType}, varint(1),
         },
      },
      "Purchased": {
         tag{2, bytesType}, lengthPrefix{
            tag{9, varintType}, varint(1),
         },
      },
      "4K": {
         tag{2, bytesType}, lengthPrefix{
            tag{14, varintType}, varint(1),
         },
      },
      "360°": {
         tag{2, bytesType}, lengthPrefix{
            tag{15, varintType}, varint(1),
         },
      },
      "Location": {
         tag{2, bytesType}, lengthPrefix{
            tag{23, varintType}, varint(1),
         },
      },
      "HDR": {
         tag{2, bytesType}, lengthPrefix{
            tag{25, varintType}, varint(1),
         },
      },
      "VR180": {
         tag{2, bytesType}, lengthPrefix{
            tag{26, varintType}, varint(1),
         },
      },
   },
}

type lengthPrefix = protopack.LengthPrefix

type Message protopack.Message

func Continuation(videoID string) Message {
   return Message{
      tag{2, bytesType}, lengthPrefix{
         tag{2, bytesType}, protopack.String(videoID),
      },
      tag{3, varintType}, varint(6),
      tag{6, bytesType}, lengthPrefix{
         tag{4, bytesType}, lengthPrefix{
            tag{4, bytesType}, protopack.String(videoID),
         },
      },
   }
}

func (m Message) Encode() string {
   b := protopack.Message(m).Marshal()
   return base64.StdEncoding.EncodeToString(b)
}

type tag = protopack.Tag

type varint = protopack.Varint
