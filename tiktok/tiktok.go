package tiktok

import (
   "encoding/json"
   "github.com/89z/mech"
   "github.com/89z/parse/net"
   "net/http"
)

const (
   agent = "Mozilla/5.0 (Windows)"
   referer = "https://www.tiktok.com/"
)

var LogLevel mech.LogLevel

type nextData struct {
   Props struct {
      PageProps struct {
         ItemInfo struct {
            ItemStruct ItemStruct
         }
      }
   }
}

func NewItemStruct(addr string) (*ItemStruct, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", agent)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   for _, script := range net.ReadHTML(res.Body, "script") {
      if script.Attr["id"] == "__NEXT_DATA__" {
         var next nextData
         err := json.Unmarshal(script.Data, &next)
         if err != nil {
            return nil, err
         }
         return &next.Props.PageProps.ItemInfo.ItemStruct, nil
      }
   }
   return nil, mech.NotFound{"__NEXT_DATA__"}
}

type ItemStruct struct {
   Author struct {
      UniqueID string
   }
   ID string
   Video struct {
      PlayAddr string
   }
}

func (i ItemStruct) Request() (*http.Request, error) {
   req, err := http.NewRequest("GET", i.Video.PlayAddr, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Referer", referer)
   return req, nil
}
