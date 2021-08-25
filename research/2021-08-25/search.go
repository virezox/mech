package main

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type youTubeI struct {
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
}

func (i youTubeI) encode() (*bytes.Buffer, error) {
   b := new(bytes.Buffer)
   if err := json.NewEncoder(b).Encode(i); err != nil {
      return nil, err
   }
   return b, nil
}

func main() {
   var i youTubeI
   i.Context.Client.ClientName = "MWEB"
   i.Context.Client.ClientVersion = "2.19700101"
   i.Query = "k4M53xndqiU"
   body, err := i.encode()
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search", body,
   )
   req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   var mw mweb
   json.NewDecoder(res.Body).Decode(&mw)
}

type mweb struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer *struct {
               Contents []struct {
                  CompactVideoRenderer *struct {
                     VideoID string
                  }
               }
            }
         }
      }
   }
}

func (m mweb) videoID() string {
   for _, sect := range m.Contents.SectionListRenderer.Contents {
      if sect.ItemSectionRenderer != nil {
         for _, item := range sect.ItemSectionRenderer.Contents {
            return item.CompactVideoRenderer.VideoID
         }
      }
   }
   return ""
}
