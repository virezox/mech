package youtube

import (
   "fmt"
   "path"
)

const (
   HeightCrop = iota
   WidthAuto
   WidthAutoHeightBlack
   WidthBlack
)

var Images = []Image{
   {"0.jpg", 360, WidthAutoHeightBlack},
   {"0.webp", 360, WidthAutoHeightBlack},
   {"1.jpg", 90, WidthBlack},
   {"1.webp", 90, WidthBlack},
   {"2.jpg", 90, WidthBlack},
   {"2.webp", 90, WidthBlack},
   {"3.jpg", 90, WidthBlack},
   {"3.webp", 90, WidthBlack},
   {"default.jpg", 90, WidthAutoHeightBlack},
   {"default.webp", 90, WidthAutoHeightBlack},
   {"hq1.jpg", 360, WidthBlack},
   {"hq1.webp", 360, WidthBlack},
   {"hq2.jpg", 360, WidthBlack},
   {"hq2.webp", 360, WidthBlack},
   {"hq3.jpg", 360, WidthBlack},
   {"hq3.webp", 360, WidthBlack},
   {"hq720.jpg", 720, WidthAuto},
   {"hq720.webp", 720, WidthAuto},
   {"hqdefault.jpg", 360, WidthAutoHeightBlack},
   {"hqdefault.webp", 360, WidthAutoHeightBlack},
   {"maxres1.jpg", 720, WidthBlack},
   {"maxres1.webp", 720, WidthBlack},
   {"maxres2.jpg", 720, WidthBlack},
   {"maxres2.webp", 720, WidthBlack},
   {"maxres3.jpg", 720, WidthBlack},
   {"maxres3.webp", 720, WidthBlack},
   {"maxresdefault.jpg", 720, WidthAuto},
   {"maxresdefault.webp", 720, WidthAuto},
   {"mq1.jpg", 180, HeightCrop},
   {"mq1.webp", 180, HeightCrop},
   {"mq2.jpg", 180, HeightCrop},
   {"mq2.webp", 180, HeightCrop},
   {"mq3.jpg", 180, HeightCrop},
   {"mq3.webp", 180, HeightCrop},
   {"mqdefault.jpg", 180, WidthAuto},
   {"mqdefault.webp", 180, WidthAuto},
   {"sd1.jpg", 480, WidthBlack},
   {"sd1.webp", 480, WidthBlack},
   {"sd2.jpg", 480, WidthBlack},
   {"sd2.webp", 480, WidthBlack},
   {"sd3.jpg", 480, WidthBlack},
   {"sd3.webp", 480, WidthBlack},
   {"sddefault.jpg", 480, WidthAutoHeightBlack},
   {"sddefault.webp", 480, WidthAutoHeightBlack},
}

var imageDirs = map[string]string{".webp": "vi_webp", ".jpg": "vi"}

type Image struct {
   File string
   Height int
   Frame int
}

func (i Image) Address(id string) string {
   dir := imageDirs[path.Ext(i.File)]
   return fmt.Sprintf("http://i.ytimg.com/%v/%v/%v", dir, id, i.File)
}
