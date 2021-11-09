package youtube

type mParam struct {
   SortBy int `json:"varint,1"`
   Filter struct {
      UploadDate int `protobuf:"varint,1"`
      Type int `protobuf:"varint,2"`
      Duration int `protobuf:"varint,3"`
      HD int `protobuf:"varint,4"`
      Subtitles int `protobuf:"varint,5"`
      CreativeCommons int `protobuf:"varint,6"`
      ThreeD int `protobuf:"varint,7"`
      Live int `protobuf:"varint,8"`
      Purchased int `protobuf:"varint,9"`
      FourK int `protobuf:"varint,14"`
      ThreeSixty int `protobuf:"varint,15"`
      Location int `protobuf:"varint,23"`
      HDR int `protobuf:"varint,25"`
      VR180 int `protobuf:"varint,26"`
   } `protobuf:"bytes,2"`
}
