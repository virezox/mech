package tiktok

import (
   "github.com/89z/mech"
   "github.com/89z/parse/json"
   "github.com/89z/parse/net"
   "net/http"
   stdjson "encoding/json"
)

const (
   agent = "Mozilla/5.0 (Windows)"
   referer = "https://www.tiktok.com/"
)

func Request(vid Video) (*http.Request, error) {
   req, err := http.NewRequest("GET", vid.PlayAddr(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Referer", referer)
   return req, nil
}

type NextData struct {
   Props struct {
      PageProps struct {
         ItemInfo struct {
            ItemStruct struct {
               Author struct {
                  UniqueID string
               }
               ID string
               Video struct {
                  PlayAddr string
               }
            }
         }
      }
   }
}

func (n NextData) Author() string {
   return n.Props.PageProps.ItemInfo.ItemStruct.Author.UniqueID
}

func (n NextData) ID() string {
   return n.Props.PageProps.ItemInfo.ItemStruct.ID
}

func (n NextData) PlayAddr() string {
   return n.Props.PageProps.ItemInfo.ItemStruct.Video.PlayAddr
}

type SigiData struct {
   ItemModule map[int]struct {
      Author string
      ID string
      Video struct {
         PlayAddr string
      }
   }
}

func (s SigiData) Author() string {
   for key := range s.ItemModule {
      return s.ItemModule[key].Author
   }
   return ""
}

func (s SigiData) ID() string {
   for key := range s.ItemModule {
      return s.ItemModule[key].ID
   }
   return ""
}

func (s SigiData) PlayAddr() string {
   for key := range s.ItemModule {
      return s.ItemModule[key].Video.PlayAddr
   }
   return ""
}

type Video interface {
   Author() string
   ID() string
   PlayAddr() string
}

func NewVideo(addr string) (Video, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", agent)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   for _, script := range net.ReadHTML(res.Body, "script") {
      switch script.Attr["id"] {
      case "__NEXT_DATA__":
         next := new(NextData)
         err := stdjson.Unmarshal(script.Data, next)
         if err != nil {
            return nil, err
         }
         return next, nil
      case "sigi-persisted-data":
         sigi := new(SigiData)
         ok := json.NewDecoder(script.Data).Object(sigi)
         if ok {
            return sigi, nil
         }
      }
   }
   return nil, mech.NotFound{"script"}
}
