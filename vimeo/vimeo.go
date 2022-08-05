package vimeo

import (
   "encoding/json"
   "github.com/89z/rosso/http"
   "strconv"
)

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

func (p Progressive) String() string {
   b := []byte("Width:")
   b = strconv.AppendInt(b, p.Width, 10)
   b = append(b, " Height:"...)
   b = strconv.AppendInt(b, p.Height, 10)
   b = append(b, " FPS:"...)
   b = strconv.AppendInt(b, p.FPS, 10)
   return string(b)
}

type Check struct {
   Request struct {
      Files struct {
         Progressive []Progressive
      }
   }
}

type Progressive struct {
   Width int64
   Height int64
   FPS int64
   URL string
}

func (p Progressive) Height_Distance(v int64) int64 {
   if p.Height > v {
      return p.Height - v
   }
   return v - p.Height
}
