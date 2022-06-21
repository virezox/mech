package vimeo

import (
   "encoding/json"
   "errors"
   "fmt"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

var Log format.Log

type Download struct {
   Width int
   Height int
   Quality string
   Size_Short string
   Link string
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

func (d Download) Format(f fmt.State, verb rune) {
   fmt.Fprint(f, "Width:", d.Width)
   fmt.Fprint(f, " Height:", d.Height)
   fmt.Fprint(f, " Quality:", d.Quality)
   fmt.Fprint(f, " Size:", d.Size_Short)
   if verb == 'a' {
      fmt.Fprint(f, " Link:", d.Link)
   }
}

func (v Video) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Name:", v.Name)
   fmt.Fprintln(f, "User:", v.User.Name)
   fmt.Fprintln(f, "Duration:", v.Duration)
   fmt.Fprintln(f, "Release:", v.Release_Time)
   if verb == 'a' {
      fmt.Fprintln(f, "Picture:", v.Pictures.Base_Link)
   }
   for _, down := range v.Download {
      fmt.Fprintln(f)
      down.Format(f, verb)
   }
}

type Clip struct {
   ID int
   UnlistedHash string
}

func New_Clip(address string) (*Clip, error) {
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
         _, err := fmt.Sscan(field, &clip.ID)
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

type JSON_Web struct {
   Token string
}

func New_JSON_Web() (*JSON_Web, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/_rv/jwt", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Requested-With", "XMLHttpRequest")
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
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

func (w JSON_Web) Video(clip *Clip) (*Video, error) {
   buf := fmt.Sprint("https://api.vimeo.com/videos/", clip.ID)
   if clip.UnlistedHash != "" {
      buf = fmt.Sprint(buf, ":", clip.UnlistedHash)
   }
   req, err := http.NewRequest("GET", buf, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "JWT " + w.Token)
   req.URL.RawQuery = "fields=duration,download,name,pictures,release_time,user"
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}
