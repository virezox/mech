package pbs

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "path"
   "strconv"
   "strings"
)

type ErrorString string

func (e ErrorString) Error() string {
   return string(e)
}

func Slug(addr string) (string, error) {
   par, err := url.Parse(addr)
   if err != nil {
      return "", err
   }
   slug := path.Base(par.Path)
   ind := strings.Index(par.Path, "/video/")
   switch ind {
   case -1:
      return "", notPresent{"/video/"}
   case 0:
      return slug, nil
   }
   par.Path = par.Path[:ind] + "/api" + par.Path[ind:]
   req, err := http.NewRequest("GET", par.String(), nil)
   if err != nil {
      return "", err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return "", ErrorString(res.Status)
   }
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
   return "", notPresent{slug}
}

const (
   android = "baXE7humuVat"
   platformVersion = "5.4.2"
)

var LogLevel format.LogLevel

type notPresent struct {
   value string
}

func (n notPresent) Error() string {
   return strconv.Quote(n.value) + " is not present"
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

type Asset struct {
   Resource struct {
      Duration int64
      MP4_Videos []AssetVideo
      Title string
   }
}

func NewAsset(slug string) (*Asset, error) {
   var str strings.Builder
   str.WriteString("http://content.services.pbs.org")
   str.WriteString("/v3/android/screens/video-assets/")
   str.WriteString(slug)
   str.WriteByte('/')
   req, err := http.NewRequest("GET", str.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-PBS-PlatformVersion", platformVersion)
   req.SetBasicAuth("android", android)
   req.URL.RawQuery = "station_id=b3291387-78a4-41e1-beb0-da2f61a96a3e"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, ErrorString(res.Status)
   }
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
