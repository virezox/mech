package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "os"
)

var (
   _ = fmt.Print
   _ = io.ReadAll
   _ = os.Stdout
)

type request struct {
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Params string `json:"params,omitempty"`
   Query string `json:"query"`
}

type response struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents []struct {
                  TvMusicVideoRenderer struct {
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
            }
         }
      }
   }
}

func main() {
   var s request
   s.Context.Client.ClientName = "TVHTML5"
   s.Context.Client.ClientVersion = "7.20200101"
   s.Params = "EgIQAQ"
   s.Query = "fleet foxes wading in waist high water"
   body := new(bytes.Buffer)
   json.NewEncoder(body).Encode(s)
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
   var r response
   json.NewDecoder(res.Body).Decode(&r)
   for _, sect := range r.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         fmt.Printf("%+v\n", item.TvMusicVideoRenderer)
      }
   }
}
