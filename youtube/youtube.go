package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

// REQUEST BODY ////////////////////////////////////////////////////////////////

var (
   clientAndroid = Client{"ANDROID", "15.01"}
   clientMweb = Client{"MWEB", "2.19700101"}
)

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type Request struct {
   Context struct {
      Client Client
   }
   Query string `json:"query"`
   VideoID string `json:"videoId"`
}

func (c Client) video(id string) Request {
   var r Request
   r.Context.Client = c
   r.VideoID = id
   return r
}

func (c Client) query(s string) Request {
   var r Request
   r.Context.Client = c
   r.Query = s
   return r
}

// RESPONSE BODY ///////////////////////////////////////////////////////////////

func (r Request) Post(path string) (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(r)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", "https://www.youtube.com" + path, buf)
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

// RESPONSE SEARCH /////////////////////////////////////////////////////////////

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

func SearchMweb(query string) (*Search, error) {
   res, err := clientMweb.query(query).Post("/youtubei/v1/search")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s := new(Search)
   json.NewDecoder(res.Body).Decode(s)
   return s, nil
}

// RESPONSE PLAYER /////////////////////////////////////////////////////////////

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
   VideoDetails struct {
      Author string
      ShortDescription string
      Title string
      ViewCount int `json:"viewCount,string"`
   }
}

func PlayerAndroid(id string) (*Player, error) {
   res, err := clientAndroid.video(id).Post("/youtubei/v1/player")
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

func PlayerMweb(id string) (*Player, error) {
   res, err := clientMweb.video(id).Post("/youtubei/v1/player")
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
