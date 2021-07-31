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

type search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents []map[string]struct {
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

func main() {
   f, err := os.Open("TVHTML5.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   var s search
   json.NewDecoder(f).Decode(&s)
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         if tv, ok := item["tvMusicVideoRenderer"]; ok {
            fmt.Printf("%+v\n", tv)
         }
      }
   }
}
