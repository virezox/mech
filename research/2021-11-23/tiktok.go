package tiktok

import (
   "github.com/89z/mech"
   "github.com/89z/parse/html"
   "github.com/89z/parse/json"
   "net/http"
   stdjson "encoding/json"
)

const agent = "Mozilla/5.0 (Windows)"

type data interface {
   playAddr() string
}

type nextData struct {
   Props struct {
      PageProps struct {
         ItemInfo struct {
            ItemStruct struct {
               Video struct {
                  PlayAddr string
               }
            }
         }
      }
   }
}

func (n nextData) playAddr() string {
   return n.Props.PageProps.ItemInfo.ItemStruct.Video.PlayAddr
}

type sigiData struct {
   ItemModule map[int]struct {
      Video struct {
         PlayAddr string
      }
   }
}

func (s sigiData) playAddr() string {
   for key := range s.ItemModule {
      return s.ItemModule[key].Video.PlayAddr
   }
   return ""
}

func newData(addr string) (data, error) {
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
   var ids string
   for _, script := range html.Parse(res.Body, "script") {
      id := script.Attr["id"]
      switch id {
      case "__NEXT_DATA__":
         next := new(nextData)
         if err := stdjson.Unmarshal(script.Data, next); err != nil {
            return nil, err
         }
         return next, nil
      case "sigi-persisted-data":
         sigi := new(sigiData)
         ok := json.NewDecoder(script.Data).Decode(sigi, '{')
         if ok {
            return sigi, nil
         }
      default:
         if id != "" {
            if ids != "" {
               ids += ","
            }
            ids += id
         }
      }
   }
   return nil, mech.NotFound{ids}
}
