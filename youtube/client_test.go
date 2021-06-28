package youtube_test

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "testing"
   "time"
)

var clients = []Client{
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

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type search struct {
   Context struct {
      Client `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
}

func (c Client) searchRequest() error {
   var s search
   s.Context.Client = c
   s.Query = "nelly furtado say it right"
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(s)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search", buf,
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   fmt.Println("search", res.Status, c.ClientName)
   return nil
}


type player struct {
   Context struct {
      Client `json:"client"`
   } `json:"context"`
   VideoID string `json:"videoId"`
}

func (c Client) playerRequest() error {
   var p player
   p.Context.Client = c
   p.VideoID = "XeojXq6ySs4"
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(p)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      return  err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   fmt.Println("player", res.Status, c.ClientName)
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   fmt.Println("len", len(data), c.ClientName)
   if bytes.Contains(data, []byte("\n        \"url\"")) {
      fmt.Println("decrypt pass", c.ClientName)
   } else {
      fmt.Println("decrypt fail", c.ClientName)
   }
   if bytes.Contains(data, []byte(`"publishDate"`)) {
      fmt.Println("publishDate pass", c.ClientName)
   } else {
      fmt.Println("publishDate fail", c.ClientName)
   }
   return nil
}

func TestClients(t *testing.T) {
   for _, client := range clients {
      // 1. player request?
      // 2. decrypted media?
      // 3. publishDate?
      // 4. size?
      if err := client.playerRequest(); err != nil {
         t.Fatal(err)
      }
      time.Sleep(100 * time.Millisecond)
      // 5. search request?
      if err := client.searchRequest(); err != nil {
         t.Fatal(err)
      }
      time.Sleep(100 * time.Millisecond)
   }
}
