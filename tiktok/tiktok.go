package tiktok

import (
   "github.com/89z/mech"
   "github.com/89z/parse/html"
   "github.com/89z/parse/json"
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
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var ids []string
   for _, script := range html.Parse(res.Body, "script") {
      id := script.Attr["id"]
      switch id {
      case "__NEXT_DATA__":
         next := new(NextData)
         if err := stdjson.Unmarshal(script.Data, next); err != nil {
            return nil, err
         }
         return next, nil
      case "sigi-persisted-data":
         sigi := new(SigiData)
         ok := json.NewDecoder(script.Data).Object(sigi)
         if ok {
            return sigi, nil
         }
      default:
         if id != "" {
            ids = append(ids, id)
         }
      }
   }
   return nil, mech.NotFound{ids}
}
