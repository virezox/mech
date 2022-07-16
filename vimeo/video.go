package vimeo

import (
   "encoding/json"
   "github.com/89z/rosso/http"
   "net/url"
   "strconv"
   "strings"
)

type Video struct {
   Name string
   User struct {
      Name string
   }
   Duration int64
   Release_Time string
   Pictures struct {
      Base_Link string
   }
   Download []Download
}

type Download struct {
   Width int64
   Height int64
   Quality string
   Size_Short string
   Link string
}

var Client = http.Default_Client

type Clip struct {
   ID int64
   Unlisted_Hash string
}

func New_Clip(raw string) (*Clip, error) {
   address, err := url.Parse(raw)
   if err != nil {
      return nil, err
   }
   fields := strings.FieldsFunc(address.Path, func(r rune) bool {
      return r == '/'
   })
   var clip Clip
   for _, field := range fields {
      if clip.ID >= 1 {
         clip.Unlisted_Hash = field
      } else if field != "video" {
         clip.ID, err = strconv.ParseInt(field, 10, 64)
         if err != nil {
            return nil, err
         }
      }
   }
   for _, key := range []string{"h", "unlisted_hash"} {
      hash := address.Query().Get(key)
      if hash != "" {
         clip.Unlisted_Hash = hash
      }
   }
   return &clip, nil
}

type JSON_Web struct {
   Token string
}

func New_JSON_Web() (*JSON_Web, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/_rv/jwt", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Requested-With", "XMLHttpRequest")
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(JSON_Web)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

func (self JSON_Web) Video(clip *Clip) (*Video, error) {
   buf := []byte("https://api.vimeo.com/videos/")
   buf = strconv.AppendInt(buf, clip.ID, 10)
   if clip.Unlisted_Hash != "" {
      buf = append(buf, ':')
      buf = append(buf, clip.Unlisted_Hash...)
   }
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "JWT " + self.Token)
   req.URL.RawQuery = "fields=duration,download,name,pictures,release_time,user"
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}
