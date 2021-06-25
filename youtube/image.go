package youtube
import "fmt"

var Images = []Image{
   {Base: "0", Ext: "jpg"},
   {Base: "1", Ext: "jpg"},
   {Base: "2", Ext: "jpg"},
   {Base: "3", Ext: "jpg"},
   {Base: "default", Ext: "jpg"},
   {Base: "hq1", Ext: "jpg"},
   {Base: "hq2", Ext: "jpg"},
   {Base: "hq3", Ext: "jpg"},
   {Base: "hq720", Ext: "jpg"},
   {Base: "hqdefault", Ext: "jpg"},
   {Base: "maxres1", Ext: "jpg"},
   {Base: "maxres2", Ext: "jpg"},
   {Base: "maxres3", Ext: "jpg"},
   {Base: "maxresdefault", Ext: "jpg"},
   {Base: "mq1", Ext: "jpg"},
   {Base: "mq2", Ext: "jpg"},
   {Base: "mq3", Ext: "jpg"},
   {Base: "mqdefault", Ext: "jpg"},
   {Base: "sd1", Ext: "jpg"},
   {Base: "sd2", Ext: "jpg"},
   {Base: "sd3", Ext: "jpg"},
   {Base: "sddefault", Ext: "jpg"},
   {Vi: "_webp", Base: "0", Ext: "webp"},
   {Vi: "_webp", Base: "1", Ext: "webp"},
   {Vi: "_webp", Base: "2", Ext: "webp"},
   {Vi: "_webp", Base: "3", Ext: "webp"},
   {Vi: "_webp", Base: "default", Ext: "webp"},
   {Vi: "_webp", Base: "hq1", Ext: "webp"},
   {Vi: "_webp", Base: "hq2", Ext: "webp"},
   {Vi: "_webp", Base: "hq3", Ext: "webp"},
   {Vi: "_webp", Base: "hq720", Ext: "webp"},
   {Vi: "_webp", Base: "hqdefault", Ext: "webp"},
   {Vi: "_webp", Base: "maxres1", Ext: "webp"},
   {Vi: "_webp", Base: "maxres2", Ext: "webp"},
   {Vi: "_webp", Base: "maxres3", Ext: "webp"},
   {Vi: "_webp", Base: "maxresdefault", Ext: "webp"},
   {Vi: "_webp", Base: "mq1", Ext: "webp"},
   {Vi: "_webp", Base: "mq2", Ext: "webp"},
   {Vi: "_webp", Base: "mq3", Ext: "webp"},
   {Vi: "_webp", Base: "mqdefault", Ext: "webp"},
   {Vi: "_webp", Base: "sd1", Ext: "webp"},
   {Vi: "_webp", Base: "sd2", Ext: "webp"},
   {Vi: "_webp", Base: "sd3", Ext: "webp"},
   {Vi: "_webp", Base: "sddefault", Ext: "webp"},
}

type Image struct {
   Vi string
   Base string
   Ext string
}

func (i Image) Address(id string) string {
   return fmt.Sprintf(
      "http://i.ytimg.com/vi%v/%v/%v.%v", i.Vi, id, i.Base, i.Ext,
   )
}
