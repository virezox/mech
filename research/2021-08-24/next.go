package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
   "net/http/httputil"
   "os"
   "time"
)

type youTubeI struct {
   Context struct {
      Client youtube.Client `json:"client"`
   } `json:"context"`
   Continuation string `json:"continuation"`
}

var clients = []struct{name, version string}{
   {"IOS", "16.05.7"},
   {"WEB", "2.20210223.09.00"},
}

func main() {
   var i youTubeI
   i.Continuation = youtube.Continuation("q5UnT4Ik6KU").Encode()
   buf := new(bytes.Buffer)
   for _, client := range clients {
      i.Context.Client.Name = client.name
      i.Context.Client.Version = client.version
      json.NewEncoder(buf).Encode(i)
      req, err := http.NewRequest(
         "POST", "https://www.youtube.com/youtubei/v1/next", buf,
      )
      if err != nil {
         panic(err)
      }
      q := req.URL.Query()
      q.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
      req.URL.RawQuery = q.Encode()
      d, err := httputil.DumpRequest(req, true)
      if err != nil {
         panic(err)
      }
      os.Stdout.Write(d)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      fmt.Println(res.Status)
      f, err := os.Create(client.name + ".json")
      if err != nil {
         panic(err)
      }
      defer f.Close()
      f.ReadFrom(res.Body)
      time.Sleep(100 * time.Millisecond)
   }
}

type next struct {
   ContinuationContents struct {
      ItemSectionContinuation struct {
         Contents []struct {
            CommentThreadRenderer struct {
               Comment struct {
                  CommentRenderer struct {
                     ContentText struct {
                        Runs []struct {
                           Text string
                        }
                     }
                  }
               }
            }
         }
      }
   }
}
