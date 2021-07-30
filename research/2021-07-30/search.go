package main

import (
   "bytes"
   "encoding/json"
   "net/http"
   "os"
)

type compactVideoRenderer struct {
   title struct {
      runs []struct {
         text string
      }
   }
}

type search struct {
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Params string `json:"params,omitempty"`
   Query string `json:"query"`
}

func main() {
   var s search
   s.Context.Client.ClientName = "MWEB"
   s.Context.Client.ClientVersion = "2.19700101"
   // type video
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
   res, err := new(http.Client).Do(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
