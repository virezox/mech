package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "os"
)

type youTubeI struct {
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Params string `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
}

func main() {
   var i youTubeI
   i.Params = "EgIQAQ"
   i.Query = "oneohtrix point never Nil Admirari"
   body := new(bytes.Buffer)
   for _, client := range clients {
      fmt.Println(client)
      i.Context.Client.ClientName = "MWEB"
      i.Context.Client.ClientVersion = client.version
      json.NewEncoder(body).Encode(i)
      req, err := http.NewRequest(
         "POST", "https://www.youtube.com/youtubei/v1/search", body,
      )
      if err != nil {
         panic(err)
      }
      req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      f, err := os.Create(client.name + ".json")
      if err != nil {
         panic(err)
      }
      defer f.Close()
      f.ReadFrom(res.Body)
   }
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents []Item
            }
         }
      }
   }
}

type Item struct {
   TvMusicVideoRenderer *struct {
      LengthText struct {
         SimpleText string
      }
      NavigationEndpoint struct {
         WatchEndpoint struct {
            VideoID string
         }
      }
      PrimaryText struct {
         SimpleText string
      }
   }
}
