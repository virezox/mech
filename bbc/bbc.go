package bbc

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
)

const (
   producer = "http://walter-producer-cdn.api.bbci.co.uk/content/cps/news/"
   VideoType = "bbc.mobile.news.video"
)

const mediaSelector =
   "http://open.live.bbc.co.uk/mediaselector/6/select/version/2.0/format/json" +
   "/mediaset/mobile-phone-main/vpid/"

type newsItem struct {
   Relations []struct {
      PrimaryType string
      Content json.RawMessage
   }
}

type NewsVideo struct {
   ExternalID string
}

func NewNewsVideo(id int) (*NewsVideo, error) {
   req, err := http.NewRequest(
      "GET", producer + strconv.Itoa(id), nil,
   )
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var item newsItem
   if err := json.NewDecoder(res.Body).Decode(&item); err != nil {
      return nil, err
   }
   for _, rel := range item.Relations {
      if rel.PrimaryType == VideoType {
         video := new(NewsVideo)
         err := json.Unmarshal(rel.Content, video)
         if err != nil {
            return nil, err
         }
         return video, nil
      }
   }
   return nil, mech.NotFound{VideoType}
}

type Selector struct {
   Media []struct {
      Connection []struct {
         Href string
      }
   }
}

func NewSelector(externalID string) (*Selector, error) {
   req, err := http.NewRequest(
      "GET", mediaSelector + externalID, nil,
   )
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sel := new(Selector)
   if err := json.NewDecoder(res.Body).Decode(sel); err != nil {
      return nil, err
   }
   return sel, nil
}
