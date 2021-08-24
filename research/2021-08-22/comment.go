package comment

import (
   "encoding/base64"
   p "google.golang.org/protobuf/testing/protopack"
)

func continuation(videoID string) string {
   m := p.Message{
      p.Tag{2, p.BytesType}, p.LengthPrefix{
         p.Tag{2, p.BytesType}, p.String(videoID),
      },
      p.Tag{3, p.VarintType}, p.Varint(6),
      p.Tag{6, p.BytesType}, p.LengthPrefix{
         p.Tag{4, p.BytesType}, p.LengthPrefix{
            p.Tag{4, p.BytesType}, p.String(videoID),
         },
      },
   }
   b := m.Marshal()
   return base64.StdEncoding.EncodeToString(b)
}
