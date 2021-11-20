package pbs

import (
   "encoding/json"
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
   req.Header.Set("X-PBS-PlatformVersion", platformVersion)
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
