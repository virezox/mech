package youtube
import "fmt"

var Images = []Image{
   {"vi", "0", "jpg"},
   {"vi", "1", "jpg"},
   {"vi", "2", "jpg"},
   {"vi", "3", "jpg"},
   {"vi", "default", "jpg"},
   {"vi", "hq1", "jpg"},
   {"vi", "hq2", "jpg"},
   {"vi", "hq3", "jpg"},
   {"vi", "hq720", "jpg"},
   {"vi", "hqdefault", "jpg"},
   {"vi", "maxres1", "jpg"},
   {"vi", "maxres2", "jpg"},
   {"vi", "maxres3", "jpg"},
   {"vi", "maxresdefault", "jpg"},
   {"vi", "mq1", "jpg"},
   {"vi", "mq2", "jpg"},
   {"vi", "mq3", "jpg"},
   {"vi", "mqdefault", "jpg"},
   {"vi", "sd1", "jpg"},
   {"vi", "sd2", "jpg"},
   {"vi", "sd3", "jpg"},
   {"vi", "sddefault", "jpg"},
   {"vi_webp", "0", "webp"},
   {"vi_webp", "1", "webp"},
   {"vi_webp", "2", "webp"},
   {"vi_webp", "3", "webp"},
   {"vi_webp", "default", "webp"},
   {"vi_webp", "hq1", "webp"},
   {"vi_webp", "hq2", "webp"},
   {"vi_webp", "hq3", "webp"},
   {"vi_webp", "hq720", "webp"},
   {"vi_webp", "hqdefault", "webp"},
   {"vi_webp", "maxres1", "webp"},
   {"vi_webp", "maxres2", "webp"},
   {"vi_webp", "maxres3", "webp"},
   {"vi_webp", "maxresdefault", "webp"},
   {"vi_webp", "mq1", "webp"},
   {"vi_webp", "mq2", "webp"},
   {"vi_webp", "mq3", "webp"},
   {"vi_webp", "mqdefault", "webp"},
   {"vi_webp", "sd1", "webp"},
   {"vi_webp", "sd2", "webp"},
   {"vi_webp", "sd3", "webp"},
   {"vi_webp", "sddefault", "webp"},
}

type Image struct {
   Dir string
   Base string
   Ext string
}

func (i Image) Address(id string) string {
   return fmt.Sprintf(
      "http://i.ytimg.com/%v/%v/%v.%v", i.Dir, id, i.Base, i.Ext,
   )
}
