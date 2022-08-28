package vimeo

import (
   "encoding/json"
   "github.com/89z/rosso/http"
   "net/url"
   "strconv"
   "strings"
)

type Clip struct {
   ID int64
   Unlisted_Hash string
}

func New_Clip(address string) (*Clip, error) {
   ref, err := url.Parse(address)
   if err != nil {
      return nil, err
   }
   fields := strings.FieldsFunc(ref.Path, func(r rune) bool {
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
      hash := ref.Query().Get(key)
      if hash != "" {
         clip.Unlisted_Hash = hash
      }
   }
   return &clip, nil
}

type Video struct {
   Duration int64
   Name string
   Pictures struct {
      Base_Link string
   }
   Release_Time string
   User struct {
      Name string
   }
   Download []struct {
      Width int64
      Height int64
      Link string
      Quality string
      Size_Short string
   }
}

func New_JSON_Web() (*JSON_Web, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/_next/jwt", nil)
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

var Client = http.Default_Client

type JSON_Web struct {
   Token string
}

func (w JSON_Web) Video(clip *Clip) (*Video, error) {
   b := []byte("https://api.vimeo.com/videos/")
   b = strconv.AppendInt(b, clip.ID, 10)
   if clip.Unlisted_Hash != "" {
      b = append(b, ':')
      b = append(b, clip.Unlisted_Hash...)
   }
   req, err := http.NewRequest("GET", string(b), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "JWT " + w.Token)
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
