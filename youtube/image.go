package youtube
import "fmt"

var Images = []Image{
   {Right: "0", Ext: "jpg"},
   {Right: "1", Ext: "jpg"},
   {Right: "2", Ext: "jpg"},
   {Right: "3", Ext: "jpg"},
   {Right: "default", Ext: "jpg"},
   {Left: "hq", Right: "1", Ext: "jpg"},
   {Left: "hq", Right: "2", Ext: "jpg"},
   {Left: "hq", Right: "3", Ext: "jpg"},
   {Left: "hq", Right: "720", Ext: "jpg"},
   {Left: "hq", Right: "default", Ext: "jpg"},
   {Left: "maxres", Right: "1", Ext: "jpg"},
   {Left: "maxres", Right: "2", Ext: "jpg"},
   {Left: "maxres", Right: "3", Ext: "jpg"},
   {Left: "maxres", Right: "default", Ext: "jpg"},
   {Left: "mq", Right: "1", Ext: "jpg"},
   {Left: "mq", Right: "2", Ext: "jpg"},
   {Left: "mq", Right: "3", Ext: "jpg"},
   {Left: "mq", Right: "default", Ext: "jpg"},
   {Left: "sd", Right: "1", Ext: "jpg"},
   {Left: "sd", Right: "2", Ext: "jpg"},
   {Left: "sd", Right: "3", Ext: "jpg"},
   {Left: "sd", Right: "default", Ext: "jpg"},
   {Vi: "_webp", Right: "0", Ext: "webp"},
   {Vi: "_webp", Right: "1", Ext: "webp"},
   {Vi: "_webp", Right: "2", Ext: "webp"},
   {Vi: "_webp", Right: "3", Ext: "webp"},
   {Vi: "_webp", Right: "default", Ext: "webp"},
   {Vi: "_webp", Left: "hq", Right: "1", Ext: "webp"},
   {Vi: "_webp", Left: "hq", Right: "2", Ext: "webp"},
   {Vi: "_webp", Left: "hq", Right: "3", Ext: "webp"},
   {Vi: "_webp", Left: "hq", Right: "720", Ext: "webp"},
   {Vi: "_webp", Left: "hq", Right: "default", Ext: "webp"},
   {Vi: "_webp", Left: "maxres", Right: "1", Ext: "webp"},
   {Vi: "_webp", Left: "maxres", Right: "2", Ext: "webp"},
   {Vi: "_webp", Left: "maxres", Right: "3", Ext: "webp"},
   {Vi: "_webp", Left: "maxres", Right: "default", Ext: "webp"},
   {Vi: "_webp", Left: "mq", Right: "1", Ext: "webp"},
   {Vi: "_webp", Left: "mq", Right: "2", Ext: "webp"},
   {Vi: "_webp", Left: "mq", Right: "3", Ext: "webp"},
   {Vi: "_webp", Left: "mq", Right: "default", Ext: "webp"},
   {Vi: "_webp", Left: "sd", Right: "1", Ext: "webp"},
   {Vi: "_webp", Left: "sd", Right: "2", Ext: "webp"},
   {Vi: "_webp", Left: "sd", Right: "3", Ext: "webp"},
   {Vi: "_webp", Left: "sd", Right: "default", Ext: "webp"},
}

func GetImages(id string) []string {
   img := make([]string, len(Images))
   for n, opt := range Images {
      img[n] = fmt.Sprintf(
         "http://i.ytimg.com/vi%v/%v/%v%v.%v",
         opt.Vi, id, opt.Left, opt.Right, opt.Ext,
      )
   }
   return img
}

type Image struct {
   Vi string
   Left string
   Right string
   Ext string
}
