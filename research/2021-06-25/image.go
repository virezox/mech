package youtube

const (
   HeightCrop = iota
   WidthAuto
   WidthAutoHeightBlack
   WidthBlack
)

var Images = []Image{
   {"0", 360, WidthAutoHeightBlack},
   {"1", 90, WidthBlack},
   {"2", 90, WidthBlack},
   {"3", 90, WidthBlack},
   {"default", 90, WidthAutoHeightBlack},
   {"hq1", 360, WidthBlack},
   {"hq2", 360, WidthBlack},
   {"hq3", 360, WidthBlack},
   {"hq720", 720, WidthAuto},
   {"hqdefault", 360, WidthAutoHeightBlack},
   {"maxres1", 720, WidthBlack},
   {"maxres2", 720, WidthBlack},
   {"maxres3", 720, WidthBlack},
   {"maxresdefault", 720, WidthAuto},
   {"mq1", 180, HeightCrop},
   {"mq2", 180, HeightCrop},
   {"mq3", 180, HeightCrop},
   {"mqdefault", 180, WidthAuto},
   {"sd1", 480, WidthBlack},
   {"sd2", 480, WidthBlack},
   {"sd3", 480, WidthBlack},
   {"sddefault", 480, WidthAutoHeightBlack},
}

type Image struct {
   Base string
   Height int
   Frame int
}
