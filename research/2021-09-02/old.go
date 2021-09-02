package youtube
import p "google.golang.org/protobuf/testing/protopack"

var Params = map[string]map[string]p.Message{
   "FEATURES": {
      "HD": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{4, p.VarintType}, p.Varint(1),
         },
      },
      "Subtitles/CC": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{5, p.VarintType}, p.Varint(1),
         },
      },
      "Creative Commons": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{6, p.VarintType}, p.Varint(1),
         },
      },
      "3D": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{7, p.VarintType}, p.Varint(1),
         },
      },
      "Live": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{8, p.VarintType}, p.Varint(1),
         },
      },
      "Purchased": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{9, p.VarintType}, p.Varint(1),
         },
      },
      "4K": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{14, p.VarintType}, p.Varint(1),
         },
      },
      "360°": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{15, p.VarintType}, p.Varint(1),
         },
      },
      "Location": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{23, p.VarintType}, p.Varint(1),
         },
      },
      "HDR": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{25, p.VarintType}, p.Varint(1),
         },
      },
      "VR180": {
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{26, p.VarintType}, p.Varint(1),
         },
      },
   },
}
