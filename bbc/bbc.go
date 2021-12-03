package bbc

import (
   "encoding/json"
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "path"
   "strconv"
)

// bbc.com/news/av/10462520
func Valid(id string) bool {
   if len(id) == 8 {
      return true
   }
   return false
}

const producer = "http://walter-producer-cdn.api.bbci.co.uk/content/cps/news/"

type newsItem struct {
   Relations []struct {
      PrimaryType string
      Content json.RawMessage
   }
}

const VideoType = "bbc.mobile.news.video"

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

type selector struct {
   Media []struct {
      Kind string
      Connection json.RawMessage
   }
}

type NewsVideo struct {
   ExternalID string
}

const mediaSelector =
   "http://open.live.bbc.co.uk" +
   "/mediaselector/6/select/version/2.0/mediaset/pc/vpid/"

func (c Connection) HLS() ([]m3u.Format, error) {
   req, err := http.NewRequest("GET", c.Href, nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(c.Href)
   return m3u.Decode(res.Body, dir)
}

type Connection struct {
   Protocol string
   Supplier string
   TransferFormat string
   Href string
}

func (n NewsVideo) Connection() (*Connection, error) {
   req, err := http.NewRequest("GET", mediaSelector + n.ExternalID, nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var sel selector
   if err := json.NewDecoder(res.Body).Decode(&sel); err != nil {
      return nil, err
   }
   for _, media := range sel.Media {
      if media.Kind == "video" {
         var cons []Connection
         err := json.Unmarshal(media.Connection, &cons)
         if err != nil {
            return nil, err
         }
         for _, con := range cons {
            if con.Protocol == "http" && con.Supplier == "mf_akamai" {
               if con.TransferFormat == "hls" {
                  return &con, nil
               }
            }
         }
         return nil, mech.NotFound{"http,hls,mf_akamai"}
      }
   }
   return nil, mech.NotFound{"video"}
}
