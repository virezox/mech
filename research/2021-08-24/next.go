package next

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
   "net/http/httputil"
   "os"
)

type youTubeI struct {
   Context struct {
      Client youtube.Client `json:"client"`
   } `json:"context"`
   Continuation string `json:"continuation"`
}

var clients = []struct{name, version string}{
   {"ANDROID", "16.07.34"},
   {"ANDROID_CREATOR", "21.06.103"},
   {"ANDROID_EMBEDDED_PLAYER", "16.20"},
   {"ANDROID_KIDS", "6.02.3"},
   {"ANDROID_MUSIC", "4.32"},
   {"IOS", "16.05.7"},
   {"IOS_CREATOR", "20.47.100"},
   {"IOS_KIDS", "5.42.2"},
   {"IOS_MUSIC", "4.16.1"},
   {"MWEB", "2.19700101"},
   {"TVHTML5", "7.20210224.00.00"},
   {"WEB", "2.20210223.09.00"},
   {"WEB_CREATOR", "1.20210223.01.00"},
   {"WEB_EMBEDDED_PLAYER", "1.20210620.0.1"},
   {"WEB_KIDS", "2.1.3"},
   {"WEB_REMIX", "0.1"},
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
