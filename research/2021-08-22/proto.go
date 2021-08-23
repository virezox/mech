type result struct {
   PageParams struct {
      VideoID string `protobuf:"2"`
   } `protobuf:"2"`
   Six uint `protobuf:"3"`
   OffsetInformation struct {
      PageInfo struct {
         VideoID string `protobuf:"4"`
         Sort uint `protobuf:"6"`
      } `protobuf:"4"`
      Offset uint `protobuf:"5"`
   } `protobuf:"6"`
}
