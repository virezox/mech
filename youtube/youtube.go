package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type Context struct {
   Client Client
}

type Player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         AvailableCountries []string
         PublishDate string
      }
   }
   PlayabilityStatus struct {
      Reason string
   }
   StreamingData struct {
      AdaptiveFormats Formats
   }
   VideoDetails VideoDetails
}

type PlayerRequest struct {
   Context Context
   VideoID string `json:"videoId"`
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   ViewCount int `json:"viewCount,string"`
}

type CompactVideoRenderer struct {
   VideoID string
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct{
            ItemSectionRenderer struct {
               Contents	[]struct{
                  CompactVideoRenderer CompactVideoRenderer
               }
            }
         }
      }
   }
}

type SearchRequest struct {
   Context Context
   Query string `json:"query"`
}


var (
   ClientAndroid = Client{"ANDROID", "15.01"}
   ClientMWeb = Client{"MWEB", "2.19700101"}
)

func Android(id string) (*Player, error) {
   res, err := ClientAndroid.PlayerRequest(id).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   a := new(Player)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}

func (c Client) PlayerRequest(id string) PlayerRequest {
   var p PlayerRequest
   p.Context.Client = c
   p.VideoID = id
   return p
}

func MWeb(id string) (*Player, error) {
   res, err := ClientMWeb.PlayerRequest(id).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   mw := new(Player)
   if err := json.NewDecoder(res.Body).Decode(mw); err != nil {
      return nil, err
   }
   return mw, nil
}

func (p PlayerRequest) post() (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
}


const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func (r Search) Videos() []CompactVideoRenderer {
   var vids []CompactVideoRenderer
   for _, sect := range r.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.CompactVideoRenderer)
      }
   }
   return vids
}

func NewSearchRequest(query string) SearchRequest {
   var s SearchRequest
   s.Context.Client = ClientMWeb
   s.Query = query
   return s
}

func (s SearchRequest) Post() (*Search, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(s)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search", buf,
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   r := new(Search)
   if err := json.NewDecoder(res.Body).Decode(r); err != nil {
      return nil, err
   }
   return r, nil
}
