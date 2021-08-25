package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
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

var clients = []struct{name, version string}{
   {"ANDROID", "16.02"},
   {"ANDROID_CREATOR", "21.06.103"},
   {"ANDROID_EMBEDDED_PLAYER", "16.02"},
   {"ANDROID_KIDS", "6.02.3"},
   {"ANDROID_MUSIC", "4.32"},
   {"IOS", "16.02"},
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
   i.Context.Client.ClientName = "WEB"
   i.Context.Client.ClientVersion = "2.20210223.09.00"
   i.Query = "k4M53xndqiU"
   body, err := i.encode()
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search", body,
   )
   req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
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
   var mw mweb
   if err := json.NewDecoder(res.Body).Decode(&mw); err != nil {
      panic(err)
   }
   if mw.videoID() != "" {
      fmt.Println(i.Context.Client, "pass")
   } else {
      fmt.Println(i.Context.Client, "fail")
   }
}

func (m mweb) videoID() string {
   for _, sect := range m.Contents.SectionListRenderer.Contents {
      if sect.ItemSectionRenderer != nil {
         for _, item := range sect.ItemSectionRenderer.Contents {
            if item.CompactVideoRenderer != nil {
               return item.CompactVideoRenderer.VideoID
            }
         }
      }
   }
   return ""
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
