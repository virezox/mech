package vimeo

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "github.com/89z/rosso/http"
   "io"
   "net/url"
   "strconv"
   "strings"
)

type Clip struct {
   ID int64
   Unlisted_Hash string
}

func New_Clip(reference string) (*Clip, error) {
   ref, err := url.Parse(reference)
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

func (c Clip) Check(password string) (*Check, error) {
   // URL
   ref := []byte("https://player.vimeo.com/video/")
   ref = strconv.AppendInt(ref, c.ID, 10)
   ref = append(ref, "/check-password"...)
   // body
   body := new(bytes.Buffer)
   body.WriteString("password=")
   io.WriteString(base64.NewEncoder(base64.StdEncoding, body), password)
   req, err := http.NewRequest("POST", string(ref), body)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   check := new(Check)
   if err := json.NewDecoder(res.Body).Decode(check); err != nil {
      return nil, err
   }
   return check, nil
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
