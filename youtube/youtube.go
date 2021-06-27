// YouTube
package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const (
   chunk = 10_000_000
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

type Context struct {
   Client struct {
      ClientName string `json:"clientName"`
      ClientVersion string `json:"clientVersion"`
   } `json:"client"`
}

type Search struct {
   Context `json:"context"`
   Query string `json:"query"`
}

type player struct {
   Context `json:"context"`
   VideoID string `json:"videoId"`
}

func newPlayer(id, name, version string) player {
   var p player
   p.Context.Client.ClientName = name
   p.Context.Client.ClientVersion = version
   p.VideoID = id
   return p
}

func (p player) post() (*http.Response, error) {
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


type Result struct {
   Contents struct {
      TwoColumnSearchResultsRenderer `json:"twoColumnSearchResultsRenderer"`
   }
}

func (r Result) VideoRenderers() []VideoRenderer {
   var vids []VideoRenderer
   for _, sect := range r.Contents.PrimaryContents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.VideoRenderer)
      }
   }
   return vids
}

type TwoColumnSearchResultsRenderer struct {
   PrimaryContents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents []struct {
                  VideoRenderer `json:"videoRenderer"`
               }
            }
         }
      }
   }
}

type VideoRenderer struct {
   VideoID string
}
