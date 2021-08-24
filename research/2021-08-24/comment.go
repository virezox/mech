package comment

import (
   "encoding/base64"
   "google.golang.org/protobuf/testing/protopack"
)

const (
   BytesType = protopack.BytesType
   VarintType = protopack.VarintType
)

type (
   LengthPrefix = protopack.LengthPrefix
   Message = protopack.Message
   String = protopack.String
   Tag = protopack.Tag
   Varint = protopack.Varint
)

func continuation(videoID string) string {
   m := Message{
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
   b := m.Marshal()
   return base64.StdEncoding.EncodeToString(b)
}
