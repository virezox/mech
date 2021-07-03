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

type player struct {
   Context struct {
      Client `json:"client"`
   } `json:"context"`
   VideoID string `json:"videoId"`
}

type result struct {
   decrypt bool
   publishDate bool
   search bool
   size int
}

func (r *result) playerRequest(c Client) error {
   r.decrypt = true
   r.publishDate = true
   for _, id := range []string{"XeojXq6ySs4"} {
      var p player
      p.Context.Client = c
      p.VideoID = id
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
      data, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      // size
      r.size = len(data)
      // decrypt
      if !bytes.Contains(data, []byte("\"itag\": 251,\n        \"url\"")) {
         r.decrypt = false
      }
      // publishDate
      if !bytes.Contains(data, []byte(`"publishDate"`)) {
         r.publishDate = false
      }
   }
   return nil
}

type search struct {
   Context struct {
      Client `json:"client"`
   } `json:"context"`
   Query string `json:"query"`
}

func TestClients(t *testing.T) {
   results := make(map[string]result)
   for _, c := range clients {
      fmt.Println(c)
      var r result
      if err := r.playerRequest(c); err != nil {
         t.Fatal(err)
      }
      time.Sleep(100 * time.Millisecond)
      if err := r.searchRequest(c); err != nil {
         t.Fatal(err)
      }
      results[c.ClientName] = r
      time.Sleep(100 * time.Millisecond)
   }
   done := make(map[string]bool)
   for kOne, vOne := range results {
      for kTwo, vTwo := range results {
         if done[kOne + kTwo] {
            continue
         }
         done[kTwo + kOne] = true
         // decrypt
         if !vOne.decrypt && !vTwo.decrypt {
            continue
         }
         // publishDate
         if !vOne.publishDate && !vTwo.publishDate {
            continue
         }
         // search
         if !vOne.search && !vTwo.search {
            continue
         }
         // print
         fmt.Printf(
            "%v %v %+v %v %+v\n", vOne.size + vTwo.size, kOne, vOne, kTwo, vTwo,
         )
      }
   }
}
