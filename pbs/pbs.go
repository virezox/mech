package pbs

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "path"
   "strings"
   "time"
)

const origin = "http://content.services.pbs.org"

type Video struct {
   Profile string
   URL string
}

type Asset struct {
   Resource struct {
      Duration Duration
      MP4_Videos []Video
      Title string
   }
}

func NewAsset(slug string) (*Asset, error) {
   req, err := http.NewRequest(
      "GET", origin + "/v3/android/screens/video-assets/" + slug + "/", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("x-pbs-platformversion", "5.4.2")
   req.SetBasicAuth("android", "baXE7humuVat")
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   ass := new(Asset)
   if err := json.NewDecoder(res.Body).Decode(ass); err != nil {
      return nil, err
   }
   return ass, nil
}

type Duration int64

func (d Duration) String() string {
   dur := time.Duration(d) * time.Second
   return dur.String()
}

type Episode struct {
   Episode struct {
      Assets []struct {
         Object_Type string
         Slug string
      }
   }
   Slug string
}

func NewEpisode(addr string) (*Episode, error) {
   ind := strings.Index(addr, "/video/")
   if ind == -1 {
      return nil, mech.NotFound{"/video/"}
   }
   addr = addr[:ind] + "/api" + addr[ind:]
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var video struct {
      Episodes []Episode
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   slug := path.Base(addr)
   for _, ep := range video.Episodes {
      if ep.Slug == slug {
         return &ep, nil
      }
   }
   return nil, mech.NotFound{slug}
}

func (e Episode) FullLength() (*Asset, error) {
   for _, asset := range e.Episode.Assets {
      if asset.Object_Type == "full_length" {
         return NewAsset(asset.Slug)
      }
   }
   return nil, mech.NotFound{"full_length"}
}

type Progress struct {
   *http.Response
   met []string
   x, xMax int
   y int64
}

func NewProgress(res *http.Response) *Progress {
   return &Progress{
      Response: res,
      met: []string{"B", "kB", "MB", "GB"},
      xMax: 10_000_000,
   }
}

func (p *Progress) Read(buf []byte) (int, error) {
   if p.x == 0 {
      bytes := mech.NumberFormat(float64(p.y), p.met)
      fmt.Println(mech.Percent(p.y, p.ContentLength), bytes)
   }
   num, err := p.Body.Read(buf)
   if err != nil {
      return 0, err
   }
   p.y += int64(num)
   p.x += num
   if p.x >= p.xMax {
      p.x = 0
   }
   return num, nil
}
