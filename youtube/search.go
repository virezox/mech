package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

var Mweb = Client{"MWEB", "2.19700101"}

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
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

func NewSearch(query string) (*Search, error) {
   var req request
   req.Context.Client = Mweb
   req.Query = query
   res, err := req.post("/youtubei/v1/search")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s := new(Search)
   if err := json.NewDecoder(res.Body).Decode(s); err != nil {
      return nil, err
   }
   return s, nil
}

func (r Search) Videos() []CompactVideoRenderer {
   var vids []CompactVideoRenderer
   for _, sect := range r.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.CompactVideoRenderer)
      }
   }
   return vids
}

type request struct {
   Context struct {
      Client Client `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
   VideoID string `json:"videoId"`
}

func (r request) post(path string) (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(r)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + path, buf)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, req.Method, reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
}
