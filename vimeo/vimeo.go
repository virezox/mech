package vimeo

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

type Video struct {
   Duration int64
   Release_Time string
   Name string
   Pictures struct {
      Base_Link string
   }
   Download []Download
}

func (v Video) String() string {
   buf := []byte("Duration: ")
   buf = append(buf, v.Time().String()...)
   buf = append(buf, "\nRelease: "...)
   buf = append(buf, v.Release_Time...)
   buf = append(buf, "\nName: "...)
   buf = append(buf, v.Name...)
   if v.Pictures.Base_Link != "" {
      buf = append(buf, "\nPicture: "...)
      buf = append(buf, v.Pictures.Base_Link...)
   }
   for _, down := range v.Download {
      buf = append(buf, '\n')
      buf = append(buf, down.String()...)
   }
   return string(buf)
}

type Download struct {
   Public_Name string
   Width int64
   Height int64
   Size_Short string
   Link string
}

func (d Download) String() string {
   var buf []byte
   buf = append(buf, "Name:"...)
   buf = append(buf, d.Public_Name...)
   buf = append(buf, " Width:"...)
   buf = strconv.AppendInt(buf, d.Width, 10)
   buf = append(buf, " Height:"...)
   buf = strconv.AppendInt(buf, d.Height, 10)
   buf = append(buf, " Size:"...)
   buf = append(buf, d.Size_Short...)
   if d.Link != "" {
      buf = append(buf, " Link:"...)
      buf = append(buf, d.Link...)
   }
   return string(buf)
}

func NewClip(address string) (*Clip, error) {
   addr, err := url.Parse(address)
   if err != nil {
      return nil, err
   }
   fields := strings.FieldsFunc(addr.Path, func(r rune) bool {
      return r == '/'
   })
   var clip Clip
   for _, field := range fields {
      if clip.ID >= 1 {
         clip.UnlistedHash = field
      } else if field != "video" {
         clip.ID, err = strconv.ParseInt(field, 10, 64)
         if err != nil {
            return nil, err
         }
      }
   }
   for _, key := range []string{"h", "unlisted_hash"} {
      hash := addr.Query().Get(key)
      if hash != "" {
         clip.UnlistedHash = hash
      }
   }
   return &clip, nil
}

func (w JsonWeb) Video(clip *Clip) (*Video, error) {
   buf := []byte("https://api.vimeo.com/videos/")
   buf = strconv.AppendInt(buf, clip.ID, 10)
   if clip.UnlistedHash != "" {
      buf = append(buf, ':')
      buf = append(buf, clip.UnlistedHash...)
   }
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "JWT " + w.Token)
   req.URL.RawQuery = "fields=duration,download,name,pictures,release_time"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

var LogLevel format.LogLevel

type Clip struct {
   ID int64
   UnlistedHash string
}

type JsonWeb struct {
   Token string
}

func NewJsonWeb() (*JsonWeb, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/_rv/jwt", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Requested-With", "XMLHttpRequest")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(JsonWeb)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

func (v Video) Time() time.Duration {
   return time.Duration(v.Duration) * time.Second
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}

