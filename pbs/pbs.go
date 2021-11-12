package pbs

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "path"
   "strings"
   "time"
)

const (
   Origin = "http://content.services.pbs.org"
   android = "baXE7humuVat"
   platformVersion = "5.4.2"
)

func Slug(addr string) (string, error) {
   par, err := url.Parse(addr)
   if err != nil {
      return "", err
   }
   ind := strings.Index(par.Path, "/video/")
   if ind == -1 {
      return "", mech.NotFound{"/video/"}
   }
   slug := path.Base(par.Path)
   if ind == 0 {
      return slug, nil
   }
   par.Path = par.Path[:ind] + "/api" + par.Path[ind:]
   req, err := http.NewRequest("GET", par.String(), nil)
   if err != nil {
      return "", err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   var vid Video
   if err := json.NewDecoder(res.Body).Decode(&vid); err != nil {
      return "", err
   }
   for _, ep := range vid.Episodes {
      if ep.Slug == slug {
         for _, asset := range ep.Episode.Assets {
            if asset.Object_Type == "full_length" {
               return asset.Slug, nil
            }
         }
      }
   }
   return "", mech.NotFound{slug}
}

type Asset struct {
   Resource struct {
      Duration Duration
      MP4_Videos []AssetVideo
      Title string
   }
}

func NewAsset(slug string) (*Asset, error) {
   req, err := http.NewRequest(
      "GET", Origin + "/v3/android/screens/video-assets/" + slug + "/", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("x-pbs-platformversion", platformVersion)
   req.SetBasicAuth("android", android)
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

type AssetVideo struct {
   Profile string
   URL string
}

type Duration int64

func (d Duration) String() string {
   dur := time.Duration(d) * time.Second
   return dur.String()
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

type Video struct {
   Episodes []struct {
      Episode struct {
         Assets []struct {
            Object_Type string
            Slug string
         }
      }
      Slug string
   }
}
